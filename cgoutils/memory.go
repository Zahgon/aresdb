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

package cgoutils

// #cgo LDFLAGS: -L${SRCDIR}/../lib -lmem
// #include "string.h"
// #include "memory.h"
import "C"
import (
	"unsafe"
)

// GetFlags return flags about the memory management.
func GetFlags() C.DeviceMemoryFlags { _ = "STUB: not implemented"; return *new(C.DeviceMemoryFlags) }

func IsDeviceMemoryImplementation() bool { _ = "STUB: not implemented"; return false }

func IsPooledMemory() bool { _ = "STUB: not implemented"; return false }

func SupportHashReduction() bool { _ = "STUB: not implemented"; return false }

// HostAlloc allocates memory in C.
func HostAlloc(bytes int) unsafe.Pointer { _ = "STUB: not implemented"; return *new(unsafe.Pointer) }

// HostFree frees memory allocated in C.
func HostFree(p unsafe.Pointer) { _ = "STUB: not implemented"; return }

// HostMemCpy copies memory between two host addresses
func HostMemCpy(dst unsafe.Pointer, src unsafe.Pointer, bytes int) {
	_ = "STUB: not implemented"
	return
}

// MakeSliceFromCPtr make a slice that points to data that cptr points to.
// cptr must be a c-allocated pointer as the garbage collector will not update
// that uintptr's value if the golang object movee.
func MakeSliceFromCPtr(cptr uintptr, length int) []byte { _ = "STUB: not implemented"; return nil }

// CreateCudaStream creates a Cuda stream.
func CreateCudaStream(device int) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// WaitForCudaStream block waits until all pending operations are finished on
// the specified Cuda stream.
func WaitForCudaStream(stream unsafe.Pointer, device int) { _ = "STUB: not implemented"; return }

// DestroyCudaStream destroys the specified Cuda stream.
func DestroyCudaStream(stream unsafe.Pointer, device int) { _ = "STUB: not implemented"; return }

// DeviceAllocate allocates the specified amount of memory on the device.
func DeviceAllocate(bytes, device int) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// DeviceFree frees the specified memory from the device.
func DeviceFree(ptr unsafe.Pointer, device int) { _ = "STUB: not implemented"; return }

// AsyncMemCopyFunc is a abstraction of DeviceToDevice, DeviceToHost, HostToDevice memcopy functions
type AsyncMemCopyFunc func(dst, src unsafe.Pointer, bytes int, stream unsafe.Pointer, device int)

// AsyncCopyHostToDevice asynchronously copies the host buffer to the device
// buffer on the specified stream.
func AsyncCopyHostToDevice(
	dst, src unsafe.Pointer, bytes int, stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	return
}

// AsyncCopyDeviceToDevice asynchronously copies the src device buffer to the
// dst device buffer buffer on the specified stream.
func AsyncCopyDeviceToDevice(
	dst, src unsafe.Pointer, bytes int, stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	return
}

// AsyncCopyDeviceToHost asynchronously copies the device buffer to the host
// buffer on the specified stream.
func AsyncCopyDeviceToHost(
	dst, src unsafe.Pointer, bytes int, stream unsafe.Pointer, device int) {
	_ = "STUB: not implemented"
	return
}

// GetDeviceCount returns the number of GPU devices
func GetDeviceCount() int { _ = "STUB: not implemented"; return 0 }

// GetDeviceGlobalMemoryInMB returns the total global memory(MB) for a given device
func GetDeviceGlobalMemoryInMB(device int) int { _ = "STUB: not implemented"; return 0 }

// CudaProfilerStart starts/resumes the profiler.
func CudaProfilerStart() { _ = "STUB: not implemented"; return }

// CudaProfilerStop stops/pauses the profiler.
func CudaProfilerStop() { _ = "STUB: not implemented"; return }

// GetDeviceMemoryInfo returns information about total size and free size of device memory in bytes for a specfic
// device.
func GetDeviceMemoryInfo(device int) (int, int) { _ = "STUB: not implemented"; return 0, 0 }

// doCGoCall does the cgo call by converting CGoCallResHandle to C.int and *C.char and calls doCGoCall.
// The reason to have this wrapper is because CGo types are bound to package name, thereby even C.int are different types
// under different packages.
func doCGoCall(f func() C.CGoCallResHandle) uintptr { _ = "STUB: not implemented"; return 0 }
