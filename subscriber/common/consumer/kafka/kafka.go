//	Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package kafka

import (
	"context"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/subscriber/common/consumer"
	"github.com/uber/aresdb/subscriber/common/rules"
	"github.com/uber/aresdb/subscriber/config"
	"go.uber.org/zap"
)

// KafkaConsumer implements Consumer interface
type KafkaConsumer struct {
	sarama.ConsumerGroup
	*sarama.Config
	sync.Mutex

	group      string
	topicArray []string
	logger     *zap.Logger
	scope      tally.Scope
	msgCh      chan consumer.Message

	// WARNING: The following channels should not be closed by the lib users
	closeAttempted bool
	closeCh        chan struct{}
}

// KafkaMessage implements Message interface
type KafkaMessage struct {
	*sarama.ConsumerMessage

	consumer    consumer.Consumer
	clusterName string
	session     sarama.ConsumerGroupSession
}

// CGHandler represents a Sarama consumer group handler
type CGHandler struct {
	consumer       *KafkaConsumer
	msgCounter     map[string]map[int32]tally.Counter
	msgByteCounter map[string]map[int32]tally.Counter
	msgOffsetGauge map[string]map[int32]tally.Gauge
	msgLagGauge    map[string]map[int32]tally.Gauge
}

// GetConsumerGroupName will return the consumer group name to use or being used
// for given deployment and job name
func GetConsumerGroupName(deployment, jobName string, aresCluster string) string {
	_ = "STUB: not implemented"
	return ""
}

func getKafkaVersion(v string) sarama.KafkaVersion {
	_ = "STUB: not implemented"
	return *new(sarama.KafkaVersion)
}

// NewKafkaConsumer creates kafka consumer
func NewKafkaConsumer(jobConfig *rules.JobConfig, serviceConfig config.ServiceConfig) (consumer.Consumer, error) {
	_ = "STUB: not implemented"
	return *new(consumer.Consumer), nil
}

// Name returns the name of this consumer group.
func (c *KafkaConsumer) Name() string {
	_ = "STUB: not implemented"

	// Topics returns the names of the topics being consumed.
	return ""
}

func (c *KafkaConsumer) Topics() []string { _ = "STUB: not implemented"; return nil }

// Errors returns a channel of errors for the topic. To prevent deadlocks,
// users must read from the error channel.
//
// All errors returned from this channel can be safely cast to the
// consumer.Error interface, which allows structured access to the topic
// name and partition number.
func (c *KafkaConsumer) Errors() <-chan error { _ = "STUB: not implemented"; return nil }

// Closed returns a channel that unblocks when the consumer successfully shuts
// down.
func (c *KafkaConsumer) Closed() <-chan struct{} {
	_ = "STUB: not implemented"

	// SetClosed is used for testing
	return nil
}

func (c *KafkaConsumer) SetClosed(closeCh chan struct{}) { _ = "STUB: not implemented"; return }

// Messages returns a channel of messages for the topic.
//
// If the consumer is not configured with nonzero buffer size, the Errors()
// channel must be read in conjunction with Messages() to prevent deadlocks.
func (c *KafkaConsumer) Messages() <-chan consumer.Message {
	_ = "STUB: not implemented"

	// SetMessages is used for testing
	return nil
}

func (c *KafkaConsumer) SetMessages(msgCh chan consumer.Message) {
	_ = "STUB: not implemented"

	// CommitUpTo marks this message and all previous messages in the same partition
	// as processed. The last processed offset for each partition is periodically
	// flushed to ZooKeeper; on startup, consumers begin processing after the last
	// stored offset.
	return
}

func (c *KafkaConsumer) CommitUpTo(msg consumer.Message) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *KafkaConsumer) startConsuming(ctx context.Context, cgHandler *CGHandler) {
	_ = "STUB: not implemented"
	return
}

// those four Metrics are of the format {"<topic name>":{<partition id>: <offset>, ...}, ...}

// initialize counter map

// check if context was cancelled, signaling that the consumer should stop

func (c *KafkaConsumer) processMsg(msg *sarama.ConsumerMessage, cgHandler *CGHandler,
	highWaterOffset int64, session sarama.ConsumerGroupSession) {
	_ = "STUB: not implemented"
	return
}

func (c *KafkaConsumer) Close() error { _ = "STUB: not implemented"; return nil }

func (m *KafkaMessage) Key() []byte { _ = "STUB: not implemented"; return nil }

func (m *KafkaMessage) Value() []byte { _ = "STUB: not implemented"; return nil }

func (m *KafkaMessage) Topic() string { _ = "STUB: not implemented"; return "" }

func (m *KafkaMessage) Partition() int32 { _ = "STUB: not implemented"; return 0 }

func (m *KafkaMessage) Offset() int64 { _ = "STUB: not implemented"; return 0 }

func (m *KafkaMessage) Ack() { _ = "STUB: not implemented"; return }

func (m *KafkaMessage) Nack() {
	_ = "STUB: not implemented"
	// No op for now since Kafka based DLQ is not implemented
	return
}

func (m *KafkaMessage) Cluster() string { _ = "STUB: not implemented"; return "" }

// Setup is run at the beginning of a new session, before ConsumeClaim
func (h *CGHandler) Setup(sarama.ConsumerGroupSession) error {
	_ = "STUB: not implemented"

	// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
	return nil
}

func (h *CGHandler) Cleanup(sarama.ConsumerGroupSession) error {
	_ = "STUB: not implemented"

	// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
	return nil
}

func (h *CGHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	_ = "STUB: not implemented"

	// NOTE:
	// Do not move the code below to a goroutine.
	// The `ConsumeClaim` itself is called within a goroutine, see:
	// https://github.com/Shopify/sarama/blob/master/consumer_group.go#L27-L29
	return nil
}
