//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package job

import (
	"sync"

	"github.com/uber-go/tally"
	"github.com/uber/aresdb/client"
	controllerCli "github.com/uber/aresdb/controller/client"
	"github.com/uber/aresdb/subscriber/common/consumer"
	"github.com/uber/aresdb/subscriber/common/message"
	"github.com/uber/aresdb/subscriber/common/rules"
	"github.com/uber/aresdb/subscriber/common/sink"
	"github.com/uber/aresdb/subscriber/common/tools"
	"github.com/uber/aresdb/subscriber/config"
)

// NewConsumer is the type of function each consumer that implements Consumer should provide for initialization.
type NewConsumer func(jobConfig *rules.JobConfig, serviceConfig config.ServiceConfig) (consumer.Consumer, error)

// NewDecoder is the type of function each decoder that implements decoder should provide for initialization.
type NewDecoder func(jobConfig *rules.JobConfig, serviceConfig config.ServiceConfig) (decoder message.Decoder, err error)

// NewSink is the type of function each decoder that implements sink should provide for initialization.
type NewSink func(
	serviceConfig config.ServiceConfig, jobConfig *rules.JobConfig, cluster string,
	sinkCfg config.SinkConfig, aresControllerClient controllerCli.ControllerClient) (sink.Sink, error)

// StreamingProcessor defines a individual processor that connects to a Kafka high level consumer,
// processes the messages based on the type of job and saves to database
type StreamingProcessor struct {
	ID                   int
	context              *ProcessorContext
	jobConfig            *rules.JobConfig
	cluster              string
	serviceConfig        config.ServiceConfig
	scope                tally.Scope
	aresControllerClient controllerCli.ControllerClient
	sink                 sink.Sink
	sinkInitFunc         NewSink
	highLevelConsumer    consumer.Consumer
	consumerInitFunc     NewConsumer
	parser               *message.Parser
	decoder              message.Decoder
	batcher              *tools.Batcher
	msgSizes             chan int64
	shutdown             chan bool
	close                chan bool
	errors               chan ProcessorError
	failureHandler       FailureHandler
}

// NewStreamingProcessor returns Processor to consume, process and save data to db.
func NewStreamingProcessor(id int, jobConfig *rules.JobConfig, aresControllerClient controllerCli.ControllerClient, sinkInitFunc NewSink, consumerInitFunc NewConsumer, decoderInitFunc NewDecoder,
	errors chan ProcessorError, msgSizes chan int64, serviceConfig config.ServiceConfig) (Processor, error) {
	_ = "STUB: not implemented"
	return *new(Processor), nil
}

// Initialize downstream DB

// initialize failure handler

// Initialize Kafka consumer

// Initialize the decoder based on topic

// Initialize message parser

func initFailureHandler(serviceConfig config.ServiceConfig,
	jobConfig *rules.JobConfig, db sink.Sink) FailureHandler {
	_ = "STUB: not implemented"
	return *new(FailureHandler)
}

// initDatabase will initialize the database for writing ingest data
func initSink(
	jobConfig *rules.JobConfig, serviceConfig config.ServiceConfig, aresControllerClient controllerCli.ControllerClient, sinkInitFunc NewSink) (sink.Sink, error) {
	_ = "STUB: not implemented"
	return *new(sink.Sink), nil
}

// GetID will return ID of this processor
func (s *StreamingProcessor) GetID() int {
	_ = "STUB: not implemented"

	// GetContext will return context of this processor
	return 0
}

func (s *StreamingProcessor) GetContext() *ProcessorContext {
	_ = "STUB: not implemented"

	// Stop will stop the processor
	return nil
}

func (s *StreamingProcessor) Stop() { _ = "STUB: not implemented"; return }

// Restart will stop the processor and start the process again in the case failure detected
func (s *StreamingProcessor) Restart() { _ = "STUB: not implemented"; return }

// not calling Stop() due to different flag setting

// wating for the original routine to stop

// wating for some time to avoid keep restarting in short of period, or quit if stop is called during restart

// restart

// quit restarting if failed to re-initialized database or consumer

func (s *StreamingProcessor) reInitialize() error {
	_ = "STUB: not implemented"
	// maybe we can try to re-initialize if anything failed
	return nil
}

// Initialize Kafka consumer

// Run will start the Processor that reads from high level kafka consumer,
// decodes the message and add the row to batcher for saving to ares.
func (s *StreamingProcessor) Run() { _ = "STUB: not implemented"; return }

// reset back the running flag

// Update message count in context

// log message size and report for throttling

// decode message and add to batcher for parse and save

// initBatcher will initialize the batcher with 1 worker per processor, so to
// maintain the order of offset commits
func (s *StreamingProcessor) initBatcher() { _ = "STUB: not implemented"; return }

// decodeMessage will decode the given Kafka message and return Message, which defines
// actual raw Kafka message, decoded message and timestamp of the message
func (s *StreamingProcessor) decodeMessage(msg consumer.Message) (*message.Message, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// saveToDestination will parse given decoded message based on transformations in JobConfig
// and save it to configured destination
func (s *StreamingProcessor) saveToDestination(batch []interface{}, destination sink.Destination) {
	_ = "STUB: not implemented"
	return
}

func (s *StreamingProcessor) writeRow(rows []client.Row, destination sink.Destination) {
	_ = "STUB: not implemented"
	return
}

// saveToDB will parse and save given batches
func (s *StreamingProcessor) saveToDB(batches chan []interface{}, wg *sync.WaitGroup) {
	_ = "STUB: not implemented"
	return
}

// reportMessageAge will report the message age for the message
func (s *StreamingProcessor) reportMessageAge(msg *message.Message) {
	_ = "STUB: not implemented"
	return
}
