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
	"os"
)

// FileSystem is a file system interface
type FileSystem interface {
	ReadFile(filename string) ([]byte, error)
	ReadDir(dirname string) ([]os.FileInfo, error)
	Stat(path string) (os.FileInfo, error)
	Mkdir(name string, perm os.FileMode) error
	MkdirAll(path string, perm os.FileMode) error
	Remove(path string) error
	RemoveAll(path string) error
	OpenFileForWrite(name string, flag int, perm os.FileMode) (io.WriteCloser, error)
}

// OSFileSystem implements FileSystem using os package
type OSFileSystem struct{}

// ReadFile reads whole file into byte buffer
func (OSFileSystem) ReadFile(name string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// ReadDir reads file infos under given directory
}

func (OSFileSystem) ReadDir(dirname string) ([]os.FileInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil

	// Mkdir makes directory with given name and permission
}

func (OSFileSystem) Mkdir(name string, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// MkdirAll makes directory with necessary parent directories in path
func (OSFileSystem) MkdirAll(path string, perm os.FileMode) error {
	_ = "STUB: not implemented"
	return nil
}

// OpenFileForWrite opens a file for write
func (OSFileSystem) OpenFileForWrite(name string, flag int, perm os.FileMode) (io.WriteCloser, error) {
	_ = "STUB: not implemented"
	return *new(io.WriteCloser), nil
}

// Remove removes a file
func (OSFileSystem) Remove(path string) error { _ = "STUB: not implemented"; return nil }

// RemoveAll removes a file and all its children
func (OSFileSystem) RemoveAll(path string) error { _ = "STUB: not implemented"; return nil }

// Stat tries gets file info for t
func (OSFileSystem) Stat(path string) (os.FileInfo, error) {
	_ = "STUB: not implemented"
	return *new(os.FileInfo), nil
}
