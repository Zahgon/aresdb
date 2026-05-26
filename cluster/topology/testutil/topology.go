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

package testutil

import (
	"github.com/m3db/m3/src/cluster/shard"
	"github.com/uber/aresdb/cluster/topology"
)

// MustNewTopologyMap returns a new topology.Map with provided parameters.
// It's a utility method to make tests easier to write.
func MustNewTopologyMap(replicas int, assignment map[string][]shard.Shard) topology.Map {
	_ = "STUB: not implemented"
	return *new(topology.Map)
}

// NewTopologyView returns a new TopologyView with provided parameters.
// It's a utility method to make tests easier to write.
func NewTopologyView(replicas int, assignment map[string][]shard.Shard) TopologyView {
	_ = "STUB: not implemented"
	return *new(TopologyView)
}

// TopologyView represents a snaphshot view of a topology.Map.
type TopologyView struct {
	Replicas   int
	Assignment map[string][]shard.Shard
}

// Map returns the topology.Map corresponding to a TopologyView.
func (v TopologyView) Map() (topology.Map, error) {
	_ = "STUB: not implemented"
	return *new(topology.Map), nil
}

// HostShardStates is a human-readable way of describing an initial state topology
// on a host-by-host basis.
type HostShardStates map[string][]shard.Shard

// NewStateSnapshot creates a new initial topology state snapshot using HostShardStates
// as input.
func NewStateSnapshot(hostShardStates HostShardStates) *topology.StateSnapshot {
	_ = "STUB: not implemented"
	return nil
}
