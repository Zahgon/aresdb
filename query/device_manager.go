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

package query

import (
	"sync"

	"github.com/uber/aresdb/common"
	queryCom "github.com/uber/aresdb/query/common"
)

const (
	mb2bytes                 = 1 << 20
	defaultDeviceUtilization = 1
	defaultTimeout           = 10
)

// DeviceInfo stores memory information per device
type DeviceInfo struct {
	// device id
	DeviceID int `json:"deviceID"`
	// number of queries being served by device
	QueryCount int `json:"queryCount"`
	// device capacity.
	TotalMemory int `json:"totalMemory"`
	// device available capacity.
	TotalAvailableMemory int `json:"totalAvailableMemory"`
	// total free memory
	FreeMemory int `json:"totalFreeMemory"`
	// query to memory map
	QueryMemoryUsageMap map[*queryCom.AQLQuery]int `json:"-"`
}

// DeviceManager has the following functionalities:
// 1. Keep track of number of queries being served by this device and memory usage info
// 2. Estimate the memory requirement for a given query and determine if a device has enough memory to process a query
// 3. Assign queries to chosen device according to routing strategy specified
type DeviceManager struct {
	// lock to sync ops.
	*sync.RWMutex `json:"-"`
	// device to DeviceInfo map
	DeviceInfos []*DeviceInfo `json:"deviceInfos"`
	// default DeviceChoosingTimeout for finding a device
	Timeout int `json:"timeout"`
	// Max available memory, this can be used to early determined whether a query can be satisfied or not.
	MaxAvailableMemory int `json:"maxAvailableMemory"`
	deviceAvailable    *sync.Cond
	// device choose strategy
	strategy deviceChooseStrategy
}

// NewDeviceManager is used to init a DeviceManager.
func NewDeviceManager(cfg common.QueryConfig) *DeviceManager { _ = "STUB: not implemented"; return nil }

// retrieve device counts

// Bootstrap device.

// getDeviceInfo returns the DeviceInfo struct for a given deviceID.
func getDeviceInfo(device int, deviceMemoryUtilization float32) *DeviceInfo {
	_ = "STUB: not implemented"
	return nil
}

// FindDevice finds a device to run a given query. If a device is not found, it will wait until
// the DeviceChoosingTimeout seconds elapse.
func (d *DeviceManager) FindDevice(query *queryCom.AQLQuery, requiredMem int, preferredDevice int, timeout int) int {
	_ = "STUB: not implemented"
	return 0
}

// no DeviceChoosingTimeout passed by request, using default DeviceChoosingTimeout.

// findDevice finds a device to run a given query according to certain strategy.If no such device can't
// be found, return -1. Caller needs to hold the write lock.
func (d *DeviceManager) findDevice(query *queryCom.AQLQuery, requiredMem int, preferredDevice int) int {
	_ = "STUB: not implemented"
	return 0
}

// try to choose preferredDevice if it meets requirements.

// choose candidateDevice if preferredDevice does not meet requirements

// reserve memory for this query.

// ReleaseReservedMemory adjust total free global memory for a given device after a query is complete
func (d *DeviceManager) ReleaseReservedMemory(device int, query *queryCom.AQLQuery) {
	_ = "STUB: not implemented"
	// Don't even need the lock,
	return
}

// reportMemoryUsage reports the memory usage of specified device. Caller needs to hold the lock.
func (deviceInfo *DeviceInfo) reportMemoryUsage() { _ = "STUB: not implemented"; return }

// deviceChooseStrategy defines the interface to choose an available device for
// specific query.
type deviceChooseStrategy interface {
	chooseDevice(requiredMem int) int
}

// leastAvailableMemoryStrategy is to pick up device with least query count and
// least memory that's larger than required memory of the query.
type leastQueryCountAndMemoryStrategy struct {
	deviceManager *DeviceManager
}

// chooseDevice finds a device to run a given query according to certain strategy
// If no such device, return -1.
func (s leastQueryCountAndMemoryStrategy) chooseDevice(requiredMem int) int {
	_ = "STUB: not implemented"
	return 0
}
