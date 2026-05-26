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
	"errors"

	"github.com/m3db/m3/src/cluster/services"
	aresShard "github.com/uber/aresdb/cluster/shard"
)

var errInstanceHasNoShardsAssignment = errors.New("invalid instance with no shards assigned")

// host is the implementation of interface Host
type host struct {
	id      string
	address string
}

func (h *host) ID() string { _ = "STUB: not implemented"; return "" }

func (h *host) Address() string { _ = "STUB: not implemented"; return "" }

func (h *host) String() string { _ = "STUB: not implemented"; return "" }

// NewHost creates a new host
func NewHost(id, address string) Host { _ = "STUB: not implemented"; return *new(Host) }

// hostShardSet is the implementation of the interface HostShardSet
type hostShardSet struct {
	host     Host
	shardSet aresShard.ShardSet
}

// NewHostShardSet creates a new host shard set
func NewHostShardSet(host Host, shardSet aresShard.ShardSet) HostShardSet {
	_ = "STUB: not implemented"
	return *new(HostShardSet)
}

func (h *hostShardSet) Host() Host { _ = "STUB: not implemented"; return *new(Host) }

func (h *hostShardSet) ShardSet() aresShard.ShardSet {
	_ = "STUB: not implemented"

	// NewHostShardSetFromServiceInstance creates a new
	// host shard set derived from a service instance
	return *new(aresShard.ShardSet)
}

func NewHostShardSetFromServiceInstance(si services.ServiceInstance) (HostShardSet, error) {
	_ = "STUB: not implemented"
	return *new(HostShardSet), nil
}
