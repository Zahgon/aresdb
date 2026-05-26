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

package client

import (
	"net/http"
	"sync"

	"github.com/uber-go/tally"
	memCom "github.com/uber/aresdb/memstore/common"
	metaCom "github.com/uber/aresdb/metastore/common"
	"go.uber.org/zap"
)

const (
	// default request time out in seconds
	defaultRequestTimeout = 120
	// default schema refresh interval in seconds
	defaultSchemaRefreshInterval = 600
	dataIngestionHeader          = "application/upsert-data"
	applicationJSONHeader        = "application/json"
	defaultStringEnumLength      = 1024
)

// Row represents a row of insert data.
type Row []interface{}

// Connector is the connector interface for ares.
type Connector interface {
	// Insert inserts rows to ares
	// returns number of rows inserted and error.
	// updateModes are optional, if ignored for all columns, no need to set
	// if set, then all columns needs to be set
	Insert(tableName string, columnNames []string, rows []Row, updateModes ...memCom.ColumnUpdateMode) (int, error)
	// Close the connection
	Close()
}

// UpsertBatchBuilder is an interface of upsertBatch on client side
type UpsertBatchBuilder interface {
	PrepareUpsertBatch(tableName string, columnNames []string, updateModes []memCom.ColumnUpdateMode, rows []Row) ([]byte, int, error)
}

// enumCasesWrapper is a response/request body which wraps enum cases
type enumCasesWrapper struct {
	EnumCases []string
}

type TableSchema struct {
	Table *metaCom.Table
	// maps from column name to columnID for convenience
	ColumnDict map[string]int
}

// enumDict maps from enum value to enumID
type enumDict map[string]int

// UpsertBatchBuilderImpl implements interface UpsertBatchBuilder
type UpsertBatchBuilderImpl struct {
	sync.RWMutex

	logger        *zap.SugaredLogger
	metricScope   tally.Scope
	schemaHandler *CachedSchemaHandler
}

// connector is the ares connector implementation
type connector struct {
	cfg                ConnectorConfig
	httpClient         http.Client
	upsertBatchBuilder UpsertBatchBuilder
	schemaHandler      *CachedSchemaHandler
}

// ConnectorConfig holds the configurations for ares Connector.
type ConnectorConfig struct {
	// Address is in the format of host:port
	Address string `yaml:"address" json:"address"`
	// DeviceChoosingTimeout value is the request timeout in seconds for http calls
	// if <= 0, will use default
	Timeout int `yaml:"timeout" json:"timeout"`
	// SchemaRefreshInterval is the interval in seconds for the connector to
	// fetch and refresh schema from ares
	// if <= 0, will use default
	SchemaRefreshInterval int `yaml:"schemaRefreshInterval" json:"schemaRefreshInterval"`
}

func NewUpsertBatchBuilderImpl(logger *zap.SugaredLogger, scope tally.Scope, schemaHandler *CachedSchemaHandler) UpsertBatchBuilder {
	_ = "STUB: not implemented"
	return *new(UpsertBatchBuilder)
}

// NewConnector returns a new ares Connector
func (cfg ConnectorConfig) NewConnector(logger *zap.SugaredLogger, metricScope tally.Scope) Connector {
	_ = "STUB: not implemented"
	return *new(Connector)
}

// Insert inserts a batch of rows into ares
func (c *connector) Insert(tableName string, columnNames []string, rows []Row, updateModes ...memCom.ColumnUpdateMode) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// if no update modes at all, use default

// Do nothing when there is no row to insert

//TODO: currently always use shard zero for single instance version

//TODO: break status code check and error check into two parts for more specific handling like retrying on 5xx

// Close the connection
func (c *connector) Close() { _ = "STUB: not implemented"; return }

// computeHLLValue populate hyperloglog value
func computeHLLValue(dataType memCom.DataType, value interface{}) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// prepareUpsertBatch prepares the upsert batch for upsert,
// returns upsertBatch byte array, number of rows in upsert batch and error.
func (c *connector) prepareUpsertBatch(tableName string, columnNames []string, updateModes []memCom.ColumnUpdateMode, rows []Row) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// return error if primary key is missing

// return error if time column is missing

// checkPrimaryKeys checks whether primary key is missing
func checkPrimaryKeys(schema *TableSchema, columnNames []string) error {
	_ = "STUB: not implemented"
	return nil
}

// checkTimeColumnExistence checks if time column is missing for fact table
func checkTimeColumnExistence(schema *TableSchema, columnNames []string) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *connector) dataPath(tableName string, shard int) string {
	_ = "STUB: not implemented"
	return ""
}

func (u *UpsertBatchBuilderImpl) prepareEnumCases(isEnumArrayCol bool, tableName, columnName string, colIndex, columnID int, rows []Row, abandonRows map[int]struct{}, caseInsensitive bool, disableAutoExpand bool) error {
	_ = "STUB: not implemented"
	return nil
}

// PrepareUpsertBatch prepares the upsert batch for upsert,
// returns upsertBatch byte array, number of rows in upsert batch and error.
func (u *UpsertBatchBuilderImpl) PrepareUpsertBatch(tableName string, columnNames []string,
	updateModes []memCom.ColumnUpdateMode, rows []Row) ([]byte, int, error) {
	_ = "STUB: not implemented"
	return nil, 0, nil
}

// use abandonRows to record abandoned row index due to invalid data

// following conditions only overwrite is supported:
// 1. dimension table (TODO: might support min/max in the future if needed)
// 2. primary key column
// 3. archiving sort column
// 4. data type not in uint8, int8, uint16, int16, uint32, int32, float32

// prevent primary key being nil

// skip rows if time column is nil for fact table

// no error handling here as it should already be covered in prepareEnumCases

// If enum value is not found from predefined enum cases and default value is not set, we set it to nil.

// Set value to the last row.
// compute hll value to insert

// here use original column data type to compute hll value

// directly insert value
