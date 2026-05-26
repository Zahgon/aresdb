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
	"math"
	"unsafe"

	"github.com/uber/aresdb/utils"
)

// NullDataValue is a global data value that stands a null value where the newly added
// columns haven't received any data.
var NullDataValue = DataValue{}

// SizeOfGeoPoint is the size of GeoPointGo in memory
const SizeOfGeoPoint = unsafe.Sizeof(GeoPointGo{})

// ZeroLengthArrayFlag is the value to represent 0 length array value
// used in offset part of OffsetLength vector
// when length part is non-zero value, the offset is real offset in memory
// when length is 0 and offset is 0, the array value is invalid (by default)
// when length is 0 and offset is ZeroLengthArrayFlag, this is a 0 length array value
const ZeroLengthArrayFlag = math.MaxUint32

// CompareFunc represents compare function
type CompareFunc func(a, b unsafe.Pointer) int

// CompareBool compares boolean value
func CompareBool(a, b bool) int { _ = "STUB: not implemented"; return 0 }

// CompareInt8 compares int8 value
func CompareInt8(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareUint8 compares uint8 value
func CompareUint8(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareInt16 compares int16 value
func CompareInt16(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareUint16 compares uint16 value
func CompareUint16(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareInt32 compares int32 value
func CompareInt32(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareUint32 compares uint32 value
func CompareUint32(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareInt64 compares int64 value
func CompareInt64(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareFloat32 compares float32 value
func CompareFloat32(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareUUID compare UUID values
func CompareUUID(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareGeoPoint compare GeoPoint Values
func CompareGeoPoint(a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// CompareArray compare array values, the main purpose of this comparsion is for equal comparison
// larger/less comparision may not be accurate
func CompareArray(dataType DataType, a, b unsafe.Pointer) int { _ = "STUB: not implemented"; return 0 }

// skip 4 bytes for length

// ArrayLengthCompare compare
func ArrayLengthCompare(v1, v2 *DataValue) int { _ = "STUB: not implemented"; return 0 }

// GetCompareFunc get the compare function for specific data type
func GetCompareFunc(dataType DataType) CompareFunc {
	_ = "STUB: not implemented"
	return *new(CompareFunc)
}

// GoDataValue represents a value backed in golang memory
type GoDataValue interface {
	// GetBytes returns number of bytes copied in golang memory for this value
	GetBytes() int
	// GetSerBytes return the number of bytes required for serialize this value
	GetSerBytes() int
	Write(writer *utils.StreamDataWriter) error
	Read(reader *utils.StreamDataReader) error
}

// DataValueIterator is a iterator of data value
type DataValueIterator interface {
	// read the current data value
	read() DataValue
	// advance iterator
	next()
	// whether iterator is done
	done() bool
}

// DataValue is the wrapper to encapsulate validity, bool value and other value type
// into a single struct to make it easier for value comparison.
type DataValue struct {
	// Used for golang vector party
	GoVal    GoDataValue
	OtherVal unsafe.Pointer
	DataType DataType
	CmpFunc  CompareFunc
	Valid    bool

	IsBool  bool
	BoolVal bool
}

// GeoPointGo represents GeoPoint Golang Type
type GeoPointGo [2]float32

// GeoShapeGo represents GeoShape Golang Type
type GeoShapeGo struct {
	Polygons [][]GeoPointGo
}

// Array value representation in Go for UpsertBatch
type ArrayValue struct {
	// item data type
	DataType DataType
	// item list
	Items []interface{}
}

// Compare compares two value wrapper.
func (v1 DataValue) Compare(v2 DataValue) int { _ = "STUB: not implemented"; return 0 }

// ConvertToHumanReadable convert DataValue to meaningful golang data types
func (v1 DataValue) ConvertToHumanReadable(dataType DataType) interface{} {
	_ = "STUB: not implemented"
	return nil
}

// in string format, lng goes first and lat second

// in string format, lng goes first and lat second
// https://en.wikipedia.org/wiki/Well-known_text_representation_of_geometry

// ValueFromString converts raw string value to actual value given input data type.
func ValueFromString(str string, dataType DataType) (val DataValue, err error) {
	_ = "STUB: not implemented"
	return *new(DataValue), nil
}

// We need to convert it from i64 to i8 since strconv.ParseXXX
// always returns the largest bit size value.

// GetBytes implements GoDataValue interface
func (gs *GeoShapeGo) GetBytes() int { _ = "STUB: not implemented"; return 0 }

// GetSerBytes implements GoDataValue interface
func (gs *GeoShapeGo) GetSerBytes() int {
	_ = "STUB: not implemented"

	// 1. numPolygons (uint32)
	return 0
}

// numPoints (uint32)

// 8 bytes per point [2]float32

// Read implements Read interface for GoDataValue
func (gs *GeoShapeGo) Read(dataReader *utils.StreamDataReader) error {
	_ = "STUB: not implemented"
	return nil
}

// Write implements Read interface for GoDataValue
func (gs *GeoShapeGo) Write(dataWriter *utils.StreamDataWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// GetLength return item numbers for the array value
func (av *ArrayValue) GetLength() int { _ = "STUB: not implemented"; return 0 }

// AddItem add new item into array
func (av *ArrayValue) AddItem(item interface{}) { _ = "STUB: not implemented"; return }

// GetSerBytes return the bytes will be used in upsertbatch serialized format
func (av *ArrayValue) GetSerBytes() int { _ = "STUB: not implemented"; return 0 }

// we always align to 8 bytes in upsertbatch

// NewArrayValue create a new ArrayValue instance
func NewArrayValue(dataType DataType) *ArrayValue { _ = "STUB: not implemented"; return nil }

// Write serialize data into writer
// Serialized Array data format:
// number of items: 4 bytes
// item values: per item bytes * number of items, align to byte
// item validity:  1 bit * number of items
// final align to 8 bytes
func (av *ArrayValue) Write(writer *utils.BufferWriter) error {
	_ = "STUB: not implemented"
	return nil
}

// add value for each item

// add validity bit for each item

// ArrayValueReader is an aux class to reader item data from bytes buffer
type ArrayValueReader struct {
	itemType DataType
	value    unsafe.Pointer
	length   int
}

// NewArrayValueReader is to create ArrayValueReader to read from upsertbatch, which includes the item number
func NewArrayValueReader(dataType DataType, value unsafe.Pointer) *ArrayValueReader {
	_ = "STUB: not implemented"
	return nil
}

// GetLength return item numbers inside the array
func (reader *ArrayValueReader) GetLength() int { _ = "STUB: not implemented"; return 0 }

// GetBytes returns the bytes counts this value occopies
func (reader *ArrayValueReader) GetBytes() int { _ = "STUB: not implemented"; return 0 }

// GetBool returns bool value for Bool item type at index
func (reader *ArrayValueReader) GetBool(index int) bool { _ = "STUB: not implemented"; return false }

// Get returns the buffer pointer for the index-th item
func (reader *ArrayValueReader) Get(index int) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// IsItemValid check if the item in index-th place is valid or not
func (reader *ArrayValueReader) IsItemValid(index int) bool {
	_ = "STUB: not implemented"
	return false
}

// CalculateListElementBytes returns the total size in bytes needs to be allocated for a list type column for a single
// row along with the validity vector start.
func CalculateListElementBytes(dataType DataType, length int) int {
	_ = "STUB: not implemented"
	return 0
}

// there is a item number at beginning
// element_number_bits => 8 * 4 (4 bytes)
// DataTypeBits(dataType) * length => element_bits, round to byte
// 1 * length => null bits, round to byte
// (element_number_bits + element_bits + null_bits + 63) / 64 => round by 64 bits (8 bytes)

func CalculateListNilOffset(dataType DataType, length int) int { _ = "STUB: not implemented"; return 0 }

// GetDataValue returns the DataValue for the given column value.
func GetDataValue(col interface{}, columnIDInSchema int, columnType string) (DataValue, error) {
	_ = "STUB: not implemented"
	return *new(DataValue), nil
}

type baseDataValueIterator struct {
	curIdx int
	size   int
}

func (iter *baseDataValueIterator) read() DataValue {
	_ = "STUB: not implemented"
	return *new(DataValue)
}

func (iter *baseDataValueIterator) next() { _ = "STUB: not implemented"; return }

func (iter *baseDataValueIterator) done() bool { _ = "STUB: not implemented"; return false }

type primaryKeyDataValueIterator struct {
	baseDataValueIterator

	primaryKeyCols []int
	batchReader    BatchReader
	row            int
}

// NewPrimaryKeyDataValueIterator creates DataValueIterator from upsert batch, primary key cols and row number
func NewPrimaryKeyDataValueIterator(batchReader BatchReader, row int, primaryKeyCols []int) DataValueIterator {
	_ = "STUB: not implemented"
	return *new(DataValueIterator)
}

func (iter *primaryKeyDataValueIterator) read() DataValue {
	_ = "STUB: not implemented"
	return *new(DataValue)
}

type sliceDataValueIterator struct {
	baseDataValueIterator
	values []DataValue
}

// NewSliceDataValueIterator creates DataValueIterator from slice of DataValue
func NewSliceDataValueIterator(values []DataValue) DataValueIterator {
	_ = "STUB: not implemented"
	return *new(DataValueIterator)
}

func (iter *sliceDataValueIterator) read() DataValue {
	_ = "STUB: not implemented"
	return *new(DataValue)
}
