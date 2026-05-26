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

package utils

import "unsafe"

const (
	murmur3C1_32 uint32 = 0xcc9e2d51
	murmur3C2_32 uint32 = 0x1b873593
)

// Murmur3Sum32 implements Murmur3Sum32 hash algorithm
func Murmur3Sum32(key unsafe.Pointer, bytes int, seed uint32) uint32 {
	_ = "STUB: not implemented"
	return 0
}

func rotl64(x uint64, r int8) uint64 { _ = "STUB: not implemented"; return 0 }

func fmix64(k uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// Murmur3Sum128 implements murmur3sum128 hash algorithm
func Murmur3Sum128(key unsafe.Pointer, bytes int, seed uint32) (out [2]uint64) {
	_ = "STUB: not implemented"
	return nil
}

// Murmur3Sum64 use Murmur3Sum128 to generate 64bit hash
func Murmur3Sum64(key unsafe.Pointer, bytes int, seed uint32) uint64 {
	_ = "STUB: not implemented"
	return 0
}
