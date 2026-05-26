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

package query

// #cgo LDFLAGS: -L${SRCDIR}/../lib -lalgorithm
// #include "time_series_aggregate.h"
import "C"
import (
	"unsafe"

	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/query/common"
	"github.com/uber/aresdb/query/expr"
)

// DataTypeToCDataType mapps from memstore data type to c data types
var DataTypeToCDataType = map[memCom.DataType]C.enum_DataType{
	memCom.Bool:      C.Bool,
	memCom.Int8:      C.Int8,
	memCom.Uint8:     C.Uint8,
	memCom.Int16:     C.Int16,
	memCom.Uint16:    C.Uint16,
	memCom.Int32:     C.Int32,
	memCom.Int64:     C.Int64,
	memCom.Uint32:    C.Uint32,
	memCom.Float32:   C.Float32,
	memCom.SmallEnum: C.Uint8,
	memCom.BigEnum:   C.Uint16,
	memCom.GeoPoint:  C.GeoPoint,
	memCom.UUID:      C.UUID,
}

// UnaryExprTypeToCFunctorType maps from unary operator to C UnaryFunctorType
var UnaryExprTypeToCFunctorType = map[expr.Token]C.enum_UnaryFunctorType{
	expr.NOT:                 C.Not,
	expr.UNARY_MINUS:         C.Negate,
	expr.IS_NULL:             C.IsNull,
	expr.IS_NOT_NULL:         C.IsNotNull,
	expr.BITWISE_NOT:         C.BitwiseNot,
	expr.GET_WEEK_START:      C.GetWeekStart,
	expr.GET_MONTH_START:     C.GetMonthStart,
	expr.GET_QUARTER_START:   C.GetQuarterStart,
	expr.GET_YEAR_START:      C.GetYearStart,
	expr.GET_DAY_OF_MONTH:    C.GetDayOfMonth,
	expr.GET_DAY_OF_YEAR:     C.GetDayOfYear,
	expr.GET_MONTH_OF_YEAR:   C.GetMonthOfYear,
	expr.GET_QUARTER_OF_YEAR: C.GetQuarterOfYear,
	expr.GET_HLL_VALUE:       C.GetHLLValue,
	expr.ARRAY_LENGTH:        C.ArrayLength,
}

// BinaryExprTypeToCFunctorType maps from binary operator to C BinaryFunctorType
var BinaryExprTypeToCFunctorType = map[expr.Token]C.enum_BinaryFunctorType{
	expr.AND:              C.And,
	expr.OR:               C.Or,
	expr.EQ:               C.Equal,
	expr.NEQ:              C.NotEqual,
	expr.LT:               C.LessThan,
	expr.LTE:              C.LessThanOrEqual,
	expr.GT:               C.GreaterThan,
	expr.GTE:              C.GreaterThanOrEqual,
	expr.ADD:              C.Plus,
	expr.SUB:              C.Minus,
	expr.MUL:              C.Multiply,
	expr.DIV:              C.Divide,
	expr.MOD:              C.Mod,
	expr.BITWISE_AND:      C.BitwiseAnd,
	expr.BITWISE_OR:       C.BitwiseOr,
	expr.BITWISE_XOR:      C.BitwiseXor,
	expr.FLOOR:            C.Floor,
	expr.CONVERT_TZ:       C.Plus,
	expr.ARRAY_CONTAINS:   C.ArrayContains,
	expr.ARRAY_ELEMENT_AT: C.ArrayElementAt,
	// TODO: expr.BITWISE_LEFT_SHIFT ?
	// TODO: expr.BITWISE_RIGHT_SHIFT ?
}

type rootAction func(functorType uint32, stream unsafe.Pointer, device int, inputs []C.InputVector, exp expr.Expr)

func makeForeignColumnInput(columnIndex int, recordIDs unsafe.Pointer, table foreignTable, timezoneLookup unsafe.Pointer, timezoneLookupSize int) C.InputVector {
	_ = "STUB: not implemented"
	return *new(C.InputVector)
}

func makeDefaultValue(value memCom.DataValue) C.DefaultValue {
	_ = "STUB: not implemented"
	return *new(C.DefaultValue)
}

