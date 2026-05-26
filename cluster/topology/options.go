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
	"time"

	"github.com/m3db/m3/src/cluster/client"
	"github.com/m3db/m3/src/cluster/services"
	"github.com/uber/aresdb/cluster/shard"
	"github.com/uber/aresdb/utils"
)

const (
	defaultServiceName = "aresDB"
	defaultInitTimeout = 0 // Wait indefinitely by default for topology
	defaultReplicas    = 3
)

var (
	errNoConfigServiceClient = errors.New("no config service client")
	errInvalidReplicas       = errors.New("replicas must be equal to or greater than 1")
)

// staticOptions is the implementation of the interface StaticOptions
type staticOptions struct {
	shardSet      shard.ShardSet
	hostShardSets []HostShardSet
	replicas      int
}

// NewStaticOptions creates a new set of static topology options
func NewStaticOptions() StaticOptions { _ = "STUB: not implemented"; return *new(StaticOptions) }

func (o *staticOptions) SetShardSet(value shard.ShardSet) StaticOptions {
	_ = "STUB: not implemented"
	return *new(StaticOptions)
}

func (o *staticOptions) ShardSet() shard.ShardSet {
	_ = "STUB: not implemented"
	return *new(shard.ShardSet)
}

func (o *staticOptions) SetHostShardSets(value []HostShardSet) StaticOptions {
	_ = "STUB: not implemented"
	return *new(StaticOptions)
}

func (o *staticOptions) HostShardSets() []HostShardSet { _ = "STUB: not implemented"; return nil }

func (o *staticOptions) SetReplicas(value int) StaticOptions {
	_ = "STUB: not implemented"
	return *new(StaticOptions)
}

func (o *staticOptions) Replicas() int {
	_ = "STUB: not implemented"

	// dynamicOptions is the implementation of the interface DynamicOptions
	return 0
}

type dynamicOptions struct {
	configServiceClient     client.Client
	serviceID               services.ServiceID
	servicesOverrideOptions services.OverrideOptions
	queryOptions            services.QueryOptions
	instrumentOptions       utils.Options
	initTimeout             time.Duration
}

// NewDynamicOptions creates a new set of dynamic topology options
func NewDynamicOptions() DynamicOptions { _ = "STUB: not implemented"; return *new(DynamicOptions) }

func (o *staticOptions) Validate() error { _ = "STUB: not implemented"; return nil }

// Make a mapping of each shard to a set of hosts and check each
// shard has at least the required replicas mapped to
// NB(r): We allow greater than the required replicas in case
// node is streaming in and needs to take writes

func (o *dynamicOptions) SetConfigServiceClient(c client.Client) DynamicOptions {
	_ = "STUB: not implemented"
	return *new(DynamicOptions)
}

func (o *dynamicOptions) ConfigServiceClient() client.Client {
	_ = "STUB: not implemented"
	return *new(client.Client)
}

func (o *dynamicOptions) SetServiceID(s services.ServiceID) DynamicOptions {
	_ = "STUB: not implemented"
	return *new(DynamicOptions)
}

func (o *dynamicOptions) ServiceID() services.ServiceID {
	_ = "STUB: not implemented"
	return *new(services.ServiceID)
}

func (o *dynamicOptions) SetServicesOverrideOptions(opts services.OverrideOptions) DynamicOptions {
	_ = "STUB: not implemented"
	return *new(DynamicOptions)
}

func (o *dynamicOptions) ServicesOverrideOptions() services.OverrideOptions {
	_ = "STUB: not implemented"
	return *new(services.OverrideOptions)
}

func (o *dynamicOptions) SetQueryOptions(qo services.QueryOptions) DynamicOptions {
	_ = "STUB: not implemented"
	return *new(DynamicOptions)
}

func (o *dynamicOptions) QueryOptions() services.QueryOptions {
	_ = "STUB: not implemented"
	return *new(services.QueryOptions)
}

func (o *dynamicOptions) SetInstrumentOptions(io utils.Options) DynamicOptions {
	_ = "STUB: not implemented"
	return *new(DynamicOptions)
}

func (o *dynamicOptions) InstrumentOptions() utils.Options {
	_ = "STUB: not implemented"
	return *new(utils.Options)
}

func (o *dynamicOptions) Validate() error { _ = "STUB: not implemented"; return nil }
