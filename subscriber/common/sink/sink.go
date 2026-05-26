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
	"github.com/uber/aresdb/client"
	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/subscriber/common/rules"
)

// Sink is abstraction for interactions with downstream storage layer
type Sink interface {
	// Cluster returns the DB cluster name
	Cluster() string

	// Save will save the rows into underlying database
	Save(destination Destination, rows []client.Row) error

	// Shutdown will close the connections to the database
	Shutdown()
}

// Destination contains the table and columns that each job is storing data into
// also records the behavior when encountering key errors
type Destination struct {
	// Table is table name
	Table string
	// ColumnNames are the list of column names after sorted
	ColumnNames []string
	// PrimaryKeys maps primary key columnName to its columnID after sorted
	PrimaryKeys map[string]int
	// PrimaryKeysInSchema maps primary key columnName to its columnID defined in schema
	PrimaryKeysInSchema map[string]int
	// AresUpdateModes defines update modes
	AresUpdateModes []memCom.ColumnUpdateMode
	// NumShards is the number of shards in the aresDB cluster
	NumShards uint32
}

func Shard(rows []client.Row, destination Destination, jobConfig *rules.JobConfig) (map[uint32][]client.Row, int) {
	_ = "STUB: not implemented"
	return nil, 0
}

// in this case, there is no sharding in this aresDB cluster

// convert primaryKey to byte array

// calculate shard

func shardFn(key []byte, numShards uint32) uint32 { _ = "STUB: not implemented"; return 0 }

func getPrimaryKeyBytes(row client.Row, destination Destination, jobConfig *rules.JobConfig, keyLength int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// create empty key with keyLength capacity

// convert the string to bytes if primaryKey value is string
