package disruptor

import "sync/atomic"

type Cursor [8]int64 // prevent false sharing of the sequence cursor by padding the CPU cache line with 64 *bytes* of data.

const defaultCursorValue = -1

func NewCursor() (cursor *Cursor) {
	cursor[0] = defaultCursorValue
	return
}

func (cur *Cursor) Store(value int64) { atomic.StoreInt64(&cur[0], value) }
func (cur *Cursor) Load() int64       { return atomic.LoadInt64(&cur[0]) }
