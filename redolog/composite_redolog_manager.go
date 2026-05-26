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
	"github.com/uber/aresdb/diskstore"
	"github.com/uber/aresdb/memstore/common"
	metaCom "github.com/uber/aresdb/metastore/common"
)

// compositeRedologManager is the class to take data ingestion from all data source (kafka, http, etc.), write to local redolog when necessary,
// and then call storeFunc to store data
type compositeRedoLogManager struct {
	sync.RWMutex `json:"-"`
	// table name
	Table string `json:"table"`
	// table shard id
	Shard int `json:"shard"`
	// Local file redolog manager
	fileRedoLogManager *FileRedoLogManager
	// Kafka consumer if kafka import is supported
	kafkaRedoLogManager *kafkaRedoLogManager
}

// NewCompositeRedoLogManager create compositeRedoLogManager oibject
func newCompositeRedoLogManager(namespace, table, suffix string, shard int, tableConfig *metaCom.TableConfig,
	consumer sarama.Consumer, diskStore diskstore.DiskStore,
	commitFunc func(string, int, int64) error,
	checkPointFunc func(string, int, int64) error,
	getCommitOffsetFunc func(string, int) (int64, error),
	getCheckpointOffsetFunc func(string, int) (int64, error)) *compositeRedoLogManager {
	_ = "STUB: not implemented"
	return nil
}

// Iterator walk through redolog batch from both file and kafka
func (s *compositeRedoLogManager) Iterator() (NextUpsertFunc, error) {
	_ = "STUB: not implemented"
	return *new(NextUpsertFunc), nil
}

// WaitForRecoveryDone block call to wait for recovery finish
func (s *compositeRedoLogManager) WaitForRecoveryDone() { _ = "STUB: not implemented"; return }

// IsAppendEnabled returns whether appending is enabled
func (s *compositeRedoLogManager) IsAppendEnabled() bool {
	_ = "STUB: not implemented"

	// AppendToRedoLog append upsert batch into redolog file or commit offset
	return false
}

func (s *compositeRedoLogManager) AppendToRedoLog(upsertBatch *common.UpsertBatch) (int64, uint32) {
	_ = "STUB: not implemented"
	return 0, 0
}

// UpdateMaxEventTime update max event time for related redolog file
func (s *compositeRedoLogManager) UpdateMaxEventTime(eventTime uint32, redoFile int64) {
	_ = "STUB: not implemented"
	return
}

// CheckpointRedolog clean up obsolete redolog files and save checkpoint offset
func (s *compositeRedoLogManager) CheckpointRedolog(cutoff uint32, redoFileCheckpointed int64, batchOffset uint32) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *compositeRedoLogManager) GetTotalSize() int { _ = "STUB: not implemented"; return 0 }

func (s *compositeRedoLogManager) GetNumFiles() int { _ = "STUB: not implemented"; return 0 }

func (s *compositeRedoLogManager) GetBatchReceived() int { _ = "STUB: not implemented"; return 0 }

func (s *compositeRedoLogManager) GetBatchRecovered() int { _ = "STUB: not implemented"; return 0 }

func (s *compositeRedoLogManager) Close() { _ = "STUB: not implemented"; return }

// MarshalJSON marshals a fileRedologManager into json.
func (s *compositeRedoLogManager) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Avoid json.Marshal loop calls.
	return nil, nil
}
