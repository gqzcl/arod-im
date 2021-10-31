package disruptor

import "errors"

type Wireup struct {
	waiter         WaitStrategy
	capacity       int64
	consumerGroups [][]Consumer
}
type Option func(*Wireup)

func New(options ...Option) Disruptor {
	if this, err := NewWireup(options...); err != nil {
		panic(err)
	} else {
		return NewDisruptor(this.Build())
	}
}

func NewWireup(options ...Option) (*Wireup, error) {
	this := &Wireup{}

	WithWaitStrategy(NewWaitStrategy())(this)

	for _, option := range options {
		option(this)
	}

	if err := this.validate(); err != nil {
		return nil, err
	}

	return this, nil
}
func (w *Wireup) validate() error {
	if w.waiter == nil {
		return errMissingWaitStrategy
	}

	if w.capacity <= 0 {
		return errCapacityTooSmall
	}

	if w.capacity&(w.capacity-1) != 0 {
		return errCapacityPowerOfTwo
	}

	if len(w.consumerGroups) == 0 {
		return errMissingConsumers
	}

	for _, consumerGroup := range w.consumerGroups {
		if len(consumerGroup) == 0 {
			return errMissingConsumersInGroup
		}

		for _, consumer := range consumerGroup {
			if consumer == nil {
				return errEmptyConsumer
			}
		}
	}

	return nil
}

func (w *Wireup) Build() (Writer, Reader) {
	var writerSequence = NewCursor()
	readers, readBarrier := w.buildReaders(writerSequence)
	return NewWriter(writerSequence, readBarrier, w.capacity), compositeReader(readers)
}
func (w *Wireup) buildReaders(writerSequence *Cursor) (readers []Reader, upstream Barrier) {
	upstream = writerSequence

	for _, consumerGroup := range w.consumerGroups {
		var consumerGroupSequences []*Cursor

		for _, consumer := range consumerGroup {
			currentSequence := NewCursor()
			readers = append(readers, NewReader(currentSequence, writerSequence, upstream, w.waiter, consumer))
			consumerGroupSequences = append(consumerGroupSequences, currentSequence)
		}

		upstream = NewCompositeBarrier(consumerGroupSequences...)
	}

	return readers, upstream
}

func WithWaitStrategy(value WaitStrategy) Option { return func(this *Wireup) { this.waiter = value } }
func WithCapacity(value int64) Option            { return func(this *Wireup) { this.capacity = value } }
func WithConsumerGroup(value ...Consumer) Option {
	return func(this *Wireup) { this.consumerGroups = append(this.consumerGroups, value) }
}

var (
	errMissingWaitStrategy     = errors.New("a wait strategy must be provided")
	errCapacityTooSmall        = errors.New("the capacity must be at least 1")
	errCapacityPowerOfTwo      = errors.New("the capacity be a power of two, e.g. 2, 4, 8, 16")
	errMissingConsumers        = errors.New("no consumers have been provided")
	errMissingConsumersInGroup = errors.New("the consumer group does not have any consumers")
	errEmptyConsumer           = errors.New("an empty consumer was specified in the consumer group")
)
