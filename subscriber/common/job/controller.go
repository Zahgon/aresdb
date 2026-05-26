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

package job

import (
	"sync"

	"github.com/curator-go/curator"
	"github.com/m3db/m3/src/cluster/services"
	controllerCli "github.com/uber/aresdb/controller/client"
	"github.com/uber/aresdb/subscriber/common/rules"
	"github.com/uber/aresdb/subscriber/config"
	"go.uber.org/fx"
)

// Module configures Drivers and Controller.
var Module = fx.Options(
	fx.Provide(
		NewController,
	),
	fx.Invoke(StartController),
)

// Params defines the base objects for jobConfigs.
type Params struct {
	fx.In

	LifeCycle        fx.Lifecycle
	ServiceConfig    config.ServiceConfig
	JobConfigs       rules.JobConfigs
	SinkInitFunc     NewSink
	ConsumerInitFunc NewConsumer
	DecoderInitFunc  NewDecoder
}

// Result defines the objects that the job module provides.
type Result struct {
	fx.Out

	Controller *Controller
}

const (
	// defaultRefreshInterval is 10 minutes
	defaultRefreshInterval = 10
)

// Controller is responsible for syncing up with aresDB control
type Controller struct {
	sync.RWMutex

	serviceConfig config.ServiceConfig
	// aresControllerClient is aresDB controller client
	aresControllerClient controllerCli.ControllerClient
	// Drivers are all running jobs
	Drivers Drivers
	// jobNS is current active job namespace
	jobNS string
	// aresClusterNS is current active aresDB cluster namespace
	aresClusterNS string
	// assignmentHashCode is current assignment hash code
	assignmentHashCode string
	// zkClient is zookeeper client
	zkClient curator.CuratorFramework
	// etcdServices is etcd services client
	etcdServices services.Services
	// sinkInitFunc is func of NewSink
	sinkInitFunc NewSink
	// consumerInitFunc is func of NewConsumer
	consumerInitFunc NewConsumer
	// decoderInitFunc is func of NewDecoder
	decoderInitFunc NewDecoder
}

// ZKNodeSubscriber defines the information stored in ZKNode subscriber
type ZKNodeSubscriber struct {
	// Name is subscriber instanceId
	Name string `json:"name"`
	// Host is host name of subscriber
	Host string `json:"host"`
}

// NewController creates controller
func NewController(params Params) *Controller { _ = "STUB: not implemented"; return nil }

func connectEtcdServices(params Params) (services.Services, error) {
	_ = "STUB: not implemented"
	return *new(services.Services), nil
}

// create a config service client to access to the etcd cluster services.

func registerHeartBeatService(params Params, servicesClient services.Services) error {
	_ = "STUB: not implemented"
	return nil
}

func createZKClient(params Params) curator.CuratorFramework {
	_ = "STUB: not implemented"
	return *new(curator.CuratorFramework)
}

// Using the CuratorFrameworkBuilder gives fine grained control over creation options

// RegisterOnZK registes aresDB subscriber instance in zookeeper as an ephemeral node
func (c *Controller) RegisterOnZK() error { _ = "STUB: not implemented"; return nil }

// SyncUpJobConfigs sync up jobConfigs with aresDB controller
func (c *Controller) SyncUpJobConfigs() { _ = "STUB: not implemented"; return }

// Check if the hash of the assignment is changed or not

// Get assignment from aresDB controller since hash is changed

// Add or Update jobs

// case1: existing jobConfig

// case1.1: delete the driver because aresCluster is deleted

// case1.2: restart the driver because jobConfig version is changed,

// case1.3 add a new driver because a new aresCluster is added

// case2: a new jobConfig

// case2.1: add a new driver for each aresCluster

// case2.2: delete the aresCluster from ActiveAresClusters because it is deleted from assignment

// Delete jobs

// case3: jobConfig is deleted

// Update local hash codes

func (c *Controller) updateAssignmentHash() (update bool, newHash string) {
	_ = "STUB: not implemented"
	// get the hash of the assignment
	return false, ""
}

func (c *Controller) addDriver(
	jobConfig *rules.JobConfig, aresCluster string, aresClusterDrivers map[string]*Driver, stop bool) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *Controller) deleteDriver(driver *Driver, aresCluster string, aresClusterDrivers map[string]*Driver) {
	_ = "STUB: not implemented"
	return
}

func (c *Controller) startDriver(
	jobConfig *rules.JobConfig, aresCluster string, aresClusterDrivers map[string]*Driver, stop bool) bool {
	_ = "STUB: not implemented"
	// 0. Clone jobConfig
	return false
}

// 1. Stop the job driver

// 2. create a new driver

// 3. Start the job driver

func (c *Controller) startEtcdHBService(params Params) { _ = "STUB: not implemented"; return }

// RestartEtcdHBService registers heartbeat again if etcd cluster changes are detected
func (c *Controller) RestartEtcdHBService(params Params) { _ = "STUB: not implemented"; return }

// TODO: unadevertises old heartbeat and closes etcd client should be added once M3 provides

// StartController starts periodically sync up with aresDB controller
func StartController(c *Controller) { _ = "STUB: not implemented"; return }
