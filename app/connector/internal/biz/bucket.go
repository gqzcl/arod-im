package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/panjf2000/gnet/v2"
	"sync"
)

// BucketRepo  is a Bucket repo.
type BucketRepo interface {
	RemoveCh(address string)
	AddCh(address string, c gnet.Conn)
	SingleSend(address string, msg []byte)
}

// BucketUsecase  is a Bucket usecase.
type BucketUsecase struct {
	bucket BucketRepo
	log    *log.Helper
}

func NewBucketUsecase(bucket BucketRepo, logger log.Logger) *BucketUsecase {
	return &BucketUsecase{
		bucket: bucket,
		log:    log.NewHelper(logger),
	}
}

// DeleteCh remove Connection with ID connectId
func (b *BucketUsecase) DeleteCh(address string) {
	b.bucket.RemoveCh(address)
}

func (b *BucketUsecase) AddCh(address string, c gnet.Conn) {
	b.bucket.AddCh(address, c)
}

func (b *BucketUsecase) SingleSend(address string, msg []byte) {
	b.bucket.SingleSend(address, msg)
}

type Bucket struct {
	channels sync.Map // <uid channel>
	//channels map[string]*Channel
	rooms sync.Map // <rid room>

	//chans []chan *v1.BroadcastRoomReq
}

// NewBucket 初始化Bucket
func NewBucket() (b *Bucket) {
	b = new(Bucket)
	// 设置Xmap负载因子
	//f := &xmm.Factory{}
	//mm, _ := f.CreateMemory(0.75)
	// 初始化channels和rooms
	//b.channels, _ = xmap.NewMap(mm, xds.String, xds.Interface)
	//b.channels = make(map[string]*Channel)
	//b.rooms, _ = xmap.NewMap(mm, xds.String, xds.Struct)
	return
}

//func (b *Bucket) AddChannel(key string, ch *Channel) {
//	b.channels.Store(key, ch)
//	//b.channels[key] = ch
//}
//
//// GetChannel 通过key获取Channel
//func (b *Bucket) GetChannel(key string) *Channel {
//	c, ok := b.channels.Load(key)
//	if !ok {
//		// TODO 返回错误
//		return nil
//	}
//	return c.(*Channel)
//}
//
//// DelChannel 从Bucket中删除一个Channel
//func (b *Bucket) DelChannel(dch *Channel) {
//	room := dch.Room
//	//if ch, ok, _ := b.channels.Get(dch.Addr); ok {
//	//	b.channels.Remove(ch.(Channel).Addr)
//	//	// TODO error
//	//}
//	if room != nil && room.Del(dch) {
//		b.DelRoom(room)
//	}
//}
//
//// GetRoom 通过rid获取Room
//func (b *Bucket) GetRoom(rid string) *Room {
//	r, _ := b.rooms.Load(rid)
//	return r.(*Room)
//}
//
//// DelRoom 从bucket中删除room
//func (b *Bucket) DelRoom(room *Room) {
//	b.rooms.Delete(room.ID)
//	// TODO error
//	room.Close()
//}
//
//func (b *Bucket) Broadcast(msg []byte) {
//	// TODO
//	b.channels.Range(func(key, val interface{}) bool {
//		c := val.(*Channel)
//		err := c.Push(msg)
//		if err != nil {
//			return false
//		}
//		fmt.Println("push msg to ", key)
//		return true
//	})
//}
