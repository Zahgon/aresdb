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

package sink

import (
	"github.com/Shopify/sarama"
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/client"
	controllerCli "github.com/uber/aresdb/controller/client"
	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/subscriber/common/rules"
	"github.com/uber/aresdb/subscriber/config"
)

// default schema refresh interval in seconds
const defaultSchemaRefreshInterval = 600

type KafkaPublisher struct {
	sarama.SyncProducer
	client.UpsertBatchBuilder

	ServiceConfig config.ServiceConfig
	JobConfig     *rules.JobConfig
	Scope         tally.Scope
	ClusterName   string
}

func NewKafkaPublisher(serviceConfig config.ServiceConfig, jobConfig *rules.JobConfig, cluster string,
	sinkCfg config.SinkConfig, aresControllerClient controllerCli.ControllerClient) (Sink, error) {
	_ = "STUB: not implemented"
	return *new(Sink), nil
}

// Wait for all in-sync replicas to ack the message

// replace httpSchemaFetcher with gateway client
// httpSchemaFetcher := NewHttpSchemaFetcher(httpClient, cfg.Address, metricScope)

// schema refresh is based on job assignment refresh, so disable at here

// Shutdown will clean up resources that needs to be cleaned up
func (kp *KafkaPublisher) Shutdown() { _ = "STUB: not implemented"; return }

// Save saves a batch of row objects into a destination
func (kp *KafkaPublisher) Save(destination Destination, rows []client.Row) error {
	_ = "STUB: not implemented"
	return nil
}

// case1: no sharding --  publish rows to random kafka partition

// case2: sharding -- publish rows to specified partition

// Cluster returns the DB cluster name
func (kp *KafkaPublisher) Cluster() string { _ = "STUB: not implemented"; return "" }

func (kp *KafkaPublisher) buildKafkaMessage(msgs *[]*sarama.ProducerMessage, rowsIgnored *int, tableName string, shardID int32, columnNames []string, rows []client.Row,
	updateModes ...memCom.ColumnUpdateMode) {
	_ = "STUB: not implemented"
	return
}
