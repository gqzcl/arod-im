package sender

import "sync"

type Room struct {
	RoomId    string
	rLock     sync.RWMutex
	next      *Channel
	drop      bool
	Online    int32 // dirty read is ok
	AllOnline int32
}

// NewRoom new a room struct, store Channel room info.
func NewRoom(roomId string) *Room {
	return &Room{
		RoomId: roomId,
		next:   nil,
		drop:   false,
		Online: 0,
	}
}

// Put put channel into the room.
func (r *Room) Put(ch *Channel) (err error) {
	r.rLock.Lock()
	if !r.drop {
		if r.next != nil {
			r.next.Prev = ch
		}
		ch.Next = r.next
		ch.Prev = nil
		r.next = ch // insert to header
		r.Online++
	} else {
		//err = errors.ErrRoomDroped
	}
	r.rLock.Unlock()
	return
}

// Del delete channel from the room.
func (r *Room) DelCh(ch *Channel) bool {
	r.rLock.Lock()
	if ch.Prev == nil && ch.Next == nil {
		r.rLock.Unlock()
		return false
	}
	if ch.Next != nil {
		// if not footer
		ch.Next.Prev = ch.Prev
	}
	if ch.Prev != nil {
		// if not header
		ch.Prev.Next = ch.Next
	} else {
		r.next = ch.Next
	}
	ch.Next = nil
	ch.Prev = nil
	r.Online--
	r.drop = r.Online == 0
	r.rLock.Unlock()
	return r.drop
}

// Push push msg to the room, if chan full discard it.
//func (r *Room) Push(p *protocol.Proto) {
//	r.rLock.RLock()
//	for ch := r.next; ch != nil; ch = ch.Next {
//		_ = ch.Push(p)
//	}
//	r.rLock.RUnlock()
//}

// Close close the room.
//func (r *Room) Close() {
//	r.rLock.RLock()
//	for ch := r.next; ch != nil; ch = ch.Next {
//		ch.Close()
//	}
//	r.rLock.RUnlock()
//}

// OnlineNum the room all online.
func (r *Room) OnlineNum() int32 {
	if r.AllOnline > 0 {
		return r.AllOnline
	}
	return r.Online
}