// Otherwise it's the default value type we don't support yet, setting it to null to be safe.

func makeVectorPartySlice(column deviceVectorPartySlice) C.VectorPartySlice {
	_ = "STUB: not implemented"
	return *new(C.VectorPartySlice)
}

func makeArrayVectorPartySlice(column deviceVectorPartySlice) C.ArrayVectorPartySlice {
	_ = "STUB: not implemented"
	return *new(C.ArrayVectorPartySlice)
}

func makeVectorPartySliceInput(column deviceVectorPartySlice) C.InputVector {
	_ = "STUB: not implemented"
	return *new(C.InputVector)
}

func makeArrayVectorPartySliceInput(column deviceVectorPartySlice) C.InputVector {
	_ = "STUB: not implemented"
	return *new(C.InputVector)
}

func makeConstantInput(val interface{}, isValid bool) C.InputVector {
	_ = "STUB: not implemented"
	return *new(C.InputVector)
}

func makeScratchSpaceInput(values unsafe.Pointer, nulls unsafe.Pointer, dataType C.enum_DataType) C.InputVector {
	_ = "STUB: not implemented"
	return *new(C.InputVector)
}

func makeMeasureVectorOutput(measureVector unsafe.Pointer, outputDataType C.enum_DataType, aggFunc C.enum_AggregateFunction) C.OutputVector {
	_ = "STUB: not implemented"
	return *new(C.OutputVector)
}

func makeDimensionVectorOutput(dimensionVector unsafe.Pointer, valueOffset, nullOffset int, dataType C.enum_DataType) C.OutputVector {
	_ = "STUB: not implemented"
	return *new(C.OutputVector)
}

func makeScratchSpaceOutput(values unsafe.Pointer, nulls unsafe.Pointer, dataType C.enum_DataType) C.OutputVector {
	_ = "STUB: not implemented"
	return *new(C.OutputVector)
}

func makeDimensionVector(valueVector, hashVector, indexVector unsafe.Pointer, numDims common.DimCountsPerDimWidth, vectorCapacity int) C.DimensionVector {
	_ = "STUB: not implemented"
	return *new(C.DimensionVector)
}

func getOutputDataType(exprType expr.Type, outputWidthInByte int) C.enum_DataType {
	_ = "STUB: not implemented"
	return *new(C.enum_DataType)
}

// For reducing the measure output iterator cardinality.

func initIndexVector(vector unsafe.Pointer, start, size int, stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	return
}

func (bc *oopkBatchContext) filterAction(functorType uint32, stream unsafe.Pointer, device int, inputs []C.InputVector, exp expr.Expr) {
	_ = "STUB: not implemented"
	return
}

// If current batch size is already 0, short circuit to avoid issuing a noop cuda call.

func (bc *oopkBatchContext) makeWriteToMeasureVectorAction(aggFunc C.enum_AggregateFunction, outputWidthInByte int) rootAction {
	_ = "STUB: not implemented"
	return *new(rootAction)
}

// If current batch size is already 0, short circuit to avoid issuing a noop cuda call.

// write measure out to measureVectorD[1] for hll query

func (bc *oopkBatchContext) makeWriteToDimensionVectorAction(valueOffset, nullOffset, prevResultSize int) rootAction {
	_ = "STUB: not implemented"
	return *new(rootAction)
}

// If current batch size is already 0, short circuit to avoid issuing a noop cuda call.

// move dimensionVectorD to the start position of current batch
// dimension vector start position + bc.resultSize * dataBytes
// null vector start position + bc.resultSize

func makeCuckooHashIndex(primaryKeyData memCom.PrimaryKeyData, deviceData unsafe.Pointer) C.CuckooHashIndex {
	_ = "STUB: not implemented"
	return *new(C.CuckooHashIndex)
}

func (bc *oopkBatchContext) prepareForeignRecordIDs(mainTableJoinColumnIndex int, joinTableID int, table foreignTable,
	stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	// If current batch size is already 0, short circuit to avoid issuing a noop cuda call.
	return
}

