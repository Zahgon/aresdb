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

package sink

import (
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/client"
	controllerCli "github.com/uber/aresdb/controller/client"
	"github.com/uber/aresdb/subscriber/common/rules"
	"github.com/uber/aresdb/subscriber/config"
)

// AresDatabase is an implementation of Database interface for saving data to ares
type AresDatabase struct {
	ServiceConfig config.ServiceConfig
	JobConfig     *rules.JobConfig
	Scope         tally.Scope
	ClusterName   string
	Connector     client.Connector
}

// NewAresDatabase initialize an AresDatabase cluster
func NewAresDatabase(
	serviceConfig config.ServiceConfig, jobConfig *rules.JobConfig, cluster string,
	sinkCfg config.SinkConfig, aresControllerClient controllerCli.ControllerClient) (Sink, error) {
	_ = "STUB: not implemented"
	return *new(Sink), nil
}

// Shutdown will clean up resources that needs to be cleaned up
func (db *AresDatabase) Shutdown() { _ = "STUB: not implemented"; return }

// Save saves a batch of row objects into a destination
func (db *AresDatabase) Save(destination Destination, rows []client.Row) error {
	_ = "STUB: not implemented"
	return nil
}

// Cluster returns the DB cluster name
func (db *AresDatabase) Cluster() string { _ = "STUB: not implemented"; return "" }
