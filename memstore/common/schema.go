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

package common

import (
	"sync"

	metaCom "github.com/uber/aresdb/metastore/common"
)

// TableSchema stores metadata of the table such as columns and primary keys.
// It also stores the dictionaries for enum columns.
type TableSchema struct {
	sync.RWMutex `json:"-"`
	// Main schema of the table. Mutable.
	Schema metaCom.Table `json:"schema"`
	// Maps from column names to their IDs. Mutable.
	ColumnIDs map[string]int `json:"columnIDs"`
	// Maps from enum column names to their case dictionaries. Mutable.
	EnumDicts map[string]EnumDict `json:"enumDicts"`
	// DataType for each column ordered by column ID. Mutable.
	ValueTypeByColumn []DataType `json:"valueTypeByColumn"`
	// Number of bytes in the primary key. Immutable.
	PrimaryKeyBytes int `json:"primaryKeyBytes"`
	// Types of each primary key column. Immutable.
	PrimaryKeyColumnTypes []DataType `json:"primaryKeyColumnTypes"`
	// Default values of each column. Mutable. Nil means default value is not set.
	DefaultValues []*DataValue `json:"-"`
}

// EnumDict contains mapping from and to enum strings to numbers.
type EnumDict struct {
	// Either 0x100 for small_enum, or 0x10000 for big_enum.
	Capacity    int            `json:"capacity"`
	Dict        map[string]int `json:"dict"`
	ReverseDict []string       `json:"reverseDict"`
}

// NewTableSchema creates a new table schema object from metaStore table object,
// this does not set enum cases.
func NewTableSchema(table *metaCom.Table) *TableSchema { _ = "STUB: not implemented"; return nil }

// MarshalJSON marshals TableSchema into json.
func (t *TableSchema) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	// Avoid loop json.Marshal calls.
	return nil, nil
}

// SetTable sets a updated table and update TableSchema,
// should acquire lock before calling.
func (t *TableSchema) SetTable(table *metaCom.Table) { _ = "STUB: not implemented"; return }

// SetDefaultValue parses the default value string if present and sets to TableSchema.
// Schema lock should be acquired and release by caller and enum dict should already be
// created/update before this function.
func (t *TableSchema) SetDefaultValue(columnID int) {
	_ = "STUB: not implemented"
	// Default values are already set.
	return
}

// Should no happen since the enum dict should already be created.

// Should no happen since the enum value should already be created.

// Should not happen since the string value is already validated by schema handler.

// createEnumDict creates the enum dictionary for the specified column with the
// specified initial cases, and attaches it to TableSchema object.
// Caller should acquire the schema lock before calling this function.
func (t *TableSchema) CreateEnumDict(columnName string, enumCases []string) {
	_ = "STUB: not implemented"
	return
}

// GetValueTypeByColumn makes a copy of the ValueTypeByColumn so callers don't have to hold a read
// lock to access it.
func (t *TableSchema) GetValueTypeByColumn() []DataType { _ = "STUB: not implemented"; return nil }

// GetPrimaryKeyColumns makes a copy of the Schema.PrimaryKeyColumns so callers don't have to hold
// a read lock to access it.
func (t *TableSchema) GetPrimaryKeyColumns() []int { _ = "STUB: not implemented"; return nil }

// GetColumnDeletions returns a boolean slice that indicates whether a column has been deleted. Callers
// need to hold a read lock.
func (t *TableSchema) GetColumnDeletions() []bool { _ = "STUB: not implemented"; return nil }

// GetColumnIfNonNilDefault returns a boolean slice that indicates whether a column has non nil default value. Callers
// need to hold a read lock.
func (t *TableSchema) GetColumnIfNonNilDefault() []bool { _ = "STUB: not implemented"; return nil }

// GetArchivingSortColumns makes a copy of the Schema.ArchivingSortColumns so
// callers don't have to hold a read lock to access it.
func (t *TableSchema) GetArchivingSortColumns() []int { _ = "STUB: not implemented"; return nil }
