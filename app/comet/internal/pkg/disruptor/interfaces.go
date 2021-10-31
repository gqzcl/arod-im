package disruptor

import "errors"

var ErrMinimumReservationSize = errors.New("the minimum reservation size is 1 slot")

type Writer interface {
	Reserve(count int64) int64
	Commit(lower, upper int64)
}

type Reader interface {
	Read()
	Close() error
}

type Barrier interface {
	Load() int64
}

type Consumer interface {
	Consume(lower, upper int64)
}

type WaitStrategy interface {
	Gate(int64)
	Idle(int64)
}
