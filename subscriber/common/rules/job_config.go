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

package rules

import (
	"github.com/uber/aresdb/controller/models"

	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/subscriber/config"
	"go.uber.org/fx"
)

// Module configures JobConfigs.
var Module = fx.Options(
	fx.Provide(
		NewJobConfigs,
	),
)

// Params defines the base objects for jobConfigs.
type Params struct {
	fx.In

	ServiceConfig config.ServiceConfig
}

// Result defines the objects that the rules module provides.
type Result struct {
	fx.Out

	JobConfigs JobConfigs
}

// JobConfigs contains configuration and information for all jobs and therir destination ares clusters.
type JobConfigs map[string]JobAresConfig

// JobAresConfig contains configuration and information for each Ares cluster and job configuration.
type JobAresConfig map[string]*JobConfig

// JobConfig wraps job config controller
type JobConfig struct {
	models.JobConfig
	// maps from column name to columnID for convenience
	columnDict      map[string]int
	destinations    map[string]*DestinationConfig
	transformations map[string]*TransformationConfig
	primaryKeys     map[string]int
	primaryKeyBytes int
}

// DestinationConfig defines the configuration needed to save data in ares
type DestinationConfig struct {
	// Table is ares table
	Table string `json:"table" yaml:"table"`
	// Column is ares table column name
	Column string `json:"column" yaml:"column"`
	// UpdateMode is column's upsert mode
	UpdateMode memCom.ColumnUpdateMode `json:"update_mode,omitempty" yaml:"update_mode,omitempty"`
}

// TransformationConfig defiines the configuration needed to generate a specific transformation
type TransformationConfig struct {
	// Type of transformationConfig to apply for the column,
	// like timestamp, uuid, uuid_hll etc
	Type string `json:"type" yaml:"type"`
	// Source is the field name to read the value from
	Source string `json:"source" yaml:"source"`
	// Default value to use if value is empty
	Default string `json:"default" yaml:"default"`
	// Context help complex transformations to define information
	// needed to parse the values
	Context map[string]string
}

// Assignment defines the job assignment of the ares-subscriber
type Assignment struct {
	// Subscriber is ares-subscriber instance name
	Subscriber string `json:"subscriber"`
	// Jobs is a list of jobConfigs for the ares-subscriber instance
	Jobs []*JobConfig `json:"jobs"`
	// AresClusters is a table of aresClusters for the ares-subscriber instance
	AresClusters map[string]config.SinkConfig `json:"instances"`
}

// NewJobConfigs creates JobConfigs object
func NewJobConfigs(params Params) (Result, error) {
	_ = "STUB: not implemented"
	return *new(Result), nil
}

// NewAssignmentFromController parse controller assignment and create Assignment rule
func NewAssignmentFromController(from *models.IngestionAssignment) (*Assignment, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetDestinations returns a job's destination definition
func (j *JobConfig) GetDestinations() map[string]*DestinationConfig {
	_ = "STUB: not implemented"
	return nil

	// GetTranformations returns a job's tranformation definition
}

func (j *JobConfig) GetTranformations() map[string]*TransformationConfig {
	_ = "STUB: not implemented"
	return nil

	// GetPrimaryKeys returns a job's primaryKeys definition
}

func (j *JobConfig) GetPrimaryKeys() map[string]int { _ = "STUB: not implemented"; return nil }

// SetPrimaryKeyBytes sets the number of bytes needed by primaryKey
func (j *JobConfig) SetPrimaryKeyBytes(primaryKeyBytes int) { _ = "STUB: not implemented"; return }

// AppendPrimaryKeyBytes returns the number of bytes needed by primaryKey
func (j *JobConfig) GetPrimaryKeyBytes() int { _ = "STUB: not implemented"; return 0 }

// GetColumnDict returns a job's columnDict definition
func (j *JobConfig) GetColumnDict() map[string]int { _ = "STUB: not implemented"; return nil }

// PopulateAresTableConfig populates information into jobConfig fields
func (j *JobConfig) PopulateAresTableConfig() error {
	_ = "STUB: not implemented"
	// set primaryKeys and primaryKeyBytes
	return nil
}

// set destinations and transformations

func (j *JobConfig) getUpdateMode(column string) memCom.ColumnUpdateMode {
	_ = "STUB: not implemented"
	return *new(memCom.ColumnUpdateMode)
}

// AddLocalJobConfig creates a list of jobConfigs from local configuration file
func AddLocalJobConfig(serviceConfig config.ServiceConfig, jobConfigs JobConfigs) error {
	_ = "STUB: not implemented"
	return nil
}

// iterate all active jobs configured at local

// set the job configure by loading its configuration file

// iterate all active ares cluster to set jobConfig for each of them

// newJobConfig creates a jobConfig from json
func newJobConfig(value []byte) (*JobConfig, error) { _ = "STUB: not implemented"; return nil, nil }

// CloneJobConfig deep copy jobConfig
func CloneJobConfig(src *JobConfig, serviceConfig config.ServiceConfig, aresCluster string) (*JobConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// copy column defaultValue

// copy destinations map

// copy transformations map

// copy primaryKeys map[string]interface{}

// parseUpdateMode converts update mode string to memCom.ColumnUpdateMode
func parseUpdateMode(modeStr string) memCom.ColumnUpdateMode {
	_ = "STUB: not implemented"
	return *new(memCom.ColumnUpdateMode)
}
