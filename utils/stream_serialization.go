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
	"io"
)

// StreamDataReader reads primitive Golang data types from an underlying reader.
// It always advance the read iterator without rewinding it.
type StreamDataReader struct {
	reader    io.Reader
	bytesRead uint32
}

// NewStreamDataReader returns a StreamDataReader given an underlying reader.
func NewStreamDataReader(reader io.Reader) StreamDataReader {
	_ = "STUB: not implemented"
	return *new(StreamDataReader)
}

// GetBytesRead returns number of bytes read so far.
func (r StreamDataReader) GetBytesRead() uint32 {
	_ = "STUB: not implemented"

	// Read reads bytes from underlying reader and fill bs with len(bs) bytes data. Raise
	// error if there is no enough bytes left.
	return 0
}

func (r *StreamDataReader) Read(bs []byte) error { _ = "STUB: not implemented"; return nil }

// Implementations of Read are discouraged from returning a
// zero byte count with a nil error, except when len(p) == 0.
// Callers should treat a return of 0 and nil as indicating that
// nothing happened; in particular it does not indicate EOF.

// We return EOF directly without wrapping so that callers can take special actions against EOF.

// We have to check readLen here since a non-zero number of bytes at the end of the input stream
// may return either err == EOF or err == nil.

// ReadUint8 reads one uint8 from the reader and advance.
func (r *StreamDataReader) ReadUint8() (uint8, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadInt8 reads one int8 from the reader and advance.
func (r *StreamDataReader) ReadInt8() (int8, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadUint16 reads one uint16 from the reader and advance.
func (r *StreamDataReader) ReadUint16() (uint16, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadInt16 reads one int16 from the reader and advance.
func (r *StreamDataReader) ReadInt16() (int16, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadUint32 reads one uint32 from the reader and advance.
func (r *StreamDataReader) ReadUint32() (uint32, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadInt32 reads one int32 from the reader and advance.
func (r *StreamDataReader) ReadInt32() (int32, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadUint64 reads one uint64 from the reader and advance.
func (r *StreamDataReader) ReadUint64() (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// ReadFloat32 reads one float32 from the reader and advance.
func (r *StreamDataReader) ReadFloat32() (float32, error) { _ = "STUB: not implemented"; return 0, nil }

// SkipBytes read some empty bytes from the reader and discard it
func (r *StreamDataReader) SkipBytes(nbytes int) error { _ = "STUB: not implemented"; return nil }

// ReadPadding reads and ignore bytes until alignment is met.
// nbytes is the size of bytes already read.
func (r *StreamDataReader) ReadPadding(nbytes int, alignment int) error {
	_ = "STUB: not implemented"
	return nil
}

// StreamDataWriter writes primitive Golang data types to an underlying writer.
// It always advance the write iterator without rewinding it.
type StreamDataWriter struct {
	writer       io.Writer
	bytesWritten uint32
}

// NewStreamDataWriter returns a StreamDataWriter given an underlying writer.
func NewStreamDataWriter(writer io.Writer) StreamDataWriter {
	_ = "STUB: not implemented"
	return *new(StreamDataWriter)
}

// GetBytesWritten returns number of bytes written by current writer so far.
func (w StreamDataWriter) GetBytesWritten() uint32 { _ = "STUB: not implemented"; return 0 }

func (w *StreamDataWriter) Write(bs []byte) error {
	_ = "STUB: not implemented"
	// We can ignore n here since Write must return a non-nil error if it returns n < len(p).
	return nil
}

// WriteUint8 writes one uint8 to the writer and advance.
func (w *StreamDataWriter) WriteUint8(v uint8) error { _ = "STUB: not implemented"; return nil }

// WriteInt8 writes one int8 to the writer and advance.
func (w *StreamDataWriter) WriteInt8(v int8) error { _ = "STUB: not implemented"; return nil }

// WriteUint16 writes one uint16 to the writer and advance.
func (w *StreamDataWriter) WriteUint16(v uint16) error { _ = "STUB: not implemented"; return nil }

// WriteInt16 writes one int16 to the writer and advance.
func (w *StreamDataWriter) WriteInt16(v int16) error { _ = "STUB: not implemented"; return nil }

// WriteUint32 writes one uint32 to the writer and advance.
func (w *StreamDataWriter) WriteUint32(v uint32) error { _ = "STUB: not implemented"; return nil }

// WriteInt32 writes one int32 to the writer and advance.
func (w *StreamDataWriter) WriteInt32(v int32) error { _ = "STUB: not implemented"; return nil }

// WriteUint64 writes one uint64 to the writer and advance.
func (w *StreamDataWriter) WriteUint64(v uint64) error { _ = "STUB: not implemented"; return nil }

// WriteFloat32 writes one float32 to the writer and advance.
func (w *StreamDataWriter) WriteFloat32(v float32) error { _ = "STUB: not implemented"; return nil }

// SkipBytes write some empty bytes to the writer and advance.
func (w *StreamDataWriter) SkipBytes(nbytes int) error { _ = "STUB: not implemented"; return nil }

// WritePadding write some empty bytes until alignment is met.
// nbytes is the size of bytes already written.
func (w *StreamDataWriter) WritePadding(nbytes int, alignment int) error {
	_ = "STUB: not implemented"
	return nil
}
