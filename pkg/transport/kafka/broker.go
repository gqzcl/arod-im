package kafka

import (
	"context"
)

type Consumer interface {
	Key() string
	Value() []byte
	Header() map[string]string
	Ack() error
	Nack() error
}

type Event interface {
	Topic() string
	Message() *Message
	Ack() error
	Error() error
}

type Receiver interface {
	Options()
	Topic() string
	Unsubscribe() error
}

type Broker interface {
	Read(ctx context.Context) (Message, error)
	Ack(msg Message) error
	Nack(msg Message) error
	Close() error

	Address() string
	Connect() error
	Disconnect() error
	//Subscribe(topic string, h Handler, opts ...SubscribeOption) (Receiver, error)
	Name() string
}
