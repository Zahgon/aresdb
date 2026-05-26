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

package message

import (
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/client"
	metaCom "github.com/uber/aresdb/metastore/common"
	"github.com/uber/aresdb/subscriber/common/rules"
	"github.com/uber/aresdb/subscriber/common/sink"
	"github.com/uber/aresdb/subscriber/config"
)

// Parser holds all resources needed to parse one message
// into one or multiple row objects with respect to different destinations
type Parser struct {
	// ServiceConfig is ares-subscriber configure
	ServiceConfig config.ServiceConfig
	// JobName is job name
	JobName string
	// Cluster is ares cluster name
	Cluster string
	// destinations each message will be parsed and written into
	Destination sink.Destination
	// Transformations are keyed on the output column name
	Transformations map[string]*rules.TransformationConfig
	scope           tally.Scope
}

// NewParser will create a Parser for given JobConfig
func NewParser(jobConfig *rules.JobConfig, serviceConfig config.ServiceConfig) *Parser {
	_ = "STUB: not implemented"
	return nil
}

func (mp *Parser) populateDestination(jobConfig *rules.JobConfig) {
	_ = "STUB: not implemented"
	return
}

// sort column names in destination for consistent query order

// ParseMessage will parse given message to fit the destination
func (mp *Parser) ParseMessage(msg map[string]interface{}, destination sink.Destination) (client.Row, error) {
	_ = "STUB: not implemented"
	return *new(client.Row), nil
}

// IsMessageValid checks if the message is valid
func (mp *Parser) IsMessageValid(msg map[string]interface{}, destination sink.Destination) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckPrimaryKeys returns error if the value of primary key column is nil
func (mp *Parser) CheckPrimaryKeys(destination sink.Destination, row client.Row) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckTimeColumnExistence checks if time column is missing for fact table
func (mp *Parser) CheckTimeColumnExistence(schema *metaCom.Table, columnDict map[string]int,
	destination sink.Destination, row client.Row) error {
	_ = "STUB: not implemented"
	return nil
}

func (mp *Parser) extractSourceFieldValue(msg map[string]interface{}, fieldName string) interface{} {
	_ = "STUB: not implemented"
	return nil
}

func (mp *Parser) getValue(msg map[string]interface{}, fieldName string) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetFuncName get the function name of the calling function
func GetFuncName() string { _ = "STUB: not implemented"; return "" }
