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

import (
	"bytes"
)

// AlignOffset returns the next offset given a specific alignment.
func AlignOffset(offset int, alignment int) int { _ = "STUB: not implemented"; return 0 }

// BufferReader provides read function for different type to read from the underline buffer.
type BufferReader struct {
	buffer []byte
}

// NewBufferReader creates a BufferReader.
func NewBufferReader(buffer []byte) BufferReader {
	_ = "STUB: not implemented"
	return *new(BufferReader)
}

// ReadUint8 reads 1 byte from buffer.
func (b BufferReader) ReadUint8(offset int) (uint8, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadInt8 reads 1 byte from buffer.
func (b BufferReader) ReadInt8(offset int) (int8, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadUint16 reads 2 bytes from buffer.
func (b BufferReader) ReadUint16(offset int) (uint16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadInt16 reads 2 bytes from buffer.
func (b BufferReader) ReadInt16(offset int) (int16, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadUint32 reads 4 bytes from buffer.
func (b BufferReader) ReadUint32(offset int) (uint32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadInt32 reads 4 bytes from buffer.
func (b BufferReader) ReadInt32(offset int) (int32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadUint64 reads 8 bytes from buffer.
func (b BufferReader) ReadUint64(offset int) (uint64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadInt64 reads 8 bytes from buffer.
func (b BufferReader) ReadInt64(offset int) (int64, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadFloat32 reads 4 bytes from buffer.
func (b BufferReader) ReadFloat32(offset int) (float32, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// BufferWriter provides functions to write different data types into the underline buffer. It
// supports both random access and sequential access.
type BufferWriter struct {
	offset    int
	bitOffset int
	buffer    []byte
}

// NewBufferWriter creates a new buffer writer over a buffer.
func NewBufferWriter(buffer []byte) BufferWriter {
	_ = "STUB: not implemented"
	return *new(BufferWriter)
}

// GetOffset returns the current offset value.
func (b BufferWriter) GetOffset() int {
	_ = "STUB: not implemented"

	// AlignBytes aligns the offset/bit offset to the next byte specified in alignment.
	// The new offset may cross the buffer boundary.
	return 0
}

func (b *BufferWriter) AlignBytes(alignment int) { _ = "STUB: not implemented"; return }

// SkipBytes moves the underline offset by specified bytes.
// The new offset may cross the buffer boundary.
func (b *BufferWriter) SkipBytes(bytes int) { _ = "STUB: not implemented"; return }

// SkipBits moves the bit offset and possibly byte offset.
// The new offset may cross the buffer boundary.
func (b *BufferWriter) SkipBits(bits int) { _ = "STUB: not implemented"; return }

// AppendBool append a boolean value to buffer and advances byte/bit offset.
func (b *BufferWriter) AppendBool(value bool) error { _ = "STUB: not implemented"; return nil }

// Need to advance to the next byte.

// AppendInt8 writes a int8 value to buffer and advances offset.
func (b *BufferWriter) AppendInt8(value int8) error { _ = "STUB: not implemented"; return nil }

// AppendUint8 writes a uint8 value to buffer and advances offset.
func (b *BufferWriter) AppendUint8(value uint8) error { _ = "STUB: not implemented"; return nil }

// AppendInt16 writes a int16 value to buffer and advances offset.
func (b *BufferWriter) AppendInt16(value int16) error { _ = "STUB: not implemented"; return nil }

// AppendUint16 writes a uint16 value to buffer and advances offset.
func (b *BufferWriter) AppendUint16(value uint16) error { _ = "STUB: not implemented"; return nil }

// AppendInt32 writes a int32 value to buffer and advances offset.
func (b *BufferWriter) AppendInt32(value int32) error { _ = "STUB: not implemented"; return nil }

// AppendInt64 writes a int64 value to buffer and advances offset.
func (b *BufferWriter) AppendInt64(value int64) error { _ = "STUB: not implemented"; return nil }

// AppendUint32 writes a uint32 value to buffer and advances offset.
func (b *BufferWriter) AppendUint32(value uint32) error { _ = "STUB: not implemented"; return nil }

// AppendUint64 writes a uint64 value to buffer and advances offset.
func (b *BufferWriter) AppendUint64(value uint64) error { _ = "STUB: not implemented"; return nil }

// AppendFloat32 writes a float32 value to buffer and advances offset.
func (b *BufferWriter) AppendFloat32(value float32) error { _ = "STUB: not implemented"; return nil }

// Append writes a byte slice to buffer and advances offset.
func (b *BufferWriter) Append(bs []byte) error { _ = "STUB: not implemented"; return nil }

// WriteInt8 writes a int8 value to buffer at given offset.
func (b *BufferWriter) WriteInt8(value int8, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteUint8 writes a uint8 value to buffer and advances offset.
func (b *BufferWriter) WriteUint8(value uint8, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteInt16 writes a int16 value to buffer and advances offset.
func (b *BufferWriter) WriteInt16(value int16, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteUint16 writes a uint16 value to buffer and advances offset.
func (b *BufferWriter) WriteUint16(value uint16, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteInt32 writes a int32 value to buffer and advances offset.
func (b *BufferWriter) WriteInt32(value int32, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteInt64 writes a int64 value to buffer and advances offset.
func (b *BufferWriter) WriteInt64(value int64, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteUint32 writes a uint32 value to buffer and advances offset.
func (b *BufferWriter) WriteUint32(value uint32, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteUint64 writes a uint64 value to buffer and advances offset.
func (b *BufferWriter) WriteUint64(value uint64, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteFloat32 writes a float32 value to buffer and advances offset.
func (b *BufferWriter) WriteFloat32(value float32, offset int) error {
	_ = "STUB: not implemented"
	return nil
}

// Write implements Write in io.Writer interface
func (b *BufferWriter) Write(bs []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// WriteAt implements WriteAt in io.WriterAt interface
func (b *BufferWriter) WriteAt(bs []byte, offset int64) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ClosableBuffer is not really closable but just implements the utils.WriteSyncCloser interface.
type ClosableBuffer struct {
	*bytes.Buffer
}

// Close just implements Close function of utils.WriteSyncCloser interface.
func (cb *ClosableBuffer) Close() error {
	_ = "STUB: not implemented"

	// ClosableReader is not really closable but just implements the utils.WriteSyncCloser interface.
	return nil
}

type ClosableReader struct {
	*bytes.Reader
}

// Close just implements Close function of utils.WriteSyncCloser interface.
func (cr *ClosableReader) Close() error {
	_ = "STUB: not implemented"

	// Sync just implements Sync function of utils.WriteSyncCloser interface.
	return nil
}

func (cr *ClosableBuffer) Sync() error { _ = "STUB: not implemented"; return nil }
