package comet

import (
	"log"
	"os"

	"github.com/gqzcl/gim/internal/comet/conf"
)

var whitelist *Whitelist

// Whitelist .
type Whitelist struct {
	log  *log.Logger
	list map[int64]struct{} // whitelist for debug
}

// InitWhitelist a whitelist struct.
func InitWhitelist(c *conf.Whitelist) (err error) {
	var (
		mid int64
		f   *os.File
	)
	if f, err = os.OpenFile(c.WhiteLog, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644); err == nil {
		whitelist = new(Whitelist)
		// 新建创建一个新的记录器。out变量设置日志数据将写入的目标。prefix显示在每个生成的日志行的开头，如果提供了Lmsgprefix标志，则显示在日志头之后。flag参数定义日志记录属性。
		whitelist.log = log.New(f, "", log.LstdFlags)
		whitelist.list = make(map[int64]struct{})
		for _, mid = range c.Whitelist {
			whitelist.list[mid] = struct{}{}
		}
	}
	return
}

// Contains whitelist contains a mid or not.
func (w *Whitelist) Contains(mid int64) (ok bool) {
	if mid > 0 {
		_, ok = w.list[mid]
	}
	return
}

// Printf calls l.Output to print to the logger.
func (w *Whitelist) Printf(format string, v ...interface{}) {
	w.log.Printf(format, v...)
}
