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

package redolog

import (
	"sync"

	"github.com/Shopify/sarama"
	"github.com/uber/aresdb/common"
	"github.com/uber/aresdb/diskstore"
	metaCom "github.com/uber/aresdb/metastore/common"
)

// Master class to create shard level redolog manager
type RedoLogManagerMaster struct {
	sync.Mutex
	// namespace of the datanode
	Namespace string
	// redolog config
	RedoLogConfig *common.RedoLogConfig
	// kafka consuer if kafka consumer is configured
	consumer sarama.Consumer
	// DiskStore
	diskStore diskstore.DiskStore
	// Metastore
	metaStore metaCom.MetaStore
	// save all table partition redolog managers
	managers map[string]map[int]RedologManager
}

// NewRedoLogManagerMaster create RedoLogManagerMaster instance
func NewRedoLogManagerMaster(namespace string, c *common.RedoLogConfig, diskStore diskstore.DiskStore, metaStore metaCom.MetaStore) (*RedoLogManagerMaster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewKafkaRedoLogManagerMaster convenient function if the kafka consumer can be passed in from outside
func NewKafkaRedoLogManagerMaster(namespace string, cfg *common.RedoLogConfig, diskStore diskstore.DiskStore, metaStore metaCom.MetaStore, consumer sarama.Consumer) (*RedoLogManagerMaster, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewRedologManager create compositeRedoLogManager on specified table/shard
// each table/shard should only have one compositeRedoLogManager
func (m *RedoLogManagerMaster) NewRedologManager(table string, shard int, unsharded bool, tableConfig *metaCom.TableConfig) (RedologManager, error) {
	_ = "STUB: not implemented"
	return *new(RedologManager), nil
}

// Close one table shard Redolog manager
func (m *RedoLogManagerMaster) Close(table string, shard int) { _ = "STUB: not implemented"; return }

// Stop close all shard redolog manager and kafka consumer
func (m *RedoLogManagerMaster) Stop() { _ = "STUB: not implemented"; return }
