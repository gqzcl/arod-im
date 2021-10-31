package disruptor

import "time"

type DefaultWaitStrategy struct{}

func NewWaitStrategy() DefaultWaitStrategy { return DefaultWaitStrategy{} }

// sleep 1ns
func (w DefaultWaitStrategy) Gate(count int64) { time.Sleep(time.Nanosecond) }

// sleep 50µs
func (w DefaultWaitStrategy) Idle(count int64) { time.Sleep(time.Microsecond * 50) }
