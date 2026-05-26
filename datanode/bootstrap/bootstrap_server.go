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

package bootstrap

import (
	"context"
	"errors"
	"sync"
	"time"

	pb "github.com/uber/aresdb/datanode/generated/proto/rpc"
	"github.com/uber/aresdb/diskstore"
	"github.com/uber/aresdb/metastore/common"
)

const (
	chunkSize       = 32 * 1024
	bufferSize      = 32 * 1024
	recycleInterval = 5 * time.Second
)

var (
	errNoCallerID       = errors.New("caller node id not set in request")
	errNoSessionID      = errors.New("session id not set in request")
	errInvalidSessionID = errors.New("invalid session id")
	errInvalidRequset   = errors.New("invalid request, table/shard not match")
	errSessionExisting  = errors.New("The request table/shard already have session running from the same node")
)

type PeerDataNodeServerImpl struct {
	sync.RWMutex

	// session id generator
	sequenceID int64

	metaStore common.MetaStore
	diskStore diskstore.DiskStore

	// session id to sessionInfo map
	sessions map[int64]*sessionInfo
	// tracking of all sessions for each table/shard
	tableShardSessions map[tableShardPair][]int64
}

type tableShardPair struct {
	table   string
	shardID uint32
}

type sessionInfo struct {
	sessionID    int64
	table        string
	shardID      uint32
	nodeID       string
	addr         string
	lastLiveTime time.Time
	ttl          int64
}

func NewPeerDataNodeServer(metaStore common.MetaStore, diskStore diskstore.DiskStore) pb.PeerDataNodeServer {
	_ = "STUB: not implemented"
	return *new(pb.PeerDataNodeServer)
}

// AcquireToken is to check if any bootstrap is running in the table/shard
// if no bootstrap session is running on the table/shard, it will increase the token count, and return true
// the caller need to release the usage by calling ReleaseToken
func (p *PeerDataNodeServerImpl) AcquireToken(tableName string, shardID uint32) bool {
	_ = "STUB: not implemented"
	return false

	// lazy clean the obsolete orphan sessions
}

// AcquireToken release the token count, must call this when call AcquireToken success
func (p *PeerDataNodeServerImpl) ReleaseToken(tableName string, shardID uint32) {
	_ = "STUB: not implemented"
	// nothing to do for now
	return
}

// getNextSequence create new session id
func (p *PeerDataNodeServerImpl) getNextSequence() int64 { _ = "STUB: not implemented"; return 0 }

// Health return the healthiness status of the data server
func (p *PeerDataNodeServerImpl) Health(ctx context.Context, req *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// StartSession create new session for one table/shard/node, only One session can be established on one table/shard from one node
func (p *PeerDataNodeServerImpl) StartSession(ctx context.Context, req *pb.StartSessionRequest) (*pb.Session, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (p *PeerDataNodeServerImpl) checkReqExist(s *sessionInfo) error {
	_ = "STUB: not implemented"
	return nil
}

// KeepAlive is like client/server ping process, to notify health about each other
func (p *PeerDataNodeServerImpl) KeepAlive(stream pb.PeerDataNode_KeepAliveServer) error {
	_ = "STUB: not implemented"
	return nil
}

// update last live time

// FetchTableShardMetaData to retrieve all metadata for one table/shard
func (p *PeerDataNodeServerImpl) FetchTableShardMetaData(ctx context.Context, req *pb.TableShardMetaDataRequest) (*pb.TableShardMetaData, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// dimension table

// fact table

// adjust start/end batchID according to local retention setting and request
// we'll take the intersection batches

func (p *PeerDataNodeServerImpl) FetchVectorPartyRawData(req *pb.VectorPartyRawDataRequest, stream pb.PeerDataNode_FetchVectorPartyRawDataServer) error {
	_ = "STUB: not implemented"
	return nil
}

// clean the EOF error

// in macro second

// BenchmarkFileTransfer is used to benchmark testing, we can remove later TODO
func (p *PeerDataNodeServerImpl) BenchmarkFileTransfer(req *pb.BenchmarkRequest, stream pb.PeerDataNode_BenchmarkFileTransferServer) error {
	_ = "STUB: not implemented"
	return nil
}

// in macro second

func (p *PeerDataNodeServerImpl) validateSessionSource(sessionID int64, nodeID string) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *PeerDataNodeServerImpl) validateRequest(sessionID int64, nodeID string, table string, shard uint32) error {
	_ = "STUB: not implemented"
	return nil
}

// record new requested session
func (p *PeerDataNodeServerImpl) addSession(session *sessionInfo) {
	_ = "STUB: not implemented"
	return
}

// closeSession remove session from memory
func (p *PeerDataNodeServerImpl) cleanSession(sessionID int64, needLock bool) {
	_ = "STUB: not implemented"
	return
}

// retrieve session info using session id
func (p *PeerDataNodeServerImpl) getSession(sessionID int64) (*sessionInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// check if request table/shard is valid
func (p *PeerDataNodeServerImpl) validateTable(tableName string, shardID uint32) error {
	_ = "STUB: not implemented"
	// check if table exists
	return nil
}

//  TODO check table shard ownership from topology

func logInfoMsg(s *sessionInfo, msg string, fields ...interface{}) {
	_ = "STUB: not implemented"
	return
}

func logErrorMsg(s *sessionInfo, err error, msg string, fields ...interface{}) {
	_ = "STUB: not implemented"
	return
}
