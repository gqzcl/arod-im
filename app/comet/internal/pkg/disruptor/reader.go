package disruptor

import (
	"io"
	"sync/atomic"
)

const (
	stateRunning = 0
	stateClosed  = 1
)

type DefaultReader struct {
	state    int64
	current  *Cursor // this reader has processed up to this sequence
	written  *Cursor // the ring buffer has been written up to this sequence
	upstream Barrier // all of the readers have advanced up to this sequence
	waiter   WaitStrategy
	consumer Consumer
}

func NewReader(current, written *Cursor, upstream Barrier, waiter WaitStrategy, consumer Consumer) *DefaultReader {
	return &DefaultReader{
		state:    stateRunning,
		current:  current,
		written:  written,
		upstream: upstream,
		waiter:   waiter,
		consumer: consumer,
	}
}

func (r *DefaultReader) Read() {
	var gateCount, idleCount, lower, upper int64
	var current = r.current.Load()

	for {
		lower = current + 1
		upper = r.upstream.Load()

		if lower <= upper {
			r.consumer.Consume(lower, upper)
			r.current.Store(upper)
			current = upper
		} else if upper = r.written.Load(); lower <= upper {
			gateCount++
			idleCount = 0
			r.waiter.Gate(gateCount)
		} else if atomic.LoadInt64(&r.state) == stateRunning {
			idleCount++
			gateCount = 0
			r.waiter.Idle(idleCount)
		} else {
			break
		}
	}

	if closer, ok := r.consumer.(io.Closer); ok {
		_ = closer.Close()
	}
}

func (r *DefaultReader) Close() error {
	atomic.StoreInt64(&r.state, stateClosed)
	return nil
}
