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

package common

import (
	"unsafe"
)

const (
	BucketSize = 8
	// log2(numHashes)
	log2NumHashes = 2
	// number of hash functions
	NumHashes = 1 << log2NumHashes

	RecordIDBytes = int(unsafe.Sizeof(RecordID{}))
)

// RecordID represents a record location with BatchID as the inflated vector id
// and offset determines the offset of record inside the vector
type RecordID struct {
	BatchID int32  `json:"batchID"`
	Index   uint32 `json:"index"`
}

// Key represents the key for the item
type Key []byte

// PrimaryKeyData holds the data for transferring to GPU for query purposes
type PrimaryKeyData struct {
	Data       unsafe.Pointer
	NumBytes   int
	Seeds      [NumHashes]uint32
	KeyBytes   int
	NumBuckets int
}

// PrimaryKey is an interface for primary key index
type PrimaryKey interface {
	// Find looks up a value given key
	Find(key Key) (RecordID, bool)
	// FindOrInsert find or insert a key value pair into
	FindOrInsert(key Key, value RecordID, eventTime uint32) (existingFound bool, recordID RecordID, err error)
	// Update updates a key with a new recordID. Return whether key exists in the primary key or not.
	Update(key Key, value RecordID) bool
	// Delete deletes a key if it exists
	Delete(key Key)
	// Update the cutoff event time.
	UpdateEventTimeCutoff(eventTimeCutoff uint32)
	// GetEventTimeCutoff returns the cutoff event time.
	GetEventTimeCutoff() uint32
	// GetDataForTransfer locks the primary key for transferring data
	// the caller should unlock by calling  UnlockAfterTransfer when done
	LockForTransfer() PrimaryKeyData
	// UnlockAfterTransfer unlocks primary key
	UnlockAfterTransfer()
	// Destruct clean up all existing resources used by primary key
	Destruct()
	// Size returns the current number of items.
	Size() uint
	// Capacity returns how many items current primary key can hold.
	Capacity() uint
	// AllocatedBytes returns the size of primary key in bytes.
	AllocatedBytes() uint
}

// MarshalPrimaryKey marshals a PrimaryKey into json. We cannot define MarshalJson for PrimaryKey
// since pointer cannot be a receiver.
func MarshalPrimaryKey(pk PrimaryKey) ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

// AppendPrimaryKeyBytes writes primary keys bytes into key buffer
func AppendPrimaryKeyBytes(key []byte, primaryKeyValues DataValueIterator) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
