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

package datanode

import (
	"net/http"
	"sync"
	"time"

	"github.com/uber/aresdb/common"

	"github.com/m3db/m3/src/cluster/services"
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/api"
	"github.com/uber/aresdb/cluster/shard"
	"github.com/uber/aresdb/cluster/topology"
	mutatorsCom "github.com/uber/aresdb/controller/mutators/common"
	"github.com/uber/aresdb/diskstore"
	"github.com/uber/aresdb/memstore"
	metaCom "github.com/uber/aresdb/metastore/common"
	"github.com/uber/aresdb/redolog"
	"google.golang.org/grpc"
)

// dataNode includes metastore, memstore and diskstore
type dataNode struct {
	sync.RWMutex

	hostID          string
	startedAt       time.Time
	shardSet        shard.ShardSet
	clusterServices services.Services

	topoInitializer    topology.Initializer
	topology           topology.Topology
	numShardsInCluster int
	enumReader         mutatorsCom.EnumReader
	metaStore          metaCom.MetaStore
	memStore           memstore.MemStore
	diskStore          diskstore.DiskStore

	opts     Options
	logger   common.Logger
	metrics  datanodeMetrics
	handlers datanodeHandlers

	bootstrapManager     BootstrapManager
	redoLogManagerMaster *redolog.RedoLogManagerMaster
	grpcServer           *grpc.Server

	mapWatch topology.MapWatch
	close    chan struct{}

	readyCh chan struct{}
}

type datanodeHandlers struct {
	schemaHandler      *api.SchemaHandler
	enumHandler        *api.EnumHandler
	queryHandler       *api.QueryHandler
	dataHandler        *api.DataHandler
	nodeModuleHandler  http.Handler
	debugStaticHandler http.Handler
	debugHandler       *api.DebugHandler
	healthCheckHandler *api.HealthCheckHandler
	swaggerHandler     http.Handler
}

type datanodeMetrics struct {
	restartTimer tally.Timer
}

// NewDataNode creates a new data node
func NewDataNode(
	hostID string,
	topoInitializer topology.Initializer,
	enumReader mutatorsCom.EnumReader,
	opts Options) (DataNode, error) {
	_ = "STUB: not implemented"
	return *new(DataNode), nil
}

// Open data node for serving
func (d *dataNode) Open() error { _ = "STUB: not implemented"; return nil }

//1. start schema watch

// memstore fetch local disk schema

// 2. start debug server

// initialize topology, block wait for first topology in etcd

// 3. first shard assignment

// 5. start scheduler

// disable archiving during redolog replay

// this will start scheduler of all jobs except archiving, archiving will be started individually

// 6. start table addition watch

// 7. start active topology watch

// 8. start analyzing shard availability

// 9. start analyzing server readiness

// 10. start bootstrap retry watch

func (d *dataNode) startSchemaWatch() { _ = "STUB: not implemented"; return }

// TODO better to reuse the code directly in controller to talk to etcd

// immediate initial fetch

func (d *dataNode) Close() { _ = "STUB: not implemented"; return }

func (d *dataNode) startDebugServer() { _ = "STUB: not implemented"; return }

func (d *dataNode) startTableAdditionWatch() { _ = "STUB: not implemented"; return }

func (d *dataNode) startActiveTopologyWatch() { _ = "STUB: not implemented"; return }

// assign empty shard set when host does not appear in placement

// checkShardReadiness check which of the shards are ready (meaning all tables within the shard are bootstrapped)
// this function will partition the input shards into two parts with all ready shards in the first part
// and return the number of ready shards
// eg. given input shards [0, 1, 2, 3, 4, 5, 6, 7], if all tables in shard 2, 6 are bootstrapped
// when the algorithm is finished,
// the input shards slice will become [2, 6, 0, 1, 3, 4, 5, 7],
// and numReadyShards returned is 2
func (d *dataNode) checkShardReadiness(tables []string, shards []uint32) (numReadyShards int) {
	_ = "STUB: not implemented"
	return 0
}

func (d *dataNode) startAnalyzingServerReadiness() { _ = "STUB: not implemented"; return }

// condition for serving readiness
// 1. no shards owned by server

// 2. no nonInitializing (available/leaving) shards waiting for bootstrap

// 3. all nonInitializing (available/leaving) shards are bootstrapped
// and all dimension table shards are bootstrapped (only one shard for dim table)

func (d *dataNode) startAnalyzingShardAvailability() { _ = "STUB: not implemented"; return }

// initializing shards are shards with Initializing State
// nonInitializing shards are shards with Available and Leaving State
// we always bootstrap nonInitializing shards first
// and wait for nonInitializing shards's readiness before serving traffic

// snapshot fact tables

// Options returns the database options.
func (d *dataNode) Options() Options {
	_ = "STUB: not implemented"

	// ShardSet returns the set of shards currently associated with this datanode.
	return *new(Options)
}

func (d *dataNode) ShardSet() shard.ShardSet {
	_ = "STUB: not implemented"
	return *new(shard.ShardSet)
}

func (d *dataNode) ID() string { _ = "STUB: not implemented"; return "" }

func (d *dataNode) Serve() {
	_ = "STUB: not implemented"
	// wait for server is ready to serve
	return
}

// start advertising to the cluster

// enable archiving jobs

// start server

// Support CORS calls.

// record time from data node started to actually serving

// start batch reporter

func (d *dataNode) advertise() { _ = "STUB: not implemented"; return }

func (d *dataNode) addTable(table string) { _ = "STUB: not implemented"; return }

// dimension table defaults shard to zero
// new table does not need to copy data from peer, but need to purge old data

// new table does not need to copy data from peer, but need to purge old data

func (d *dataNode) assignShardSet(shardSet shard.ShardSet) { _ = "STUB: not implemented"; return }

// process fact tables first

// when needPeerCopy is true, we also need to purge old data before adding new shard

// add/remove dimension tables with the following rules:
// 1. add dimension tables when first shard is assigned to the data node
// 2. remove dimension tables when the last shard is removed from the data node
// 3. copy dimension table data from peer when all assigned new shards are initializing shards

// only need to copy data from peer when all new shards are initializing shards
// meaning no available/leaving shards ever owned by this data node

// only copy data from peer for dimension table
// when from zero shards to all initialing shards
// when needPeerCopy is true, we also need to purge old data before adding new shard

// GetOwnedShards returns all shard ids the datanode owns
func (d *dataNode) GetOwnedShards() []int { _ = "STUB: not implemented"; return nil }

func newDatanodeMetrics(scope tally.Scope) datanodeMetrics {
	_ = "STUB: not implemented"
	return *new(datanodeMetrics)
}

func (d *dataNode) newHandlers() datanodeHandlers {
	_ = "STUB: not implemented"
	return *new(datanodeHandlers)
}

// mixed handler for both grpc and traditional http
func mixedHandler(grpcServer *grpc.Server, httpHandler http.Handler) http.Handler {
	_ = "STUB: not implemented"
	return *new(http.Handler)
}

func (d *dataNode) startBootstrapRetryWatch() { _ = "STUB: not implemented"; return }
