//	Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package etcd

import (
	"github.com/m3db/m3/src/cluster/services"
	xwatch "github.com/m3db/m3/src/x/watch"
)

// ElectionStatus represents the leader election status
type ElectionStatus int

const (
	// Unknown status
	Unknown ElectionStatus = iota
	// Leader status
	Leader
	// Follower status
	Follower
)

// LeaderElector is the interface for start leader election
type LeaderElector interface {
	// Start starts leader election
	Start() error
	// nofity channel for change in status
	C() <-chan struct{}
	// Status check current leader election status
	Status() ElectionStatus
	// Resign leader status if leader
	Resign() error
	// Close election
	Close() error
}

type leaderElector struct {
	leaderService   services.LeaderService
	statusWatchable xwatch.Watchable
	watch           xwatch.Watch
}

func (l *leaderElector) C() <-chan struct{} { _ = "STUB: not implemented"; return nil }

func (l *leaderElector) Start() error { _ = "STUB: not implemented"; return nil }

func (l *leaderElector) Status() ElectionStatus {
	_ = "STUB: not implemented"
	return *new(ElectionStatus)
}

func (l *leaderElector) Close() error { _ = "STUB: not implemented"; return nil }

func (l *leaderElector) Resign() error { _ = "STUB: not implemented"; return nil }

// NewLeaderElector creates a leader elector
func NewLeaderElector(service services.LeaderService) LeaderElector {
	_ = "STUB: not implemented"
	return *new(LeaderElector)
}
