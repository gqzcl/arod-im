package pkg

import (
	"arod-im/api/protocol"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/pkg/errors"
)

var (
	ErrRingEmpty = errors.New("ring buffer empty")
	ErrRingFull  = errors.New("ring buffer full")
)

// Ring ring proto buffer.
type Ring struct {
	num     uint64 // ring buffer size
	mask    uint64 // ring buffer mask
	rp      uint64 // read index
	wp      uint64 // write index
	buffers []protocol.Proto
}

// NewRing new a ring buffer.
// 创建一个ring队列，长度为最小大于num的2^n
func NewRing(num int, logger log.Logger) *Ring {
	r := new(Ring)
	r.Init(uint64(num))
	return r
}

// 初始化长度为最小大于num的2^n
func (r *Ring) Init(num uint64) {
	// 2^N
	if num&(num-1) != 0 {
		for num&(num-1) != 0 {
			num &= num - 1
		}
		num <<= 1
	}
	r.buffers = make([]protocol.Proto, num)
	r.num = num
	r.mask = r.num - 1
}

// Read a proto from ringbuffer.
func (r *Ring) Read() (proto *protocol.Proto, err error) {
	if r.rp == r.wp {
		return nil, ErrRingEmpty
	}
	proto = &r.buffers[r.rp&r.mask]
	return
}

// GetAdv incr read index.
func (r *Ring) ReadAdv() {
	r.rp++
}

// Write a proto to ringbuffer.
func (r *Ring) Write() (proto *protocol.Proto, err error) {
	// wp 和 rp为uint64保证一定成立
	if r.wp-r.rp >= r.num {
		return nil, ErrRingFull
	}
	proto = &r.buffers[r.wp&r.mask]
	return
}

// SetAdv incr write index.
func (r *Ring) WriteAdv() {
	r.wp++
}

// Reset reset ring.
func (r *Ring) Reset() {
	r.rp = 0
	r.wp = 0
}
