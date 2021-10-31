package disruptor

import "math"

type compositeBarrier []*Cursor

func NewCompositeBarrier(sequences ...*Cursor) Barrier {
	if len(sequences) == 1 {
		return sequences[0]
	} else {
		return compositeBarrier(sequences)
	}
}

func (barrier compositeBarrier) Load() int64 {
	var minimum int64 = math.MaxInt64

	for _, item := range barrier {
		if sequence := item.Load(); sequence < minimum {
			minimum = sequence
		}
	}

	return minimum
}
