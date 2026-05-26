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
package consistenthasing

import (
	"errors"
)

var (
	// ErrNodeIDExists indicates a duplicated node was added to the ring
	ErrNodeIDExists = errors.New("Node with same id already exists")
)

// Node is a node in hash ring
type Node struct {
	ID     string
	HashID uint32
}

// Nodes is a slice of nodes
type Nodes []Node

// Len is for sorting nodes
func (n Nodes) Len() int {
	_ = "STUB: not implemented"

	// Less is for sorting nodes
	return 0
}

func (n Nodes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

// Swap  is for sorting nodes
func (n Nodes) Swap(i, j int) { _ = "STUB: not implemented"; return }

// Ring is a hashring
type Ring struct {
	Nodes Nodes
	idSet map[string]bool
}

// NewNode returns a new node
func NewNode(id string) *Node { _ = "STUB: not implemented"; return nil }

// NewRing returns a new ring
func NewRing() *Ring { _ = "STUB: not implemented"; return nil }

// AddNode adds a new node
func (r *Ring) AddNode(id string) error { _ = "STUB: not implemented"; return nil }

// Get node id given key
func (r *Ring) Get(key string) (int, string) { _ = "STUB: not implemented"; return 0, "" }

func hashKey(key string) uint32 { _ = "STUB: not implemented"; return 0 }
