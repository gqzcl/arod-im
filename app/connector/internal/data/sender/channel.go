package sender

import (
	jobV1 "arod-im/api/job/v1"
	"encoding/json"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
)

type Channel struct {
	Room *Room
	Next *Channel
	Prev *Channel

	conn gnet.Conn

	Uid  string
	Addr string
}

func NewChannel(conn gnet.Conn) *Channel {
	return &Channel{
		conn: conn,
	}
}

func (r *Channel) Push(msg []*jobV1.MsgBody) error {
	m, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	err = wsutil.WriteServerMessage(r.conn, ws.OpBinary, m)
	if err != nil {
		return err
	}
	return nil
}

//// Watch watch a operation.
//func (c *Channel) Watch(accepts ...int32) {
//	c.mutex.Lock()
//	for _, op := range accepts {
//		c.watchOps[op] = struct{}{}
//	}
//	c.mutex.Unlock()
//}
//
//// UnWatch unwatch an operation
//func (c *Channel) UnWatch(accepts ...int32) {
//	c.mutex.Lock()
//	for _, op := range accepts {
//		delete(c.watchOps, op)
//	}
//	c.mutex.Unlock()
//}
//
//// NeedPush verify if in watch.
//func (c *Channel) NeedPush(op int32) bool {
//	c.mutex.RLock()
//	if _, ok := c.watchOps[op]; ok {
//		c.mutex.RUnlock()
//		return true
//	}
//	c.mutex.RUnlock()
//	return false
//}
//
//// Push server push message.
//func (c *Channel) Push(p *protocol.Proto) (err error) {
//	select {
//	case c.signal <- p:
//	default:
//		err = errors.ErrSignalFullMsgDropped
//	}
//	return
//}
//
//// Ready check the channel ready or close?
//func (c *Channel) Ready() *protocol.Proto {
//	return <-c.signal
//}
//
//// Signal send signal to the channel, protocol ready.
//func (c *Channel) Signal() {
//	c.signal <- protocol.ProtoReady
//}
//
//// Close close the channel.
//func (c *Channel) Close() {
//	c.signal <- protocol.ProtoFinish
//}
