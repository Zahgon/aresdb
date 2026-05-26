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

// ColumnHeaderSize returns the total size of the column headers.
func ColumnHeaderSize(numCols int) int { _ = "STUB: not implemented"; return 0 }

// offset (4 bytes)
// enum dict length (4 bytes)
// reserved (4 bytes)
// data_type (4 bytes)
// column_id (2 bytes)
// column mode (1 byte)

// UpsertBatchHeader is a helper class used by upsert batch reader and writer to access the column
// header info.
type UpsertBatchHeader struct {
	offsetVector   []byte
	enumDictLength []byte
	typeVector     []byte
	idVector       []byte
	modeVector     []byte
}

// NewUpsertBatchHeader create upsert batch header from buffer
func NewUpsertBatchHeader(buffer []byte, numCols int) UpsertBatchHeader {
	_ = "STUB: not implemented"

	// Offset vector is of size numCols + 1.
	return *new(UpsertBatchHeader)
}

// reserved extra space

// WriteColumnOffset writes the offset of a column. It can take col index from 0 to numCols + 1.
func (u *UpsertBatchHeader) WriteColumnOffset(value int, col int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteEnumDictLength writes the offset of a column. It can take col index from 0 to numCols - 1.
func (u *UpsertBatchHeader) WriteEnumDictLength(value int, col int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteColumnType writes the type of a column.
func (u *UpsertBatchHeader) WriteColumnType(value DataType, col int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteColumnID writes the id of a column.
func (u *UpsertBatchHeader) WriteColumnID(value int, col int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteColumnFlag writes the mode of a column.
func (u *UpsertBatchHeader) WriteColumnFlag(columnMode ColumnMode, columnUpdateMode ColumnUpdateMode, col int) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadColumnOffset takes col index from 0 to numCols + 1 and returns the value stored.
func (u UpsertBatchHeader) ReadColumnOffset(col int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadEnumDictLength takes col index from 0 to numCols - 1 and returns the value stored.
func (u UpsertBatchHeader) ReadEnumDictLength(col int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadColumnType returns the type for a column.
func (u UpsertBatchHeader) ReadColumnType(col int) (DataType, error) {
	_ = "STUB: not implemented"
	return *new(DataType), nil
}

// ReadColumnID returns the logical ID for a column.
func (u UpsertBatchHeader) ReadColumnID(col int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadColumnFlag returns the mode for a column.
func (u UpsertBatchHeader) ReadColumnFlag(col int) (ColumnMode, ColumnUpdateMode, error) {
	_ = "STUB: not implemented"
	return *new(ColumnMode), *new(ColumnUpdateMode), nil
}
