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
	"github.com/uber/aresdb/memstore/common"
)

const maxBatchesPerFile = 5000
const commitInterval = 100

// kafkaRedoLogManager is kafka partition level consumer, also implementation of RedoLogManager
type kafkaRedoLogManager struct {
	sync.RWMutex

	Topic     string `json:"topic"`
	TableName string `json:"table"`
	Shard     int    `json:"shard"`

	// MaxEventTime per virtual redolog file
	MaxEventTimePerFile map[int64]uint32 `json:"maxEventTimePerFile"`
	// FirstKafkaOffset per virtual redolog file
	FirstKafkaOffsetPerFile map[int64]int64 `json:"firstKafkaOffsetPerFile"`
	SizePerFile             map[int64]int   `json:"sizePerFile"`
	TotalRedologSize        int             `json:"totalRedologSize"`

	// is this redolog manager also used for recovery (true if no disk redolog)
	includeRecovery bool

	consumer          sarama.Consumer
	partitionConsumer sarama.PartitionConsumer

	done chan struct{}

	commitFunc              func(string, int, int64) error
	checkPointFunc          func(string, int, int64) error
	getCommitOffsetFunc     func(string, int) (int64, error)
	getCheckpointOffsetFunc func(string, int) (int64, error)

	// used for external blocking check if recovery done
	recoveryChan chan bool
	recoveryDone bool
	// batch recovered counts
	batchRecovered int
	batchReceived  int
}

// newKafkaRedoLogManager creates kafka redolog manager
func newKafkaRedoLogManager(namespace, table, suffix string, shard int, consumer sarama.Consumer, includeRecovery bool,
	commitFunc func(string, int, int64) error,
	checkPointFunc func(string, int, int64) error,
	getCommitOffsetFunc func(string, int) (int64, error),
	getCheckpointOffsetFunc func(string, int) (int64, error)) *kafkaRedoLogManager {
	_ = "STUB: not implemented"
	return nil
}

// IsAppendEnabled returns whether appending is enabled
func (k *kafkaRedoLogManager) IsAppendEnabled() bool {
	_ = "STUB: not implemented"

	// AppendToRedoLog to record upsertbatch info as redolog
	return false
}

func (k *kafkaRedoLogManager) AppendToRedoLog(upsertBatch *common.UpsertBatch) (int64, uint32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (k *kafkaRedoLogManager) UpdateMaxEventTime(eventTime uint32, fileID int64) {
	_ = "STUB: not implemented"
	return
}

func (k *kafkaRedoLogManager) getFileOffset(kafkaOffset int64) (int64, uint32) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (k *kafkaRedoLogManager) CheckpointRedolog(eventTimeCutoff uint32, fileIDCheckpointed int64, batchOffset uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// file not purgeable

// fileID existing in MaxEventTimePerFile should always have entry in FirstKafkaOffsetPerFile

func (k *kafkaRedoLogManager) addMessage(fileID int64, kafkaOffset int64, size int) {
	_ = "STUB: not implemented"
	return
}

func (k *kafkaRedoLogManager) getKafkaOffsets() (int64, int64) {
	_ = "STUB: not implemented"
	return 0, 0
}

func (k *kafkaRedoLogManager) Iterator() (NextUpsertFunc, error) {
	_ = "STUB: not implemented"
	return *new(NextUpsertFunc), nil
}

// close previous created partition consumer

// partition consumer closed

// consumer closed

// consumer closed

func (k *kafkaRedoLogManager) setRecoveryDone() { _ = "STUB: not implemented"; return }

func (k *kafkaRedoLogManager) WaitForRecoveryDone() { _ = "STUB: not implemented"; return }

func (k *kafkaRedoLogManager) GetTotalSize() int { _ = "STUB: not implemented"; return 0 }

func (k *kafkaRedoLogManager) GetNumFiles() int { _ = "STUB: not implemented"; return 0 }

func (k *kafkaRedoLogManager) GetBatchReceived() int { _ = "STUB: not implemented"; return 0 }

func (k *kafkaRedoLogManager) GetBatchRecovered() int { _ = "STUB: not implemented"; return 0 }

func (k *kafkaRedoLogManager) Close() { _ = "STUB: not implemented"; return }

// MarshalJSON marshals a kafkaRedoLogManager into json.
func (k *kafkaRedoLogManager) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Avoid json.Marshal loop calls.
	return nil, nil
}
