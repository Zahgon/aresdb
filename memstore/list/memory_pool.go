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

package list

// #include "string.h"
// #include "stdlib.h"
import "C"

var (
	nativeChunkSize                    int64 = 1 << 25 // 32MB
	defaultStartSlabAllocatorChunkSize       = 16
	defaultSlabSize                          = 1 << 12 // 4096 bytes
	defaultSlabGrowthFactor                  = 1.5
)

type HostMemoryChangeReporter func(bytesChanged int64)

// HighLevelMemoryPool manages memory requests on pooled memory. It underlying uses a
// slab allocator to manage free memory chunks. When it no longer can satisfy memory
// allocation request from customer, it will request more memory from underlying
// NativeMemoryPool. All address returned back to client are an 2 element array of offset
// where the first offset is the memory allocated to the caller and second offset is the
// footer offset to current slab. Note that allocate a
// memory chunk larger than slabSize
// will fail. For more information related how slab allocator works, please refer to
// https://github.com/couchbase/go-slab.
type HighLevelMemoryPool interface {
	// Allocate allocates size byte memory and return back to client.
	Allocate(size int) [2]uintptr
	// Reallocate reallocates memory according to the size of the old allocated
	// memory and size of new allocation requests. If the size is the same, it
	// does nothing and just return the old addr. Otherwise it will allocate a new
	// memory, copy the old content to it if oldSize is non-zero and returned back to
	// client. If the oldSize is not zero, it will also free the old memory.
	Reallocate(oldBuf [2]uintptr, oldSize int, newSize int) [2]uintptr
	// Return the memory back to memory pool.
	Free(buf [2]uintptr)
	// Return the actual memory address given offset.
	Interpret(offset uintptr) uintptr
	// Return the underlying native memory pool allocator.
	GetNativeMemoryAllocator() NativeMemoryPool
	// Release the underlying memory.
	Destroy()
}

// slabMemoryPool implements the HighLevelMemoryPool using slab algorithm.
type slabMemoryPool struct {
	slabAllocator    *Arena
	nativeMemoryPool NativeMemoryPool
}

// NewHighLevelMemoryPool returns a default implementation of HighLevelMemoryPool.
func NewHighLevelMemoryPool(reporter HostMemoryChangeReporter) HighLevelMemoryPool {
	_ = "STUB: not implemented"
	return *new(HighLevelMemoryPool)
}

// Allocate is the implementation of Allocate in HighLevelMemoryPool interface.
func (mp slabMemoryPool) Allocate(size int) [2]uintptr { _ = "STUB: not implemented"; return nil }

// Reallocate is the implementation of Reallocate in HighLevelMemoryPool interface.
func (mp slabMemoryPool) Reallocate(oldBuf [2]uintptr, oldSize int, newSize int) [2]uintptr {
	_ = "STUB: not implemented"
	return nil
}

// Free is the implementation of Free in HighLevelMemoryPool interface.
func (mp slabMemoryPool) Free(buf [2]uintptr) { _ = "STUB: not implemented"; return }

// Interpret is the implementation of Interpret in HighLevelMemoryPool interface.
func (mp slabMemoryPool) Interpret(offset uintptr) uintptr { _ = "STUB: not implemented"; return 0 }

// Destroy is the implementation of Free in HighLevelMemoryPool interface.
func (mp slabMemoryPool) Destroy() { _ = "STUB: not implemented"; return }

// GetNativeMemoryAllocator is the implementation of GetNativeMemoryAllocator in HighLevelMemoryPool interface.
func (mp slabMemoryPool) GetNativeMemoryAllocator() NativeMemoryPool {
	_ = "STUB: not implemented"
	return *new(NativeMemoryPool)
}

// NativeMemoryPool is the interface to manage system memory to support high level memory pool
// allocation requests. All the pointer/address returned by this memory pool is relative to the
// base address fetched via GetBaseAddr.
type NativeMemoryPool interface {
	// Malloc returns a byte slice to caller, will allocate more memoryOffset if no enough space.
	// returned addresses are not aligned.
	Malloc(size int) uintptr
	// Destroy frees the memory managed by this memory pool. After destroy, any malloc call's
	// behaviour will be undefined.
	Destroy()
	// GetBaseAddr returns the base address managed by this pool.
	GetBaseAddr() uintptr
	// GetTotalBytes returns the total bytes occupied by this memory pool.
	GetTotalBytes() int64
}

// singleChunkNativeMemoryPool manages the system memory as a single chunk of continuous memory address.
// When this memory pool can no longer satisfy allocation requests, it will allocate another memory chunk
// big enough to hold the old memory content while satisfying the new requests. Then it will make a memcpy
// to move the old data into the new location. A small performance penalty will be paid during this period.
type singleChunkNativeMemoryPool struct {
	block              uintptr
	allocatedSize      int64
	totalSize          int64
	nChunks            int64
	hostMemoryReporter HostMemoryChangeReporter
}

// GetBaseAddr is the implementation of GetBaseAddr in NativeMemoryPool interface.
func (mp *singleChunkNativeMemoryPool) GetBaseAddr() uintptr {
	_ = "STUB: not implemented"

	// Malloc is the implementation of Malloc in NativeMemoryPool interface.
	return 0
}

func (mp *singleChunkNativeMemoryPool) Malloc(size int) uintptr {
	_ = "STUB: not implemented"
	return 0
}

// report memory change.

// copy old content.

// Destroy is the implementation of Destroy in NativeMemoryPool interface.
func (mp *singleChunkNativeMemoryPool) Destroy() { _ = "STUB: not implemented"; return }

// GetTotalBytes is the implementation of GetTotalBytes in NativeMemoryPool interface.
func (mp *singleChunkNativeMemoryPool) GetTotalBytes() int64 { _ = "STUB: not implemented"; return 0 }

// NewNativeMemoryPool returns a default implementation of NativeMemoryPool.
func NewNativeMemoryPool(reporter HostMemoryChangeReporter) NativeMemoryPool {
	_ = "STUB: not implemented"
	return *new(NativeMemoryPool)
}
