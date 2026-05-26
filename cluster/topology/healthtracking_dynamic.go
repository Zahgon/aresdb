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
	"sync"
	"time"
)

const (
	unhealthyRetryPeriodSeconds = 10
)

type healthiness struct {
	healthy             bool
	lastUpdateTimestamp time.Time
}

type healthTrackingDynamicTopoImpl struct {
	sync.RWMutex

	dynamicTopology  Topology
	hostsHealthiness map[Host]*healthiness
	closed           bool
}

func NewHealthTrackingDynamicTopology(opts DynamicOptions) (HealthTrackingDynamicTopoloy, error) {
	_ = "STUB: not implemented"
	return *new(HealthTrackingDynamicTopoloy), nil
}

func (ht *healthTrackingDynamicTopoImpl) Get() Map { _ = "STUB: not implemented"; return *new(Map) }

// filter known unhealthy hosts from dynamic topology

func (ht *healthTrackingDynamicTopoImpl) MarkHostHealthy(host Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (ht *healthTrackingDynamicTopoImpl) MarkHostUnhealthy(host Host) error {
	_ = "STUB: not implemented"
	return nil
}

func (ht *healthTrackingDynamicTopoImpl) changeHostHealthState(host Host, healthy bool) error {
	_ = "STUB: not implemented"
	return nil
}

// dummy implementation, don't use
// TODO: implement when needed
func (ht *healthTrackingDynamicTopoImpl) Watch() (MapWatch, error) {
	_ = "STUB: not implemented"
	return *new(MapWatch), nil
}

func (ht *healthTrackingDynamicTopoImpl) isClosed() bool { _ = "STUB: not implemented"; return false }

func (ht *healthTrackingDynamicTopoImpl) Close() { _ = "STUB: not implemented"; return }
