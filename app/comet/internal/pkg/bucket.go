package pkg

import (
	"sync"
	"sync/atomic"

	v1 "arod-im/api/comet/v1"
	"arod-im/api/protocol"
	"arod-im/app/comet/internal/conf"
)

// Bucket is a channel holder.
type Bucket struct {
	c           *conf.Bucket
	cLock       sync.RWMutex                // 读写锁保证chs并发安全
	channels    map[string]*Channel         // 记录key-channel
	rooms       map[string]*Room            // bucket room channels
	routines    []chan *v1.BroadcastRoomReq // 并发处理room的chan
	routinesNum uint64                      // 	routinesNum
	ipCnts      map[string]int32            // ip计数器
}

func NewBucket(c *conf.Bucket) (b *Bucket) {
	return &Bucket{
		c:           c,
		cLock:       sync.RWMutex{},
		channels:    make(map[string]*Channel),
		rooms:       make(map[string]*Room),
		routines:    make([]chan *v1.BroadcastRoomReq, 0, c.RoutineAmount),
		routinesNum: 0,
		ipCnts:      make(map[string]int32),
	}
	b = new(Bucket)
	b.channels = make(map[string]*Channel, c.Channel)
	b.ipCnts = make(map[string]int32)
	b.c = c
	b.rooms = make(map[string]*Room, c.Room)
	b.routines = make([]chan *v1.BroadcastRoomReq, c.RoutineAmount)

	for i := uint64(0); i < c.RoutineAmount; i++ {
		c := make(chan *v1.BroadcastRoomReq, c.RountineSize)
		b.routines[i] = c
		go b.roomproc(c)
	}
	return
}

// roomproc
func (b *Bucket) roomproc(c chan *v1.BroadcastRoomReq) {
	for {
		arg := <-c
		if room := b.Room(arg.RoomID); room != nil {
			room.Push(arg.Proto)
		}
	}
}

// ChannelCount channel count in the bucket
func (b *Bucket) ChannelCount() int {
	return len(b.channels)
}

// RoomCount room count in the bucket.
func (b *Bucket) RoomCount() int {
	return len(b.rooms)
}

// GetOnlineRooms get all room id where online number > 0.
func (b *Bucket) GetOnlineRooms() (res map[string]struct{}) {
	var (
		roomID string
		room   *Room
	)
	res = make(map[string]struct{})
	b.cLock.RLock()
	for roomID, room = range b.rooms {
		if room.Online > 0 {
			res[roomID] = struct{}{}
		}
	}
	b.cLock.RUnlock()
	return
}

// ChangeRoom change  room
func (b *Bucket) ChangeRoom(nrid string, ch *Channel) (err error) {
	var (
		nroom *Room
		ok    bool
		oroom = ch.Room
	)
	// change to no room
	// 如果nrid为空，则直接将ch的room置为空，同时删除原room
	if nrid == "" {
		if oroom != nil && oroom.Del(ch) {
			b.DelRoom(oroom)
		}
		ch.Room = nil
		return
	}
	b.cLock.Lock()
	// 如果id为nrid的room不存在，新建一个id为nrid的room
	if nroom, ok = b.rooms[nrid]; !ok {
		nroom = NewRoom(nrid)
		b.rooms[nrid] = nroom
	}
	b.cLock.Unlock()
	// 从原room中删除ch并从bucket中删除原room
	if oroom != nil && oroom.Del(ch) {
		b.DelRoom(oroom)
	}
	// 在新建的room中添加ch
	if err = nroom.Put(ch); err != nil {
		return
	}
	// 将ch的room置为新建的room
	ch.Room = nroom
	return
}

// Room get a room by roomid.
func (b *Bucket) Room(rid string) (room *Room) {
	b.cLock.RLock()
	room = b.rooms[rid]
	b.cLock.RUnlock()
	return
}

// DelRoom delete a room by roomid.
func (b *Bucket) DelRoom(room *Room) {
	b.cLock.Lock()
	delete(b.rooms, room.ID)
	b.cLock.Unlock()
	room.Close()
}

// BroadcastRoom broadcast a message to specified room
func (b *Bucket) BroadcastRoom(arg *v1.BroadcastRoomReq) {
	num := atomic.AddUint64(&b.routinesNum, 1) % b.c.RoutineAmount
	b.routines[num] <- arg
}

// UpRoomsCount update all room count
func (b *Bucket) UpdateRoomsCount(roomCountMap map[string]int32) {
	var (
		roomID string
		room   *Room
	)
	b.cLock.RLock()
	for roomID, room = range b.rooms {
		room.AllOnline = roomCountMap[roomID]
	}
	b.cLock.RUnlock()
}

// Put put a channel according with sub key.
func (b *Bucket) PutChannel(rid string, ch *Channel) (err error) {
	var (
		room *Room
		ok   bool
	)
	b.cLock.Lock()
	// close old channel
	if dch := b.channels[ch.Key]; dch != nil {
		dch.Close()
	}
	b.channels[ch.Key] = ch
	if rid != "" {
		// room not exist
		if room, ok = b.rooms[rid]; !ok {
			room = NewRoom(rid)
			b.rooms[rid] = room
		}
		ch.Room = room
	}
	b.ipCnts[ch.IP]++
	b.cLock.Unlock()
	if room != nil {
		err = room.Put(ch)
	}
	return
}

// Del delete the channel by sub key.
func (b *Bucket) DelChannel(dch *Channel) {
	var (
		ok   bool
		ch   *Channel
		room *Room
	)
	b.cLock.Lock()
	if ch, ok = b.channels[dch.Key]; ok {
		room = ch.Room
		if ch == dch {
			delete(b.channels, ch.Key)
		}
		// ip counter
		if b.ipCnts[ch.IP] > 1 {
			b.ipCnts[ch.IP]--
		} else {
			delete(b.ipCnts, ch.IP)
		}
	}
	b.cLock.Unlock()
	if room != nil && room.Del(ch) {
		// if empty room, must delete from bucket
		b.DelRoom(room)
	}
}

// Channel get a channel by sub key.
func (b *Bucket) GetChannel(key string) (ch *Channel) {
	b.cLock.RLock()
	ch = b.channels[key]
	b.cLock.RUnlock()
	return
}

// Broadcast push msgs to all channels in the bucket.
func (b *Bucket) Broadcast(p *protocol.Proto, op int32) {
	var ch *Channel
	b.cLock.RLock()
	for _, ch = range b.channels {
		if !ch.NeedPush(op) {
			continue
		}
		_ = ch.Push(p)
	}
	b.cLock.RUnlock()
}

// IPCount get ip count.
func (b *Bucket) IPCount() (res map[string]struct{}) {
	var (
		ip string
	)
	b.cLock.RLock()
	res = make(map[string]struct{}, len(b.ipCnts))
	for ip = range b.ipCnts {
		res[ip] = struct{}{}
	}
	b.cLock.RUnlock()
	return
}
