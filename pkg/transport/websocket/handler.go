package websocket

import (
	"github.com/panjf2000/gnet/v2"
)

// OnOpenHandler 建立连接时处理
type OnOpenHandler func(gnet.Conn)

// OnCloseHandler 断开连接时处理
type OnCloseHandler func(gnet.Conn)

// OnMessageHandler 收到消息时处理
type OnMessageHandler func(gnet.Conn, []byte) error

// OnErrorHandler 出现错误时处理
type OnErrorHandler func(gnet.Conn, error)

// OnTickHandler 定时任务
type OnTickHandler func()
