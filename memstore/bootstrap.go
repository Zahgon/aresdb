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
package memstore

import (
	"github.com/uber/aresdb/cluster/topology"
	"github.com/uber/aresdb/datanode/bootstrap"
	"github.com/uber/aresdb/datanode/client"
	"github.com/uber/aresdb/datanode/generated/proto/rpc"
	"github.com/uber/aresdb/utils"
)

// IsBootstrapped returns whether this table shard is bootstrapped.
func (shard *TableShard) IsBootstrapped() bool { _ = "STUB: not implemented"; return false }

// IsDiskDataAvailable returns whether the data is available on disk for table shard
func (shard *TableShard) IsDiskDataAvailable() bool { _ = "STUB: not implemented"; return false }

func (m *memStoreImpl) Bootstrap(
	peerSource client.PeerSource,
	origin string,
	topo topology.Topology,
	topoState *topology.StateSnapshot,
	options bootstrap.Options,
) error {
	_ = "STUB: not implemented"
	// snapshot table shards not bootstrapped
	return nil
}

// partition table shards based on whether it needs to copy data from peer
// so that we can start process those doesn't first

// if doesn't need peer copy, swap into the first half

// Bootstrap executes bootstrap for table shard
func (shard *TableShard) Bootstrap(
	peerSource client.PeerSource,
	origin string,
	topo topology.Topology,
	topoState *topology.StateSnapshot,
	options bootstrap.Options,
) error {
	_ = "STUB: not implemented"
	return nil

	// check whether shard is already bootstrapping
}

// find peer node for copy metadata and raw data

// Note: skip peer copy step when we found no peers. this is correct based on the assumption that
// if a shard replica does not find any peer in available/leaving state
// then the cluster should be in the initial phase of new placement created
// when we do replace/remove/add, the total number of shard replica remains the same
// so the number of initializing replica should match the number of leaving replica
// when we increase the number of replicas, we should always find existing available replica

// shuffle peer nodes randomly

// load metadata from disk

// preload snapshot or archive batches into memory

// preload all columns for fact table

// preload snapshot for dimension table

// start play redolog

type vpRawDataRequest struct {
	tableShardMeta *rpc.TableShardMetaData
	batchMeta      *rpc.BatchMetaData
	vpMeta         *rpc.VectorPartyMetaData
}

// fetchDataFromPeer fetch metadata and raw vector party data from peer
func (shard *TableShard) fetchDataFromPeer(
	peerID string,
	client rpc.PeerDataNodeClient,
	origin string,
	options bootstrap.Options,
) error {
	_ = "STUB: not implemented"
	return nil
}

// 1. fetch meta data

// 2. set metadata and trigger recovery

// 3. fetch raw vps

// capture batchMeta and vpMeta

// TODO: add checksum to vp file and vpMeta to avoid copying existing data on disk

func (shard *TableShard) fetchBatchMetaDataFromPeer(origin string, sessionID int64, client rpc.PeerDataNodeClient) (*rpc.TableShardMetaData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (shard *TableShard) createVectorPartyRawDataRequest(
	origin string,
	sessionID int64,
	tableMeta *rpc.TableShardMetaData,
	batchMeta *rpc.BatchMetaData,
	vpMeta *rpc.VectorPartyMetaData,
) (rawVPDataRequest *rpc.VectorPartyRawDataRequest, vpWriter utils.WriteSyncCloser, err error) {
	_ = "STUB: not implemented"
	return nil, *new(utils.WriteSyncCloser), nil
}

// fact table archive vp writer

// dimension table snapshot vp writer

func (shard *TableShard) fetchVectorPartyRawDataFromPeer(
	client rpc.PeerDataNodeClient,
	vpWriter utils.WriteSyncCloser,
	request *rpc.VectorPartyRawDataRequest,
) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (shard *TableShard) setBatchMetadata(tableShardMeta *rpc.TableShardMetaData, batchMeta *rpc.BatchMetaData) error {
	_ = "STUB: not implemented"
	return nil
}

func (shard *TableShard) setTableShardMetadata(tableShardMeta *rpc.TableShardMetaData) error {
	_ = "STUB: not implemented"
	// update kafka offsets
	return nil
}

// update archiving low water mark cutoff and backfill progress for fact table

// update snapshot pregress for dimension table

func (shard *TableShard) startStreamSession(peerID string, client rpc.PeerDataNodeClient, origin string, options bootstrap.Options) (sessionID int64, doneFn func(), err error) {
	_ = "STUB: not implemented"
	return 0, nil, nil
}

// send first keep alive request

// send loop

// receive loop

// server closed the stream

func (shard *TableShard) findBootstrapSource(
	origin string, topo topology.Topology, topoState *topology.StateSnapshot) ([]string, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This shard was not part of the topology when the bootstrapping
// process began.

// Don't take self into account

// Don't want to peer bootstrap from a node that has not yet completely
// taken ownership of the shard.

// Success cases - We can bootstrap from this host, which is enough to
// mark this shard as bootstrappable.
