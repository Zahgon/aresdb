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
	metaCom "github.com/uber/aresdb/metastore/common"
)

// DataType is the type of value supported in AresDB.
type DataType uint32

// The list of supported DataTypes.
// DataType & 0x0000FFFF: The width of the data type in bits, or width of the item data type for array.
// DataType & 0x00FF0000 >> 16: The base type of the data or array.
// DataType & 0x01000000 >> 24: Indicatation of arrary type, and item type is on DataType & 0x00FF0000 >> 16.
// DataType & 0xFE000000 >> 24: Reserved
// See https://github.com/uber/aresdb/wiki/redologs for more details.
const (
	Unknown   DataType = 0x00000000
	Bool      DataType = 0x00000001
	Int8      DataType = 0x00010008
	Uint8     DataType = 0x00020008
	Int16     DataType = 0x00030010
	Uint16    DataType = 0x00040010
	Int32     DataType = 0x00050020
	Uint32    DataType = 0x00060020
	Float32   DataType = 0x00070020
	SmallEnum DataType = 0x00080008
	BigEnum   DataType = 0x00090010
	UUID      DataType = 0x000a0080
	GeoPoint  DataType = 0x000b0040
	GeoShape  DataType = 0x000c0000
	Int64     DataType = 0x000d0040

	// array types
	ArrayBool      DataType = 0x01000001
	ArrayInt8      DataType = 0x01010008
	ArrayUint8     DataType = 0x01020008
	ArrayInt16     DataType = 0x01030010
	ArrayUint16    DataType = 0x01040010
	ArrayInt32     DataType = 0x01050020
	ArrayUint32    DataType = 0x01060020
	ArrayFloat32   DataType = 0x01070020
	ArraySmallEnum DataType = 0x01080008
	ArrayBigEnum   DataType = 0x01090010
	ArrayUUID      DataType = 0x010a0080
	ArrayGeoPoint  DataType = 0x010b0040
	ArrayInt64     DataType = 0x010d0040
)

// DataTypeName returns the literal name of the data type.
var DataTypeName = map[DataType]string{
	Unknown:   "Unknown",
	Bool:      metaCom.Bool,
	Int8:      metaCom.Int8,
	Uint8:     metaCom.Uint8,
	Int16:     metaCom.Int16,
	Uint16:    metaCom.Uint16,
	Int32:     metaCom.Int32,
	Uint32:    metaCom.Uint32,
	Float32:   metaCom.Float32,
	SmallEnum: metaCom.SmallEnum,
	BigEnum:   metaCom.BigEnum,
	UUID:      metaCom.UUID,
	GeoPoint:  metaCom.GeoPoint,
	GeoShape:  metaCom.GeoShape,
	Int64:     metaCom.Int64,

	// array types
	ArrayBool:      metaCom.ArrayBool,
	ArrayInt8:      metaCom.ArrayInt8,
	ArrayUint8:     metaCom.ArrayUint8,
	ArrayInt16:     metaCom.ArrayInt16,
	ArrayUint16:    metaCom.ArrayUint16,
	ArrayInt32:     metaCom.ArrayInt32,
	ArrayUint32:    metaCom.ArrayUint32,
	ArrayFloat32:   metaCom.ArrayFloat32,
	ArraySmallEnum: metaCom.ArraySmallEnum,
	ArrayBigEnum:   metaCom.ArrayBigEnum,
	ArrayUUID:      metaCom.ArrayUUID,
	ArrayGeoPoint:  metaCom.ArrayGeoPoint,
	ArrayInt64:     metaCom.ArrayInt64,
}

// StringToDataType maps string representation to DataType
var StringToDataType = map[string]DataType{
	metaCom.Bool:      Bool,
	metaCom.Int8:      Int8,
	metaCom.Uint8:     Uint8,
	metaCom.Int16:     Int16,
	metaCom.Uint16:    Uint16,
	metaCom.Int32:     Int32,
	metaCom.Uint32:    Uint32,
	metaCom.Float32:   Float32,
	metaCom.SmallEnum: SmallEnum,
	metaCom.BigEnum:   BigEnum,
	metaCom.UUID:      UUID,
	metaCom.GeoPoint:  GeoPoint,
	metaCom.GeoShape:  GeoShape,
	metaCom.Int64:     Int64,

	// array types
	metaCom.ArrayBool:      ArrayBool,
	metaCom.ArrayInt8:      ArrayInt8,
	metaCom.ArrayUint8:     ArrayUint8,
	metaCom.ArrayInt16:     ArrayInt16,
	metaCom.ArrayUint16:    ArrayUint16,
	metaCom.ArrayInt32:     ArrayInt32,
	metaCom.ArrayUint32:    ArrayUint32,
	metaCom.ArrayFloat32:   ArrayFloat32,
	metaCom.ArraySmallEnum: ArraySmallEnum,
	metaCom.ArrayBigEnum:   ArrayBigEnum,
	metaCom.ArrayUUID:      ArrayUUID,
	metaCom.ArrayGeoPoint:  ArrayGeoPoint,
	metaCom.ArrayInt64:     ArrayInt64,
}

// NewDataType converts an uint32 value into a DataType. It returns error if the the data type is
// invalid.
func NewDataType(value uint32) (DataType, error) {
	_ = "STUB: not implemented"
	return *new(DataType), nil
}

// IsNumeric determines whether a data type is numeric
func IsNumeric(dataType DataType) bool { _ = "STUB: not implemented"; return false }

