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

// MemAccess access memory location with starting pointer and an offset.
func MemAccess(p unsafe.Pointer, offset int) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// MemDist returns the distance between two unsafe pointer.
func MemDist(p1 unsafe.Pointer, p2 unsafe.Pointer) int64 { _ = "STUB: not implemented"; return 0 }

// MemEqual performs byte to byte comparison.
func MemEqual(a unsafe.Pointer, b unsafe.Pointer, bytes int) bool {
	_ = "STUB: not implemented"
	return false
}

// MemCopy performs memory copy of specified bytes from src to dst
func MemCopy(dst unsafe.Pointer, src unsafe.Pointer, bytes int) { _ = "STUB: not implemented"; return }

// MemSwap performs memory copy of specified bytes from src to dst
func MemSwap(dst unsafe.Pointer, src unsafe.Pointer, bytes int) { _ = "STUB: not implemented"; return }

// MemCmp performs memory comparison between two memory location start from offset
// comparing bytes byte while skip offset byte
func MemCmp(a, b unsafe.Pointer, offset, bytes int) int { _ = "STUB: not implemented"; return 0 }
