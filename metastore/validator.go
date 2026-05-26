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

package metastore

import (
	"github.com/uber/aresdb/metastore/common"
)

// TableSchemaValidator validates it a new table schema is valid, given existing schema
type TableSchemaValidator interface {
	SetOldTable(table common.Table)
	SetNewTable(table common.Table)
	Validate() error
}

// NewTableSchameValidator returns a new TableSchemaValidator. Pass nil for oldTable if none exists
func NewTableSchameValidator() TableSchemaValidator {
	_ = "STUB: not implemented"
	return *new(TableSchemaValidator)
}

type tableSchemaValidatorImpl struct {
	newTable *common.Table
	oldTable *common.Table
}

func (v *tableSchemaValidatorImpl) SetOldTable(table common.Table) {
	_ = "STUB: not implemented"
	return
}

func (v *tableSchemaValidatorImpl) SetNewTable(table common.Table) {
	_ = "STUB: not implemented"
	return
}

func (v tableSchemaValidatorImpl) Validate() (err error) { _ = "STUB: not implemented"; return nil }

// ValidateHLLConfig validates hll config
func validateColumnHLLConfig(c common.Column) error { _ = "STUB: not implemented"; return nil }

// checks performed:
//
//		table has at least 1 valid column
//		table has at least 1 valid primary key column
//	 fact table must have a time column as first column
//		fact table must have sort columns that are valid
//		each column have valid data type and default value
//		sort columns cannot have duplicate columnID
//		primary key columns cannot have duplicate columnID
//		column name cannot duplicate
//	 check hll cannot be enabled on time column
//	 check column configs
func (v tableSchemaValidatorImpl) validateIndividualSchema(table *common.Table, creation bool) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// validate data type

// validate hll config

// time column does not allow hll config

// checks performed
//
//		check that new table is valid table
//		check new table has larger version number
//		check no changes on immutable fields (table name, type, pk)
//		check updates on columns and sort columns are valid
//	 check allowMissingEventTime cannot be changed from true to false
//	 check hllConfig cannot be changed
func (v tableSchemaValidatorImpl) validateSchemaUpdate(newTable, oldTable *common.Table) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// validate columns

// even with column deletion, or recreation, column id are not reused

// check that no column configs are modified, even for deleted columns

// end validate columns

// primary key columns

// sort columns

// ValidateDefaultValue validates default value against data type
func ValidateDefaultValue(valueStr, dataTypeStr string) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// BigEnum or Small Enum ares string values, no need to validate
