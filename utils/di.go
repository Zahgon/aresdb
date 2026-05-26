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

import (
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/common"
)

// stores all common components together to avoid scattered references.
var (
	logger          common.Logger
	queryLogger     common.Logger
	reporterFactory *ReporterFactory
	config          common.AresServerConfig
)

// init loads default implementations of common components for unit tests' purpose.
func init() {
	ResetDefaults()
}

// ResetDefaults reset default config, logger and metrics settings
func ResetDefaults() { _ = "STUB: not implemented"; return }

// Init loads application specific common components settings.
func Init(c common.AresServerConfig, l common.Logger, ql common.Logger, s tally.Scope) {
	_ = "STUB: not implemented"
	return
}

// GetLogger returns the logger.
func GetLogger() common.Logger {
	_ = "STUB: not implemented"

	// GetQueryLogger returns the logger for query.
	return *new(common.Logger)
}

func GetQueryLogger() common.Logger {
	_ = "STUB: not implemented"

	// GetRootReporter returns the root metrics reporter.
	return *new(common.Logger)
}

func GetRootReporter() *Reporter { _ = "STUB: not implemented"; return nil }

// GetReporter returns reporter given tableName and shardID. If the corresponding
// reporter cannot be found. It will return the root scope.
func GetReporter(tableName string, shardID int) *Reporter { _ = "STUB: not implemented"; return nil }

// AddTableShardReporter adds a reporter for the given table and shards. It should
// be called when bootstrap the table shards or shard ownership changes.
func AddTableShardReporter(tableName string, shardID int) { _ = "STUB: not implemented"; return }

// DeleteTableShardReporter deletes the reporter for the given table and shards. It should
// be called when the table shard no longer belongs to current node.
func DeleteTableShardReporter(tableName string, shardID int) { _ = "STUB: not implemented"; return }

// GetConfig returns the application config.
func GetConfig() common.AresServerConfig {
	_ = "STUB: not implemented"
	return *new(common.AresServerConfig)
}
