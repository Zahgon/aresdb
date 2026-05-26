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

package memstore

import (
	"io"
	"os"
	"unsafe"

	"github.com/uber/aresdb/memstore/common"
)

// cLiveVectorParty is the implementation of LiveVectorParty with c allocated memory
// this vector party stores columns with fixed length data type
type cLiveVectorParty struct {
	cVectorParty
}

// SetBool implements SetBool in LiveVectorParty interface
func (vp *cLiveVectorParty) SetBool(offset int, val bool, valid bool) {
	_ = "STUB: not implemented"
	return
}

// SetBool implements SetValue in LiveVectorParty interface
func (vp *cLiveVectorParty) SetValue(offset int, val unsafe.Pointer, valid bool) {
	_ = "STUB: not implemented"
	return
}

// SetGoValue implements SetGoValue in LiveVectorParty interface
func (vp *cLiveVectorParty) SetGoValue(offset int, val common.GoDataValue, valid bool) {
	_ = "STUB: not implemented"
	return
}

// GetValue implements GetValue in LiveVectorParty interface
func (vp *cLiveVectorParty) GetValue(offset int) (unsafe.Pointer, bool) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), false
}

// goLiveVectorParty is the implementation of LiveVectorParty with go allocated memory
// this vector party stores columns with variable length data type
type goLiveVectorParty struct {
	baseVectorParty

	values            []common.GoDataValue
	hostMemoryManager common.HostMemoryManager

	totalBytes int64
}

// GetMinMaxValue implements GetMinMaxValue in LiveVectorParty interface
func (vp *cLiveVectorParty) GetMinMaxValue() (min uint32, max uint32) {
	_ = "STUB: not implemented"
	return 0, 0
}

// Allocate implements Allocate in VectorParty interface
func (vp *cLiveVectorParty) Allocate(hasCount bool) { _ = "STUB: not implemented"; return }

// Allocate implements Allocate in VectorParty interface
func (vp *goLiveVectorParty) Allocate(hasCount bool) { _ = "STUB: not implemented"; return }

// SetDataValue implements SetDataValue in VectorParty interface
// liveVectorParty ignores countsUpdateMode or counts
func (vp *goLiveVectorParty) SetDataValue(offset int, value common.DataValue,
	countsUpdateMode common.ValueCountsUpdateMode, counts ...uint32) {
	_ = "STUB: not implemented"
	return
}

// SetBool implements SetBool in LiveVectorParty interface
func (vp *goLiveVectorParty) SetBool(offset int, val bool, valid bool) {
	_ = "STUB: not implemented"
	return
}

// SetValue implements SetValue in LiveVectorParty interface
func (vp *goLiveVectorParty) SetValue(offset int, val unsafe.Pointer, valid bool) {
	_ = "STUB: not implemented"
	return
}

// GetValue implements GetValue in LiveVectorParty interface
func (vp *goLiveVectorParty) GetValue(offset int) (unsafe.Pointer, bool) {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer), false
}

// SetGoValue implements SetGoValue in LiveVectorParty interface
func (vp *goLiveVectorParty) SetGoValue(offset int, val common.GoDataValue, valid bool) {
	_ = "STUB: not implemented"
	return
}

// GetDataValue implements GetDataValue in VectorParty interface
func (vp *goLiveVectorParty) GetDataValue(offset int) common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

// GetDataValueByRow implements GetDataValueByRow in VectorParty interface
func (vp *goLiveVectorParty) GetDataValueByRow(row int) common.DataValue {
	_ = "STUB: not implemented"
	return *new(common.DataValue)
}

// GetValidity implements GetValidity in VectorParty interface
func (vp *goLiveVectorParty) GetValidity(offset int) bool { _ = "STUB: not implemented"; return false }

// GetMinMaxValue is **not supported** by goLiveVectorParty
func (vp *goLiveVectorParty) GetMinMaxValue() (min uint32, max uint32) {
	_ = "STUB: not implemented"

	// GetBytes implements GetBytes in VectorParty interface
	return 0, 0
}

func (vp *goLiveVectorParty) GetBytes() int64 { _ = "STUB: not implemented"; return 0 }

// Slice implements Slice in VectorParty interface
func (vp *goLiveVectorParty) Slice(startRow, numRows int) common.SlicedVector {
	_ = "STUB: not implemented"
	return *

	// size is the number of entries in the vector,
	new(common.SlicedVector)
}

// Write implements Write in VectorParty interface
func (vp *goLiveVectorParty) Write(writer io.Writer) error { _ = "STUB: not implemented"; return nil }

// write total bytes for reporting during loading

// write length

// count non nil values

// write number of valid values

// write values

// only write index if not all valid

// Read implements Read in VectorParty interface
func (vp *goLiveVectorParty) Read(reader io.Reader, serializer common.VectorPartySerializer) error {
	_ = "STUB: not implemented"
	return nil
}

// read total bytes for reporting during loading

// SafeDestruct implements SafeDestruct in VectorParty interface
func (vp *goLiveVectorParty) SafeDestruct() { _ = "STUB: not implemented"; return }

// Equals implements Equals in VectorParty interface
func (vp *goLiveVectorParty) Equals(other common.VectorParty) bool {
	_ = "STUB: not implemented"
	return false
}

func (vp *goLiveVectorParty) Dump(file *os.File) { _ = "STUB: not implemented"; return }

// NewLiveVectorParty creates LiveVectorParty
func NewLiveVectorParty(length int, dataType common.DataType, defaultValue common.DataValue, hostMemoryManager common.HostMemoryManager) common.LiveVectorParty {
	_ = "STUB: not implemented"
	return *new(common.LiveVectorParty)
}

// newCLiveVectorParty creates a LiveVectorParty with c allocated memory
func newCLiveVectorParty(length int, dataType common.DataType, defaultValue common.DataValue) *cLiveVectorParty {
	_ = "STUB: not implemented"
	return nil
}

// newGoLiveVetorParty creates a LiveVectorParty with go allocated memory
func newGoLiveVetorParty(length int, dataType common.DataType, hostMemoryManager common.HostMemoryManager) *goLiveVectorParty {
	_ = "STUB: not implemented"
	return nil
}
