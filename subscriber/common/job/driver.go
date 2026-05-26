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
	"net/http"
	"sync"
	"time"

	controllerCli "github.com/uber/aresdb/controller/client"

	"github.com/uber-go/tally"
	"github.com/uber/aresdb/subscriber/common/rules"
	"github.com/uber/aresdb/subscriber/config"
)

// Drivers contains information about job, ares cluster and its driver
type Drivers map[string]map[string]*Driver

// Driver will initialize and start the Processor's based on the JobConfig provided
type Driver struct {
	sync.RWMutex

	Topic                string
	StartTime            time.Time
	Shutdown             bool
	JobName              string
	AresCluster          string
	TotalProcessors      int
	RunningProcessors    int
	StoppedProcessors    int
	FailedProcessors     int
	RestartingProcessors int
	ProcessorContext     map[string]*ProcessorContext

	// driver related variable
	waitGroup           *sync.WaitGroup
	shutdown            chan bool
	jobConfig           *rules.JobConfig
	serviceConfig       config.ServiceConfig
	scope               tally.Scope
	errors              chan ProcessorError
	errorThreshold      int
	statusCheckInterval int

	// processors related variables
	processorInitFunc    NewProcessor
	processorCounter     uint
	processors           []Processor
	processorMsgCount    map[int]int64
	processorMsgSizes    chan int64
	aresControllerClient controllerCli.ControllerClient
	sinkInitFunc         NewSink
	consumerInitFunc     NewConsumer
	decoderInitFunc      NewDecoder
}

// NewProcessor is the type of function each processor that implements Processor should provide for initialization
// This function implementation should always return a new instance of the processor
type NewProcessor func(id int, jobConfig *rules.JobConfig, aresControllerClient controllerCli.ControllerClient, sinkInitFunc NewSink, consumerInitFunc NewConsumer, decoderInitFunc NewDecoder,
	errors chan ProcessorError, msgSizes chan int64, serviceConfig config.ServiceConfig) (Processor, error)

// NewDrivers return Drivers
func NewDrivers(params Params, aresControllerClient controllerCli.ControllerClient) (Drivers, error) {
	_ = "STUB: not implemented"
	return *

	// iterate local job configs
	new(Drivers), nil
}

// iterate all active ares clusters

// NewDriver will return a new Driver instance to start a ingest job
func NewDriver(
	jobConfig *rules.JobConfig, serviceConfig config.ServiceConfig, aresControllerClient controllerCli.ControllerClient, processorInitFunc NewProcessor, sinkInitFunc NewSink,
	consumerInitFunc NewConsumer, decoderInitFunc NewDecoder) (*Driver, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Start will start the Driver, which starts Processor's to read from Kafka consumer, process
// and save to database
func (d *Driver) Start() { _ = "STUB: not implemented"; return }

func (d *Driver) addProcessors(count int) { _ = "STUB: not implemented"; return }

// adding too many consumers at once can make some consumers stuck with no messages

// AddProcessor will add a new processor to driver
func (d *Driver) AddProcessor() (err error) { _ = "STUB: not implemented"; return nil }

// get a new ID for processor, we start at 1

// add processor

// start processor

// Update context

// RemoveProcessor will remove a processor from driver.
func (d *Driver) RemoveProcessor(ID int) bool { _ = "STUB: not implemented"; return false }

func (d *Driver) removeProcessor(ID int) bool {
	_ = "STUB: not implemented"
	// if no processors running, nothing to remove
	return false
}

// pick last running processor

// stop the processor

// update context

// remove the processor

// remove the processor id from message count map

// restartProcessor will restart a processor from driver.
func (d *Driver) restartProcessor(ID int) { _ = "STUB: not implemented"; return }

// MarshalJSON marshal driver into json
func (d *Driver) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// WriteContext writes context
func (d *Driver) WriteContext(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

// limitRate will rate limit the amount of data read from Kafka, so to
// not saturate the network bandwidth on mesos compute hosts
func (d *Driver) limitRate() { _ = "STUB: not implemented"; return }

// only one goroutine read/write this value

// drain tokens every second to allow more

// try produce tokens

// monitorStatus will monitor the status of all Processor's and report them to graphite
func (d *Driver) monitorStatus(ticker *time.Ticker) { _ = "STUB: not implemented"; return }

// first add processors in case processor failed to be initialized at the begining

// if shutdown not intentional, but processor is dead, trouble!

// Check if messages are processed by all processors

// Update new message count for comparing with next check

// monitorErrors will monitor errors generated by each Processor and kill them if exceeds
// configured threshold number of errors per processor
func (d *Driver) monitorErrors() { _ = "STUB: not implemented"; return }

// skip errors which are before restart

// reset processor error count

// GetErrors returns errors
func (d *Driver) GetErrors() chan ProcessorError {
	_ = "STUB: not implemented"

	// Stop will shutdown driver and its processors
	return nil
}

func (d *Driver) Stop() { _ = "STUB: not implemented"; return }

// Shutdown all processors
