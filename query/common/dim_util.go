package common

import (
	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/query/expr"
)

// GetDimensionDataType gets DataType for given expr
func GetDimensionDataType(expression expr.Expr) memCom.DataType {
	_ = "STUB: not implemented"
	return *new(memCom.DataType)
}

// special handling element_at for array enum type

// GetDimensionDataBytes gets num bytes for given expr
func GetDimensionDataBytes(expression expr.Expr) int { _ = "STUB: not implemented"; return 0 }

// DimValueVectorSize returns the size of final dim value vector on host side.
func DimValResVectorSize(resultSize int, numDimsPerDimWidth DimCountsPerDimWidth) int {
	_ = "STUB: not implemented"
	return 0
}