// processExpression does AST tree dfs traversal and apply root action on the root level,
// rootAction includes filterAction, writeToDimensionVectorAction and makeWriteToMeasureVectorAction
func (bc *oopkBatchContext) processExpression(exp, parentExp expr.Expr, tableScanners []*TableScanner, foreignTables []*foreignTable,
	stream unsafe.Pointer, device int, action rootAction) C.InputVector {
	_ = "STUB: not implemented"
	return *new(C.InputVector)
}

// main table

func makeGeoShapeBatch(shapesLatLongs devicePointer, numShapes, totalNumPoints int) C.GeoShapeBatch {
	_ = "STUB: not implemented"
	return *new(C.GeoShapeBatch)
}

func (bc *oopkBatchContext) makeGeoPointInputVector(pointTableID int, pointColumnIndex int, foreignTables []*foreignTable) C.InputVector {
	_ = "STUB: not implemented"
	return *new(C.InputVector)
}

func (bc *oopkBatchContext) writeGeoShapeDim(geo *geoIntersection,
	outputPredicate devicePointer, dimValueOffset, dimNullOffset int, sizeBeforeGeoFilter, prevResultSize int, stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	return
}

// geo dimension always take 1 byte and has type uint8
// compiler should have checked the number of geo shapes for join is less than 256

// move dimensionVectorD to the start position of current batch
// dimension vector start position + prevResultSize * dataBytes
// null vector start position + prevResultSize

func (bc *oopkBatchContext) geoIntersect(geo *geoIntersection, pointColumnIndex int,
	foreignTables []*foreignTable,
	outputPredicte devicePointer, stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	// If current batch size is already 0, short circuit to avoid issuing a noop cuda call.
	return
}

func (bc *oopkBatchContext) hll(numDims common.DimCountsPerDimWidth, isLastBatch bool, stream unsafe.Pointer, device int) (
	hllVector, dimRegCount devicePointer, hllVectorSize int64) {
	_ = "STUB: not implemented"
	return *new(devicePointer), *new(devicePointer), 0
}

// TODO: we also need a way to report this allocation in C++ code. Maybe can be done via calling a golang function from c++

func (bc *oopkBatchContext) sortByKey(numDims common.DimCountsPerDimWidth, stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	return
}

// sort the previous result with current batch together

func (bc *oopkBatchContext) reduceByKey(numDims common.DimCountsPerDimWidth, valueWidth int, aggFunc C.enum_AggregateFunction, stream unsafe.Pointer,
	device int) {
	_ = "STUB: not implemented"
	return
}

func (bc *oopkBatchContext) hashReduce(numDims common.DimCountsPerDimWidth, valueWidth int, aggFunc C.enum_AggregateFunction, stream unsafe.Pointer,
	device int) {
	_ = "STUB: not implemented"
	return
}

func (bc *oopkBatchContext) expand(numDims common.DimCountsPerDimWidth, stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	return
}

func (bc *oopkBatchContext) getDataTypeLength(dataType C.enum_DataType) int {
	_ = "STUB: not implemented"
	return 0
}

// allocate new stack frame and append to the end of the stack
func (bc *oopkBatchContext) allocateStackFrame(dataType C.enum_DataType) (values, nulls devicePointer) {
	_ = "STUB: not implemented"
	return *new(devicePointer), *new(devicePointer)
}

// width bytes * bc.size (value buffer) + 1 byte * bc.size (null buffer)

// append output buffer to the end

// shrink stack by one, but keep top element as is
func (bc *oopkBatchContext) shrinkStackFrame() { _ = "STUB: not implemented"; return }

// swap last two elements

// pop last element

func (qc *AQLQueryContext) createCutoffTimeFilter(cutoff uint32) expr.Expr {
	_ = "STUB: not implemented"
	return *new(expr.Expr)
}

// time column is always 0

// doCGoCall does the cgo call by converting CGoCallResHandle to C.int and *C.char and calls doCGoCall.
// The reason to have this wrapper is because CGo types are bound to package name, thereby even C.int are different types
// under different packages.
func doCGoCall(f func() C.CGoCallResHandle) uintptr { _ = "STUB: not implemented"; return 0 }

// bootstrapDevice is the go wrapper of BootstrapDevice. It will panic and crash the server if any exceptions are thrown
// in this function.
func bootstrapDevice() { _ = "STUB: not implemented"; return }
