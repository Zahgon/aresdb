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

package shard

import (
	"github.com/m3db/m3/src/cluster/shard"
)

// shardSet is the implementation of the interface ShardSet
type shardSet struct {
	shards   []shard.Shard
	ids      []uint32
	shardMap map[uint32]shard.Shard
}

// NewShardSet creates a new sharding scheme with a set of shards
func NewShardSet(shards []shard.Shard) ShardSet { _ = "STUB: not implemented"; return *new(ShardSet) }

func (s *shardSet) All() []shard.Shard { _ = "STUB: not implemented"; return nil }

func (s *shardSet) AllIDs() []uint32 {
	_ = "STUB: not implemented"

	// NewShards returns a new slice of shards with a specified state
	return nil
}

func NewShards(ids []uint32, state shard.State) []shard.Shard {
	_ = "STUB: not implemented"
	return nil
}

// IDs returns a new slice of shard IDs for a set of shards
func IDs(shards []shard.Shard) []uint32 { _ = "STUB: not implemented"; return nil }

// intRange returns a slice of all values between [from, to].
func intRange(from, to uint32) []uint32 { _ = "STUB: not implemented"; return nil }

// ShardsRange returns a slice of shards for all ids between [from, to],
// with shard state `s`.
func ShardsRange(from, to uint32, s shard.State) []shard.Shard {
	_ = "STUB: not implemented"
	return nil
}
