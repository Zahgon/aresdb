// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package datanode

import (
	"sync"
	"time"

	"github.com/uber/aresdb/cluster/topology"
	"github.com/uber/aresdb/datanode/bootstrap"
	"github.com/uber/aresdb/datanode/client"
)

// bootstrapManagerImpl is the implementation of the interface databaseBootstrapManager
type bootstrapManagerImpl struct {
	sync.RWMutex

	opts                        bootstrap.Options
	origin                      string
	peerSource                  client.PeerSource
	bootstrapable               bootstrap.Bootstrapable
	state                       bootstrap.BootstrapState
	hasPending                  bool
	bootstrapTableShards        map[string]bootstrap.BootstrapDetails
	lastBootstrapCompletionTime time.Time
	topo                        topology.Topology
}

// NewBootstrapManager creates bootstrap manager
func NewBootstrapManager(origin string,
	bootstrappable bootstrap.Bootstrapable,
	bootstrapOpts bootstrap.Options,
	topo topology.Topology,
) BootstrapManager {
	_ = "STUB: not implemented"
	return *new(BootstrapManager)
}

func (m *bootstrapManagerImpl) IsBootstrapped() bool { _ = "STUB: not implemented"; return false }

func (m *bootstrapManagerImpl) LastBootstrapCompletionTime() (time.Time, bool) {
	_ = "STUB: not implemented"
	return *new(time.Time), false
}

func (m *bootstrapManagerImpl) Bootstrap() error { _ = "STUB: not implemented"; return nil }

// NB(r): Already bootstrapping, now a consequent bootstrap
// request comes in - we queue this up to bootstrap again
// once the current bootstrap has completed.
// This is an edge case that can occur if during either an
// initial bootstrap or a resharding bootstrap if a new
// reshard occurs and we need to bootstrap more shards.

// Keep performing bootstraps until none pending

// New bootstrap calls should now enqueue another pending bootstrap

func (m *bootstrapManagerImpl) bootstrap() error { _ = "STUB: not implemented"; return nil }

func newInitialTopologyState(topo topology.Topology) *topology.StateSnapshot {
	_ = "STUB: not implemented"
	return nil
}
