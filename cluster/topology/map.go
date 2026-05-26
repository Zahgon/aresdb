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

package topology

import (
	xwatch "github.com/m3db/m3/src/x/watch"
	aresShard "github.com/uber/aresdb/cluster/shard"
)

// staticMap is the implementation of the interface Map
type staticMap struct {
	shardSet          aresShard.ShardSet
	hostShardSets     []HostShardSet
	hostShardSetsByID map[string]HostShardSet
	hostsByShard      [][]Host
	orderedHosts      []Host
	replicas          int
}

// NewStaticMap creates Map
func NewStaticMap(opts StaticOptions) Map { _ = "STUB: not implemented"; return *new(Map) }

func (sm *staticMap) Hosts() []Host { _ = "STUB: not implemented"; return nil }

func (sm *staticMap) HostShardSets() []HostShardSet { _ = "STUB: not implemented"; return nil }

func (sm *staticMap) LookupHostShardSet(hostID string) (HostShardSet, bool) {
	_ = "STUB: not implemented"
	return *new(HostShardSet), false
}

func (sm *staticMap) HostsLen() int { _ = "STUB: not implemented"; return 0 }

func (sm *staticMap) ShardSet() aresShard.ShardSet {
	_ = "STUB: not implemented"
	return *new(aresShard.ShardSet)
}

func (sm *staticMap) RouteShard(shard uint32) ([]Host, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (t *staticMap) Replicas() int {
	_ = "STUB: not implemented"

	// mapWatch is the implementation of the interface MapWatch
	return 0
}

type mapWatch struct {
	xwatch.Watch
}

// NewMapWatch creates MapWatch
func NewMapWatch(w xwatch.Watch) MapWatch { _ = "STUB: not implemented"; return *new(MapWatch) }

func (w *mapWatch) C() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (w *mapWatch) Get() Map { _ = "STUB: not implemented"; return *new(Map) }

func (w *mapWatch) Close() { _ = "STUB: not implemented"; return }
