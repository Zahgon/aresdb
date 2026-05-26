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

package utils

// TableDataSource defines the interface a data source need to implement so that we can render
// a tabular representation from the data source. We get number of columns from the length of
// column header. the data source itself should ensure that for each get value call with row and
// col within [0,numRows) and [0, numCols), it should return valid value and should not panic.
type TableDataSource interface {
	NumRows() int
	GetValue(row, col int) interface{}
	ColumnHeaders() []string
}

func getFormatModifier(value interface{}) string { _ = "STUB: not implemented"; return "" }

func expandColumnWidth(columnWidths []int, value interface{}, idx int) {
	_ = "STUB: not implemented"
	return
}

func sprintfStrings(format string, strs []string) string { _ = "STUB: not implemented"; return "" }

// WriteTable renders a tabular representation from underlying data source.
// If there is no column for this data source, it will return an empty string.
// All elements of the table will be right justify (left padding). Column
// splitter is "|" for now.
func WriteTable(dataSource TableDataSource) string { _ = "STUB: not implemented"; return "" }

// Return empty string if no columns.

// Compute column widths.

// Then compare with the length of each value.

// string buffer for final result.

// Prepare format for header.

// Write column header.

// Prepare format for rows.

// get formatter of first row.

// Write rows.
