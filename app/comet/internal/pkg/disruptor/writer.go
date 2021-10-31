package disruptor

import "runtime"

const SpinMask = 1024*16 - 1 // arbitrary; we'll want to experiment with different values

type DefaultWriter struct {
	written  *Cursor // the ring buffer has been written up to this sequence
	upstream Barrier // all of the readers have advanced up to this sequence
	capacity int64
	previous int64
	gate     int64
}

func NewWriter(written *Cursor, upstream Barrier, capacity int64) *DefaultWriter {
	return &DefaultWriter{
		upstream: upstream,
		written:  written,
		capacity: capacity,
		previous: defaultCursorValue,
		gate:     defaultCursorValue,
	}
}

// 存储
func (w *DefaultWriter) Reserve(count int64) int64 {
	if count <= 0 {
		panic(ErrMinimumReservationSize)
	}

	w.previous += count
	for spin := int64(0); w.previous-w.capacity > w.gate; spin++ {
		if spin&SpinMask == 0 {
			runtime.Gosched() // LockSupport.parkNanos(1L); http://bit.ly/1xiDINZ
		}

		w.gate = w.upstream.Load()
	}
	return w.previous
}

func (w *DefaultWriter) Commit(_, upper int64) { w.written.Store(upper) }
