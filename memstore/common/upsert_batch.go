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
)

// columnReader contains meta data for accessing the data of a column in an UpsertBatch.
type columnReader struct {
	// The logic id of the column.
	columnID int
	// The column mode.
	columnMode ColumnMode
	// The column update mode
	columnUpdateMode ColumnUpdateMode
	// DataType of the column.
	dataType DataType
	// The value vector. can be empty depending on column mode.
	valueVector []byte
	// The null vector. can be empty depending on column mode.
	nullVector []byte
	// The offset vector. Only used for variable length values. Not used yet.
	offsetVector []byte
	// Compare function if any.
	cmpFunc CompareFunc
}

// ReadGoValue returns the GoDataValue from upsert batch at given row
func (c *columnReader) ReadGoValue(row int) GoDataValue {
	_ = "STUB: not implemented"
	return *new(GoDataValue)
}

// ReadValue returns the row data (fixed sized) for a column, including the pointer to the data,
// and the validity of the value.
func (c *columnReader) ReadValue(row int) (unsafe.Pointer, bool) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), false
}

// ReadArrayValue returns the ArrayValue from upsert batch at given row
func (c *columnReader) readArrayValue(row int) (unsafe.Pointer, bool) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), false
}

// ReadValue returns the row data (boolean type) for a column, and its validity.
func (c *columnReader) ReadBool(row int) (bool, bool) {
	_ = "STUB: not implemented"
	return false, false
}

func (c *columnReader) readOffset(row int) uint32 { _ = "STUB: not implemented"; return 0 }

// readValidity return the validity value of a row in a column.
func (c *columnReader) readValidity(row int) bool { _ = "STUB: not implemented"; return false }

func readBool(buffer []byte, index int) bool { _ = "STUB: not implemented"; return false }

func writeBool(buffer []byte, index int, value bool) { _ = "STUB: not implemented"; return }

// UpsertBatch stores and indexes a serialized upsert batch of data on a particular table.
// It is used for both client-server data transfer and redo logging.
// In redo logs each batch is prepended by a 4-byte buffer size.
// The serialized buffer of the batch is in the following format:
//
//	[uint32] magic_number
//	[uint32] buffer_size
//
//	<begin of buffer>
//	[int32]  version_number
//	[int32]  num_of_rows
//	[uint16] num_of_columns
//	<reserve 14 bytes>
//	[uint32] arrival_time
//	[uint32] column_offset_0 ... [uint32] column_offset_x+1
//	[uint32] column_reserved_field1_0 ... [uint32] column_reserved_field1_x
//	[uint32] column_reserved_field2_0 ... [uint32] column_reserved_field2_x
//	[uint32] column_data_type_0 ... [uint32] column_data_type_x
//	[uint16] column_id_0 ... [uint16] column_id_x
//	[uint8] column_mode_0 ... [uint8] column_mode_x
//
//	(optional) [uint8] null_vector_0
//	(optional) [padding to 4 byte alignment uint32] offset_vector_0
//	[padding for 8 byte alignment] value_vector_0
//	...
//
//	[padding for 8 byte alignment]
//	<end of buffer>
//
// Each component in the serialized buffer is byte aligned (not pointer aligned or bit aligned).
// All serialized numbers are written in little-endian.
// The struct is used for both client serialization and server deserialization.
// See https://github.com/uber/aresdb/wiki/redo_logs for more details.
//
// Note: only fixed size values are supported currently.
type UpsertBatch struct {
	// Number of rows in the batch, must be between 0 and 65535.
	NumRows int

	// Number of columns.
	NumColumns int

	// Arrival Time of Upsert Batch
	ArrivalTime uint32

	// Serialized buffer of the batch, starts from NumRows, does not contain the 4-byte
	// buffer size.
	buffer []byte

	// When records are extracted for backfill, buffer is no longer used and we
	// use alternativeBytes to track the memory usage.
	alternativeBytes int

	// Columns to upsert on.
	columns []*columnReader

	// Column id maps the logic column id to local column index.
	columnsByID map[int]int
}

// GetBuffer returns the underline buffer used to construct the upsert batch.
func (u *UpsertBatch) GetBuffer() []byte {
	_ = "STUB: not implemented"

	// GetAlternativeBytes returns alternativeBytes
	return nil
}