// IsArrayType determins where a data type is Array
func IsArrayType(dataType DataType) bool { _ = "STUB: not implemented"; return false }

// GetElementDataType retrieve item data type for Array DataType
func GetElementDataType(dataType DataType) DataType {
	_ = "STUB: not implemented"
	return *new(DataType)
}

// DataTypeBits returns the number of bits of a data type.
func DataTypeBits(dataType DataType) int { _ = "STUB: not implemented"; return 0 }

// DataTypeForColumn returns the in memory data type for a column
func DataTypeForColumn(column metaCom.Column) DataType {
	_ = "STUB: not implemented"
	return *new(DataType)
}

// DataTypeFromString convert string representation of data type into DataType
func DataTypeFromString(str string) DataType { _ = "STUB: not implemented"; return *new(DataType) }

// DataTypeBytes returns how many bytes a value of the data type occupies.
func DataTypeBytes(dataType DataType) int { _ = "STUB: not implemented"; return 0 }

// ConvertValueForType converts data value based on data type
func ConvertValueForType(dataType DataType, value interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertToBool convert input into bool at best effort
func ConvertToBool(value interface{}) (bool, bool) { _ = "STUB: not implemented"; return false, false }

// try converting "true" "false"

// try convert as number

// ConvertToInt8 convert input into int8 at best effort
func ConvertToInt8(value interface{}) (int8, bool) { _ = "STUB: not implemented"; return 0, false }

// ConvertToUint8 convert input into uint8 at best effort
func ConvertToUint8(value interface{}) (uint8, bool) { _ = "STUB: not implemented"; return 0, false }

// ConvertToInt16 convert input into int16 at best effort
func ConvertToInt16(value interface{}) (int16, bool) { _ = "STUB: not implemented"; return 0, false }

// ConvertToUint16 convert input into uint16 at best effort
func ConvertToUint16(value interface{}) (uint16, bool) { _ = "STUB: not implemented"; return 0, false }

// ConvertToInt32 convert input into int32 at best effort
func ConvertToInt32(value interface{}) (int32, bool) { _ = "STUB: not implemented"; return 0, false }

// ConvertToUint32 convert input into uint32 at best effort
func ConvertToUint32(value interface{}) (uint32, bool) { _ = "STUB: not implemented"; return 0, false }

// ConvertToFloat32 convert input into float32 at best effort
func ConvertToFloat32(value interface{}) (float32, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

// ConvertToUint64 convert input into uint64 at best effort
func ConvertToUint64(value interface{}) (uint64, bool) { _ = "STUB: not implemented"; return 0, false }

// ConvertToInt64 convert input into int64 at best effort
func ConvertToInt64(value interface{}) (int64, bool) { _ = "STUB: not implemented"; return 0, false }

// ConvertToFloat64 convert input into float64 at best effort
func ConvertToFloat64(value interface{}) (float64, bool) {
	_ = "STUB: not implemented"
	return 0, false
}

func UUIDFromString(str string) ([2]uint64, error) { _ = "STUB: not implemented"; return nil, nil }

// ConvertToUUID convert input into uuid type ([2]uint64) at best effort
func ConvertToUUID(value interface{}) ([2]uint64, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GeoPointFromString convert string to geopoint
// we support wkt format, eg. Point(lng,lat)
// Inside AresDB system we store lat,lng format
func GeoPointFromString(str string) (point [2]float32, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ConvertToGeoPoint convert input into uuid type ([2]float32) at best effort
func ConvertToGeoPoint(value interface{}) ([2]float32, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// GeoShapeFromString convert string to geoshape
// Supported format POLYGON ((lng lat, lng lat, lng lat, ...), (...))
func GeoShapeFromString(str string) (GeoShapeGo, error) {
	_ = "STUB: not implemented"
	return *new(GeoShapeGo), nil
}

// ConvertToGeoShape converts the arbitrary value to GeoShapeGo
func ConvertToGeoShape(value interface{}) (*GeoShapeGo, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

// IsGoType determines whether a data type is golang type
func IsGoType(dataType DataType) bool {
	_ = "STUB: not implemented"
	// for now we only have GeoShape
	return false
}

// IsEnumType determines whether a data type is enum type
func IsEnumType(dataType DataType) bool { _ = "STUB: not implemented"; return false }

// GetGoDataValue return GoDataValue
func GetGoDataValue(dataType DataType) GoDataValue {
	_ = "STUB: not implemented"
	return *new(GoDataValue)
}

// ConvertToArrayValue convert input to ArrayValue at best effort
func ConvertToArrayValue(dataType DataType, value interface{}) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArrayValueFromArray convert any array to array of sepecified item data type
func ArrayValueFromArray(value []interface{}, dataType DataType) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ArrayValueFromString convert string to array of sepecified item data type
// we can support json formatted string like:
// string array: "[\"11\",\"12\",\"13\"]"
// integer array: "[11,12,13]"
// string array of uuid: "[\"1e88a975-3d26-4277-ace9-bea91b072977\",\"1e88a975-3d26-4277-ace9-bea91b072978\",\"1e88a975-3d26-4277-ace9-bea91b072979\"]"
// string array of geo-point: "[\"Point(180.0, 90.0)\",\"Point(179.0, 89.0)\",\"Point(178.0, 88.0)\"]"
// we can also support comma delimited string value as long as the item not contain comma
// "11,12,13"
func ArrayValueFromString(value string, dataType DataType) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
