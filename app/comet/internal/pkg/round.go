package pkg

import (
	"arod-im/app/comet/internal/conf"
	"arod-im/pkg/timer"
	"arod-im/pkg/wbyte"
)

// RoundOptions round options.
type RoundOptions struct {
	Timer        int
	TimerSize    int
	Reader       int
	ReadBuf      int
	ReadBufSize  int
	Writer       int
	WriteBuf     int
	WriteBufSize int
}

// Round used for connection round-robin get a reader/writer/timer for split big lock.
type Round struct {
	readers []wbyte.Pool
	writers []wbyte.Pool
	timers  []timer.Timer
	options RoundOptions
}

// NewRound new a round struct.
func NewRound(c *conf.TCP) (r *Round) {
	var i int
	var p *conf.Protocol
	r = &Round{
		options: RoundOptions{
			Reader:       int(c.Reader),
			ReadBuf:      int(c.ReadBuf),
			ReadBufSize:  int(c.ReadBufSize),
			Writer:       int(c.Writer),
			WriteBuf:     int(c.WriteBuf),
			WriteBufSize: int(c.WriteBufSize),
			Timer:        int(p.Timer),
			TimerSize:    int(p.TimerSize),
		}}
	// reader
	r.readers = make([]wbyte.Pool, r.options.Reader)
	for i = 0; i < r.options.Reader; i++ {
		r.readers[i].Init(r.options.ReadBuf, r.options.ReadBufSize)
	}
	// writer
	r.writers = make([]wbyte.Pool, r.options.Writer)
	for i = 0; i < r.options.Writer; i++ {
		r.writers[i].Init(r.options.WriteBuf, r.options.WriteBufSize)
	}
	// timer
	r.timers = make([]timer.Timer, r.options.Timer)
	for i = 0; i < r.options.Timer; i++ {
		r.timers[i].Init(r.options.TimerSize)
	}
	return
}

// Timer get a timer.
func (r *Round) Timer(rn int) *timer.Timer {
	return &(r.timers[rn%r.options.Timer])
}

// Reader get a reader memory buffer.
func (r *Round) Reader(rn int) *wbyte.Pool {
	return &(r.readers[rn%r.options.Reader])
}

// Writer get a writer memory buffer pool.
func (r *Round) Writer(rn int) *wbyte.Pool {
	return &(r.writers[rn%r.options.Writer])
}