func (u *UpsertBatch) GetAlternativeBytes() int { _ = "STUB: not implemented"; return 0 }

// convenient function to get columns len
func (u *UpsertBatch) GetColumnLen() int { _ = "STUB: not implemented"; return 0 }

// convenient function to get ColumnMode, assume no out of index
func (u *UpsertBatch) GetColumMode(col int) ColumnMode {
	_ = "STUB: not implemented"
	return *new(ColumnMode)
}

// convenient function to get ColumnUpdateMode, assume no out of index
func (u *UpsertBatch) GetColumnUpdateMode(col int) ColumnUpdateMode {
	_ = "STUB: not implemented"
	return *new(ColumnUpdateMode)
}

// convenient function to get GoDataValue
func (u *UpsertBatch) ReadGoValue(row, col int) GoDataValue {
	_ = "STUB: not implemented"
	return *new(GoDataValue)
}

// GetColumnID returns the logical id of a column.
func (u *UpsertBatch) GetColumnID(col int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// GetColumnType returns the data type of a column.
func (u *UpsertBatch) GetColumnType(col int) (DataType, error) {
	_ = "STUB: not implemented"
	return *new(DataType), nil
}

// GetColumnIndex returns the local index of a column given a logical index id.
func (u *UpsertBatch) GetColumnIndex(columnID int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// GetValue returns the data (fixed sized) stored at (row, col), including the pointer to the data,
// and the validity of the value.
func (u *UpsertBatch) GetValue(row int, col int) (unsafe.Pointer, bool, error) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), false, nil
}

// GetBool returns the data (boolean type) stored at (row, col), and the validity of the value.
func (u *UpsertBatch) GetBool(row int, col int) (bool, bool, error) {
	_ = "STUB: not implemented"
	return false, false, nil
}

// GetDataValue returns the DataValue for the given row and col index.
// It first check validity of the value, then it check whether it's a
// boolean column to decide whether to load bool value or other value
// type.
// user of GetDataValue should check row, col using NumRows and NumColumns
func (u *UpsertBatch) GetDataValue(row, col int) DataValue {
	_ = "STUB: not implemented"
	return *new(DataValue)
}

// GetDataValueWithDefault get the data value at with row and col and return defaultVal when row, col is out of bound
func (u *UpsertBatch) GetDataValueWithDefault(row, col int, defaultVal DataValue) DataValue {
	_ = "STUB: not implemented"
	return *new(DataValue)
}

// GetEventColumnIndex returns the column index of event time
func (u *UpsertBatch) GetEventColumnIndex() int {
	_ = "STUB: not implemented"
	// Validate columns in upsert batch are valid.
	return 0
}

// GetPrimaryKeyCols converts primary key columnIDs to cols in this upsert batch.
func (u *UpsertBatch) GetPrimaryKeyCols(primaryKeyColumnIDs []int) ([]int, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtractBackfillBatch extracts given rows and stores in a new UpsertBatch
// The returned new UpsertBatch is not fully serialized and can only be used for
// structured reads.
func (u *UpsertBatch) ExtractBackfillBatch(backfillRows []int) *UpsertBatch {
	_ = "STUB: not implemented"
	return nil
}

// ignore those columns with conditional updates from backfill
// clean up the column data

// GetColumnNames reads columnNames in UpsertBatch, user should not lock schema
func (u *UpsertBatch) GetColumnNames(schema *TableSchema) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadData reads data from upsert batch and convert values to meaningful representations given data type.
func (u *UpsertBatch) ReadData(start int, length int) ([][]interface{}, error) {
	_ = "STUB: not implemented"
	// Only read the column names.
	return nil, nil
}

func readUpsertBatch(buffer []byte) (*UpsertBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// numRows.

// numColumns.

// 2 byte num columns

// Header too small, error out.

// Null vector points to the beginning of the column data section.

// Round up to 8 byte padding.

// NewUpsertBatch deserializes an upsert batch on the server.
// buffer does not contain the 4-byte buffer size.
func NewUpsertBatch(buffer []byte) (*UpsertBatch, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// read version
