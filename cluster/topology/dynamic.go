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
	"sync"

	"github.com/m3db/m3/src/cluster/services"
	xwatch "github.com/m3db/m3/src/x/watch"
	"github.com/uber/aresdb/common"
)

var (
	errInvalidService            = errors.New("service topology is invalid")
	errUnexpectedShard           = errors.New("shard is unexpected")
	errMissingShard              = errors.New("shard is missing")
	errNotEnoughReplicasForShard = errors.New("replicas of shard is less than expected")
	errInvalidTopology           = errors.New("could not parse latest value from config service")
)

type dynamicInitializer struct {
	sync.Mutex

	opts DynamicOptions
	topo Topology
}

// NewDynamicInitializer returns a dynamic topology initializer
func NewDynamicInitializer(opts DynamicOptions) Initializer {
	_ = "STUB: not implemented"
	return *new(Initializer)
}

func (i *dynamicInitializer) Init() (Topology, error) {
	_ = "STUB: not implemented"
	return *new(Topology), nil
}

func (i *dynamicInitializer) TopologyIsSet() (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// Valid, just means topology is not set

type dynamicTopology struct {
	sync.RWMutex

	opts      DynamicOptions
	services  services.Services
	watch     services.Watch
	watchable xwatch.Watchable
	closed    bool
	logger    common.Logger
}

func newDynamicTopology(opts DynamicOptions) (DynamicTopology, error) {
	_ = "STUB: not implemented"
	return *new(DynamicTopology), nil
}

func (t *dynamicTopology) isClosed() bool { _ = "STUB: not implemented"; return false }

func (t *dynamicTopology) run() { _ = "STUB: not implemented"; return }

func (t *dynamicTopology) Get() Map { _ = "STUB: not implemented"; return *new(Map) }

func (t *dynamicTopology) Watch() (MapWatch, error) {
	_ = "STUB: not implemented"
	return *new(MapWatch), nil
}

func (t *dynamicTopology) Close() { _ = "STUB: not implemented"; return }

func (t *dynamicTopology) MarkShardsAvailable(
	instanceID string,
	shardIDs ...uint32,
) error {
	_ = "STUB: not implemented"
	return nil
}

func getMapFromUpdate(service services.Service, unhealthyIncluded bool) (Map, error) {
	_ = "STUB: not implemented"
	return *new(Map), nil
}

func getStaticOptions(service services.Service, unhealthyIncluded bool) (StaticOptions, error) {
	_ = "STUB: not implemented"
	return *new(StaticOptions), nil
}

func validateInstances(instances []services.ServiceInstance, unhealthyIncluded bool, replicas, numShards int) ([]uint32, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
