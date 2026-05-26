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
)

type staticInitializer struct {
	opts StaticOptions
}

// NewStaticInitializer creates a static topology initializer.
func NewStaticInitializer(opts StaticOptions) Initializer {
	_ = "STUB: not implemented"
	return *new(Initializer)
}

func (i staticInitializer) Init() (Topology, error) {
	_ = "STUB: not implemented"
	return *new(Topology), nil
}

func (i staticInitializer) TopologyIsSet() (bool, error) {
	_ = "STUB: not implemented"
	// Always has the specified static topology ready.
	return false, nil
}

type staticTopology struct {
	w xwatch.Watchable
}

// NewStaticTopology creates a static topology.
func NewStaticTopology(opts StaticOptions) Topology {
	_ = "STUB: not implemented"
	return *new(Topology)
}

func (t *staticTopology) Get() Map { _ = "STUB: not implemented"; return *new(Map) }

func (t *staticTopology) Watch() (MapWatch, error) {
	_ = "STUB: not implemented"
	// Topology is static, the returned watch will not receive any updates.
	return *new(MapWatch), nil
}

func (t *staticTopology) Close() { _ = "STUB: not implemented"; return }
