package dao

import (
	"context"
	"strconv"

	"github.com/golang/glog"
	pb "github.com/gqzcl/gim/api/logic"
	"google.golang.org/protobuf/proto"
	"gopkg.in/Shopify/sarama.v1"
)

// PushMsg 将消息推送到数据总线。
func (d *Dao) PushMsg(c context.Context, op int32, server string, keys []string, msg []byte) error {
	pushMsg := &pb.PushMsg{
		Type:      pb.PushMsg_PUSH,
		Operation: op,
		Server:    server,
		Keys:      keys,
		Msg:       msg,
	}
	// 序列化pushMsg
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return err
	}
	// ProducerMessage是传递给生产者以发送消息的元素集合。
	m := &sarama.ProducerMessage{
		// StringEncoder实现Go字符串的编码器接口，以便它们可以用作ProducerMessage中的键或值。
		// Key 此消息的分区键。现有的编码器包括StringEncoder和ByteEncoder。
		Key:   sarama.StringEncoder(keys[0]),
		Topic: d.c.Kafka.Topic,
		// 要存储在kafka中的实际消息。现有的编码器包括StringEncoder和ByteEncoder。
		Value: sarama.ByteEncoder(b),
	}
	// SendMessage生成给定的消息，并且仅在成功生成或未能生成时返回。它将返回所生成消息的分区和偏移量，如果消息未能生成，则返回错误。
	if _, _, err = d.kafkaPub.SendMessage(m); err != nil {
		glog.Errorf("PushMsg.send(push pushMsg:%v) error(%v)", pushMsg, err)
	}
	return err
}

// BroadcastRoomMsg将消息推送到数据总线。
func (d *Dao) BroadcastRoomMsg(c context.Context, op int32, room string, msg []byte) error {
	PushMsg := &pb.PushMsg{
		Type:      pb.PushMsg_ROOM,
		Operation: op,
		Room:      room,
		Msg:       msg,
	}
	b, err := proto.Marshal(PushMsg)
	if err != nil {
		return err
	}
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(room),
		Topic: d.c.Kafka.Topic,
		Value: sarama.ByteEncoder(b),
	}
	if _, _, err = d.kafkaPub.SendMessage(m); err != nil {
		glog.Errorf("PushMsg.send(broadcast_room pushMsg:%v) error(%v)", PushMsg, err)
	}
	return err
}

// BroadcastMsg将消息推送到数据总线。
func (d *Dao) BroadcastMsg(c context.Context, op, speed int32, msg []byte) error {
	pushMsg := &pb.PushMsg{
		Type:      pb.PushMsg_BROADCAST,
		Operation: op,
		Speed:     speed,
		Msg:       msg,
	}
	b, err := proto.Marshal(pushMsg)
	if err != nil {
		return err
	}
	m := &sarama.ProducerMessage{
		Key:   sarama.StringEncoder(strconv.FormatInt(int64(op), 10)),
		Topic: d.c.Kafka.Topic,
		Value: sarama.ByteEncoder(b),
	}
	if _, _, err = d.kafkaPub.SendMessage(m); err != nil {
		glog.Errorf("PushMsg.send(broadcast pushMsg:%v) error(%v)", pushMsg, err)
	}
	return err
}
