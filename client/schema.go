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
	metaCom "github.com/uber/aresdb/metastore/common"
	"go.uber.org/zap"
)

// SchemaFetcher is the interface for fetch schema and enums
type SchemaFetcher interface {
	// FetchAllSchemas fetches all schemas
	FetchAllSchemas() ([]*metaCom.Table, error)
	// FetchSchema fetch one schema for given table
	FetchSchema(table string) (*metaCom.Table, error)
	// FetchAllEnums fetches all enums for given table and column
	FetchAllEnums(tableName string, columnName string) ([]string, error)
	// ExtendEnumCases extends enum cases to given table column
	ExtendEnumCases(tableName, columnName string, enumCases []string) ([]int, error)
}

// httpSchemaFetcher is a http based schema fetcher
type httpSchemaFetcher struct {
	httpClient  http.Client
	metricScope tally.Scope
	address     string
}

// CachedSchemaHandler handles schema and enum requests with cache
type CachedSchemaHandler struct {
	*sync.RWMutex

	logger        *zap.SugaredLogger
	metricScope   tally.Scope
	schemaFetcher SchemaFetcher

	// mapping from table name to table schema
	schemas map[string]*TableSchema
	// map from table to columnID to enum dictionary
	// use columnID instead of name since column name can be reused
	// table names can be reused as well, deleting and adding a new table
	// will anyway requires job restart
	enumMappings map[string]map[int]enumDict

	// map from table to columnID to default enum id. Initialized during bootstrap
	// and will be set only if default value is non nil.
	enumDefaultValueMappings map[string]map[int]int
}

// NewCachedSchemaHandler creates a new cached schema handler
func NewCachedSchemaHandler(logger *zap.SugaredLogger, scope tally.Scope, schamaFetcher SchemaFetcher) *CachedSchemaHandler {
	_ = "STUB: not implemented"
	return nil
}

// NewHttpSchemaFetcher creates a new http schema fetcher
func NewHttpSchemaFetcher(httpClient http.Client, address string, scope tally.Scope) SchemaFetcher {
	_ = "STUB: not implemented"
	return *new(SchemaFetcher)
}

// Start starts the CachedSchemaHandler, if interval > 0, will start periodical refresh
func (cf *CachedSchemaHandler) Start(interval int) { _ = "STUB: not implemented"; return }

// TranslateEnum translates given enum value to its enumID
func (cf *CachedSchemaHandler) TranslateEnum(tableName string, columnID int, value interface{}, caseInsensitive bool) (enumID int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// here it already make sure the enum dictionary exists in cache

// FetchAllSchema fetch all schemas
func (cf *CachedSchemaHandler) FetchAllSchema() error { _ = "STUB: not implemented"; return nil }

// FetchSchemas fetch schemas in schemas of CachedSchemaHandler
func (cf *CachedSchemaHandler) FetchSchemas() { _ = "STUB: not implemented"; return }

// FetchSchema fetchs the schema of given table name
func (cf *CachedSchemaHandler) FetchSchema(tableName string) (*TableSchema, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// PrepareEnumCases prepares enum cases
func (cf *CachedSchemaHandler) PrepareEnumCases(tableName, columnName string, enumCases []string) error {
	_ = "STUB: not implemented"
	return nil
}

// It's recommended to set up elk or sentry logging to catch this error.

func (cf *CachedSchemaHandler) fetchAndSetEnumCases(table *metaCom.Table) error {
	_ = "STUB: not implemented"
	return nil
}

// Convert to lower case for comparison during ingestion.

// all mapping should be pre created

func (cf *CachedSchemaHandler) setTable(table *metaCom.Table) *TableSchema {
	_ = "STUB: not implemented"
	return nil
}

func (hf *httpSchemaFetcher) FetchAllEnums(tableName, columnName string) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hf *httpSchemaFetcher) ExtendEnumCases(tableName, columnName string, enumCases []string) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hf *httpSchemaFetcher) FetchSchema(tableName string) (*metaCom.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hf *httpSchemaFetcher) FetchAllSchemas() ([]*metaCom.Table, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (hf *httpSchemaFetcher) readJSONResponse(response *http.Response, err error, data interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (hf *httpSchemaFetcher) tablePath(tableName string) string {
	_ = "STUB: not implemented"
	return ""
}

func (hf *httpSchemaFetcher) listTablesPath() string { _ = "STUB: not implemented"; return "" }

func (hf *httpSchemaFetcher) enumDictPath(tableName, columnName string) string {
	_ = "STUB: not implemented"
	return ""
}
