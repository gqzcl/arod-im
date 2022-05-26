// Copyright 2022 gqzcl <gqzcl@qq.com>. All rights reserved.
// Use of this source code is governed by a MIT style

package kafka

import (
	"github.com/segmentio/kafka-go"
	"time"
)

type ConfigOption func(o kafka.ReaderConfig)

// WithGroupId 配置readerConfig的groupId
func WithGroupId(groupId string) ConfigOption {
	return func(o kafka.ReaderConfig) {
		o.GroupID = groupId
	}
}

// WithPartition 配置readerConfig的partition
func WithPartition(partition int) ConfigOption {
	return func(o kafka.ReaderConfig) {
		o.Partition = partition
	}
}

func WithCommitInterval(commitInterval time.Duration) ConfigOption {
	return func(o kafka.ReaderConfig) {
		o.CommitInterval = commitInterval
	}
}
