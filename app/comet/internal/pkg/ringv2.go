package pkg

import (
	"arod-im/api/protocol"
	"arod-im/app/comet/internal/pkg/disruptor"
)

const (
	BufferSize   = 1024 * 64
	BufferMask   = BufferSize - 1
	Iterations   = 128 * 1024 * 32
	Reservations = 1
)

type RingBuffer struct {
	BufferSize uint64 // ring buffer size
	BufferMask uint64 // ring buffer mask
	rp         uint64 // read index
	wp         uint64 // write index
	buffers    []protocol.Proto
}

func (consumer RingBuffer) Consume(lower, upper int64) {
	for ; lower <= upper; lower++ {
		// message := ringBuffer[lower&BufferMask]
		// if message != lower {
		// 	panic(fmt.Errorf("race condition: %d %d", message, lower))
		// }
	}
}
func NewDisruptor() disruptor.Disruptor {
	return disruptor.New(
		disruptor.WithCapacity(BufferSize),
		disruptor.WithConsumerGroup(RingBuffer{}),
	)
}
func Write() {

}
func Read() {

}
