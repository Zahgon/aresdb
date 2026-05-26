package common

import (
	"unsafe"
)

// SetValue sets the value as dataType at ith position from baseAddr.
func SetValue(baseAddr uintptr, i int, val unsafe.Pointer, dataType DataType) {
	_ = "STUB: not implemented"
	return
}

// SetBool sets the value as bool at ith position from baseAddr.
func SetBool(baseAddr uintptr, i int, val bool) { _ = "STUB: not implemented"; return }

// GetValue reads value as dataType at ith position from baseAddr
func GetValue(baseAddr uintptr, i int, dataType DataType) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

// GetBool reads the value as bool at ith position from baseAddr.
func GetBool(baseAddr uintptr, i int) bool { _ = "STUB: not implemented"; return false }
