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
	"unsafe"

	"github.com/uber/aresdb/utils"
)

// ColumnUpdateMode represents how to update data from UpsertBatch
type ColumnUpdateMode int

// UpsertBatchVersion represents the version of upsert batch
type UpsertBatchVersion uint32

const (
	// UpdateOverwriteNotNull (default) will overwrite existing value if new value is NOT null, otherwise just skip
	UpdateOverwriteNotNull ColumnUpdateMode = iota
	// UpdateForceOverwrite will simply overwrite existing value even when new data is null
	UpdateForceOverwrite
	// UpdateWithAddition will add the existing value with new value if new value is not null, existing null value will be treated as 0 in Funculation
	UpdateWithAddition
	// UpdateWithMin will save the minimum of existing and new value if new value is not null, existing null value will be treated as MAX_INT in Funculation
	UpdateWithMin
	// UpdateWithMax will save the maximum of existing and new value if new value is not null, existing null value will be treated as MIN_INT in Funculation
	UpdateWithMax
	// MaxColumnUpdateMode is the current upper limit for column update modes
	MaxColumnUpdateMode
)

const (
	V1 UpsertBatchVersion = 0xFEED0001
)

type columnBuilder struct {
	columnID       int
	dataType       DataType
	values         []interface{}
	numValidValues int
	updateMode     ColumnUpdateMode
	isTimeColumn   bool
}

// SetValue write a value into the column at given row.
func (c *columnBuilder) SetValue(row int, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

// force time column value is uint32

// AddRow grow the value array by 1.
func (c *columnBuilder) AddRow() { _ = "STUB: not implemented"; return }

// AddRow shrink the value array by 1.
func (c *columnBuilder) RemoveRow() { _ = "STUB: not implemented"; return }

// ResetRows reset the row count to 0.
func (c *columnBuilder) ResetRows() { _ = "STUB: not implemented"; return }

// Calculated BufferSize returns the size of the column data in serialized format.
func (c *columnBuilder) CalculateBufferSize(offset *int) { _ = "STUB: not implemented"; return }

// if golang memory or array type, align to 4 bytes for offset vector

// 1. uint32 for each offset value, and length = numRows + 1
// 2. last offset value is the end offset of the offset buffer

// Padding size for value vector

// Padding size for value vector

// fixed value size

// AppendToBuffer writes the column data to buffer and advances offset.
func (c *columnBuilder) AppendToBuffer(writer *utils.BufferWriter) error {
	_ = "STUB: not implemented"
	return nil

	// both gotype and array type is variable length
}

// only non goType needs to write null vector

// variable length data type needs to write offsetVector

// Padding to 4 byte alignment for offset vector

// skip offset bytes

// Padding to 8 byte alignment for value vector

// local byte offset of current value in value vector

// write values starting from current value vector offset

// write current offset if offsetWriter is defined

// Handle null value.

// only skip bits when there is no offset vector

// advance current offset

// advance current offset

// lastly write the final offset into offsetWriter

// Align at byte for bit values.

// GetMode get the mode based on number of valid values.
func (c *columnBuilder) GetMode() ColumnMode { _ = "STUB: not implemented"; return *new(ColumnMode) }

// UpsertBatchBuilder is the builder for constructing an UpsertBatch buffer. It allows random value
// write at (row, col).
type UpsertBatchBuilder struct {
	NumRows     int
	columns     []*columnBuilder
	isFactTable bool
}

// NewUpsertBatchBuilder creates a new builder for constructing an UpersetBatch.
func NewUpsertBatchBuilder() *UpsertBatchBuilder { _ = "STUB: not implemented"; return nil }

// AddColumn add a new column to the builder. Initially, new columns have all values set to null.
func (u *UpsertBatchBuilder) AddColumn(columnID int, dataType DataType) error {
	_ = "STUB: not implemented"
	return nil
}

// AddColumnWithUpdateMode add a new column to the builder with update mode info. Initially, new columns have all values set to null.
func (u *UpsertBatchBuilder) AddColumnWithUpdateMode(columnID int, dataType DataType, updateMode ColumnUpdateMode) error {
	_ = "STUB: not implemented"
	return nil
}

// AddRow increases the number of rows in the batch by 1. A new row with all nil values is appended
// to the row array.
func (u *UpsertBatchBuilder) AddRow() { _ = "STUB: not implemented"; return }

// RemoveRow decreases the number of rows in the batch by 1. The last row will be removed. It's a
// no-op if the number of rows is 0.
func (u *UpsertBatchBuilder) RemoveRow() { _ = "STUB: not implemented"; return }

// ResetRows reset the row count to 0.
func (u *UpsertBatchBuilder) ResetRows() { _ = "STUB: not implemented"; return }

// SetValue set a value to a given (row, col).
func (u *UpsertBatchBuilder) SetValue(row int, col int, value interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (u *UpsertBatchBuilder) MarkFactTable() { _ = "STUB: not implemented"; return }

// ToByteArray produces a serialized UpsertBatch in byte array.
func (u UpsertBatchBuilder) ToByteArray() ([]byte, error) {
	_ = "STUB: not implemented"
	// Create buffer.
	return nil, nil
}

// initialized size to 4 bytes (version number).

// 24 bytes consist of fixed headers:
// [int32] num_of_rows (4 bytes)
// [uint16] num_of_columns (2 bytes)
// <reserve 14 bytes>
// [uint32] arrival_time (4 bytes)

// Write upsert batch version.

// Write fixed headers.

// skip to data offset

// Write per column data their headers.

func AdditionUpdate(oldValue, newValue unsafe.Pointer, dataType DataType) {
	_ = "STUB: not implemented"
	return
}

// MinMaxUpdate update the old value if compareRes == expectedRes
func MinMaxUpdate(oldValue, newValue unsafe.Pointer, dataType DataType, cmpFunc CompareFunc, expectedRes int) {
	_ = "STUB: not implemented"
	return
}
