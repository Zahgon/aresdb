//	Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package kvstore

import (
	"github.com/m3db/m3/src/cluster/client"
	"github.com/m3db/m3/src/cluster/client/etcd"
	"github.com/m3db/m3/src/cluster/kv"
	"github.com/m3db/m3/src/cluster/services"
	"go.uber.org/zap"
)

// EtcdClient represents etcd client
type EtcdClient struct {
	Zone          string
	Environment   string
	ServiceName   string
	ClusterClient client.Client
	TxnStore      kv.TxnStore
	Services      services.Services
}

// EtcdConfig represents etcd config
type EtcdConfig struct {
	etcd.Configuration `yaml:",inline"`
	Enabled            bool `yaml:"enabled" json:"enabled"`
}

// NewEtcdClient returns a new EtcdClient
func NewEtcdClient(logger *zap.SugaredLogger, config EtcdConfig) *EtcdClient {
	_ = "STUB: not implemented"
	return nil
}
