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

package memstore

import (
	"math/rand"
	"unsafe"

	"sync"

	memCom "github.com/uber/aresdb/memstore/common"
)

const (
	// size of the stash
	stashSize = 4
	// only when load factor is larger than resizeThreshold
	// we do resize without rehashing first
	resizeThreshold = 0.9
	// growth percentage for each resize
	resizeFactor float32 = 0.2

	// reserve signature 0 to indicate the slot is empty
	emptySignature = 0
	minSignature   = 1

	// offsets for bucket
	// the data layout in a bucket is the following manner
	// RecordID[8]|signature[8]|eventTime[8](optional)|key[8]
	offsetToSignature           = memCom.BucketSize * memCom.RecordIDBytes
	offsetToEventTime           = offsetToSignature + memCom.BucketSize*1
	offsetToKeyWithEventTime    = offsetToEventTime + memCom.BucketSize*4
	offsetToKeyWithoutEventTime = offsetToEventTime
)

type stashEntry struct {
	isValid   bool
	eventTime uint32
	key       memCom.Key
	value     memCom.RecordID
}

type hashResult struct {
	bucket    unsafe.Pointer
	signature uint8
}

// CuckooIndex is a implementation of Hash Index using Cuckoo Hashing algorithm
// Lazy expiration is used to invalidate expired items
// CuckooIndex is not threadsafe
type CuckooIndex struct {
	// number of bytes of a key
	keyBytes int
	// the size in bytes each bucket takes
	bucketBytes int
	// number of buckets
	numBuckets int
	// number of entries in bucket
	numBucketEntries uint
	// number of entries in stash
	numStashEntries uint
	// maxTrials when do evict and add
	maxTrials int

	// mark whether it is a fact vs dimension table hash index
	hasEventTime bool
	// bucket array
	// the array is a byte array allocated in c
	// the data layout in a bucket is the following manner
	// RecordID[8]|signature[8]|eventTime[8](optional)|key[8]
	buckets unsafe.Pointer
	// stash is a special bucket, the only difference is its size
	// use of stash is to reduce the probability of rehashing
	// Note stash memory is allocated/de-allocated together with buckets
	// stash also have 8 slots, only 4 (configurable) of them are used
	stash unsafe.Pointer

	// extra stash entry in go struct
	// act as a temporary place before resize
	staging *stagingEntry
	// seeds holds the hash function seeds
	// use different seeds to generate different hash values
	seeds [memCom.NumHashes]uint32

	// eventTimeCutoff record the smallest timestamp that was
	eventTimeCutoff uint32

	rand *rand.Rand

	// report change of unmanaged memory.
	hostMemoryManager memCom.HostMemoryManager

	// mutex protects internal buffer for GPU transfer
	transferLock sync.RWMutex
}

type stagingEntry struct {
	eventTime uint32
	key       memCom.Key
	value     memCom.RecordID
}

func getDefaultInitNumBuckets() int { _ = "STUB: not implemented"; return 0 }

// Size returns the current number of items stored in the hash table
// including expired items yet not known to the system
func (c *CuckooIndex) Size() uint { _ = "STUB: not implemented"; return 0 }

// Update updates a key with a new recordID. Return whether key exists in the primary key or not.
func (c *CuckooIndex) Update(key memCom.Key, value memCom.RecordID) bool {
	_ = "STUB: not implemented"
	return false
}

// Find looks up a record given key
func (c *CuckooIndex) Find(key memCom.Key) (memCom.RecordID, bool) {
	_ = "STUB: not implemented"
	return *new(memCom.RecordID), false
}

// Capacity returns how many items current primary key can hold.
func (c *CuckooIndex) Capacity() uint { _ = "STUB: not implemented"; return 0 }

// AllocatedBytes returns the allocated size of primary key in bytes.
func (c *CuckooIndex) AllocatedBytes() uint { _ = "STUB: not implemented"; return 0 }

// FindOrInsert find the existing key or insert a new (key, value) pair
func (c *CuckooIndex) FindOrInsert(key memCom.Key, value memCom.RecordID, eventTime uint32) (existingFound bool, recordID memCom.RecordID, err error) {
	_ = "STUB: not implemented"
	return false, *new(memCom.RecordID), nil
}

// Delete will delete a item with given key
func (c *CuckooIndex) Delete(key memCom.Key) { _ = "STUB: not implemented"; return }

// UpdateEventTimeCutoff updates eventTimeCutoff
func (c *CuckooIndex) UpdateEventTimeCutoff(cutoff uint32) { _ = "STUB: not implemented"; return }

// GetEventTimeCutoff returns the cutoff event time.
func (c *CuckooIndex) GetEventTimeCutoff() uint32 { _ = "STUB: not implemented"; return 0 }

// LockForTransfer locks primary key for transfer and returns PrimaryKeyData
func (c *CuckooIndex) LockForTransfer() memCom.PrimaryKeyData {
	_ = "STUB: not implemented"
	return *new(memCom.PrimaryKeyData)
}

// numBuckets plus stash bucket

// UnlockAfterTransfer release transfer lock
func (c *CuckooIndex) UnlockAfterTransfer() { _ = "STUB: not implemented"; return }

func (c *CuckooIndex) hash(key unsafe.Pointer, index int) hashResult {
	_ = "STUB: not implemented"
	return *new(hashResult)
}

func (c *CuckooIndex) generateRandomSeeds() { _ = "STUB: not implemented"; return }

func (c *CuckooIndex) loadFactor() float64 { _ = "STUB: not implemented"; return 0 }

// extractSignatureByte get the most significant byte from the hash value
// this is to avoid comparison of byte array as much as possible, which is expensive
func (c *CuckooIndex) extractSignatureByte(hashValue uint32) uint8 {
	_ = "STUB: not implemented"
	return 0
}

func (c *CuckooIndex) getSignature(bucket unsafe.Pointer, index int) *uint8 {
	_ = "STUB: not implemented"
	return nil
}

func (c *CuckooIndex) getRecordID(bucket unsafe.Pointer, index int) *memCom.RecordID {
	_ = "STUB: not implemented"
	return nil
}

// call should be aware there is no eventime present, this method will return incorrect
func (c *CuckooIndex) getEventTime(bucket unsafe.Pointer, index int) *uint32 {
	_ = "STUB: not implemented"
	return nil
}

func (c *CuckooIndex) getKey(bucket unsafe.Pointer, index int) unsafe.Pointer {
	_ = "STUB: not implemented"
	return *new(unsafe.Pointer)
}

func (c *CuckooIndex) isEmpty(bucket unsafe.Pointer, index int) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CuckooIndex) recordExpired(bucket unsafe.Pointer, index int) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CuckooIndex) eventTimeExpired(eventTime uint32) bool {
	_ = "STUB: not implemented"
	return false
}

// randomSwap randomly pick a bucket position and swap with the value
func (c *CuckooIndex) randomSwap(key unsafe.Pointer, recordID *memCom.RecordID, eventTime *uint32, hashResults [memCom.NumHashes]hashResult) {
	_ = "STUB: not implemented"
	return
}

// addNew only attempts to add new item into buckets, but not stash
// and assume there is no existing item
// and will not do the cuckoo process when the process fail
func (c *CuckooIndex) addNew(key unsafe.Pointer, recordID memCom.RecordID, eventTime uint32) (added bool, hashResults [memCom.NumHashes]hashResult) {
	_ = "STUB: not implemented"
	return false, nil
}

// find existing or add new item to available slot
func (c *CuckooIndex) findOrAddNew(key unsafe.Pointer, value memCom.RecordID, eventTime uint32) (existingFound bool, added bool, recordID memCom.RecordID, hashResults [memCom.NumHashes]hashResult) {
	_ = "STUB: not implemented"
	return false, false, *new(memCom.RecordID), nil
}

// look for existing record in buckets with all hash functions
// mark potential slot to insert new record

// look for existing record in stash

// randomly evict existing item with conflict hash and reinsert
func (c *CuckooIndex) cuckooAdd(key unsafe.Pointer, recordID memCom.RecordID, eventTime uint32, hashResults [memCom.NumHashes]hashResult) bool {
	_ = "STUB: not implemented"
	return false
}

// insert to stash

// save swapped out record to staging area for resizing

// insert will insert without find existing item
func (c *CuckooIndex) insert(key unsafe.Pointer, v memCom.RecordID, eventTime uint32) bool {
	_ = "STUB: not implemented"
	return false
}

func (c *CuckooIndex) insertBucket(key unsafe.Pointer, recordID memCom.RecordID, signature uint8, eventTime uint32, bucket unsafe.Pointer, index int) {
	_ = "STUB: not implemented"
	return
}

func (c *CuckooIndex) getMaxTrials() int { _ = "STUB: not implemented"; return 0 }

// resize the hast table by growFactor
func (c *CuckooIndex) resize(resizeFactor float32) (ok bool) {
	_ = "STUB: not implemented"
	return false
}

// insert existing keys to new index

// copy to current index. note transferLock is reused.

func (c *CuckooIndex) allocatedBytes() uint { _ = "STUB: not implemented"; return 0 }

func (c *CuckooIndex) allocate() { _ = "STUB: not implemented"; return }

// allocate buckets plus stash

// newCuckooIndex create a cuckoo hashing index
func newCuckooIndex(keyBytes int, hasEventTime bool, initNumBuckets int,
	hostMemoryManager memCom.HostMemoryManager) *CuckooIndex {
	_ = "STUB: not implemented"
	return nil
}

// recordIDBytes + keyBytes + signature (1 byte)

// plus eventTime (4 bytes)

// Destruct frees all allocated memory
func (c *CuckooIndex) Destruct() { _ = "STUB: not implemented"; return }

// NewPrimaryKey create a primary key data structure
// params:
//  1. keyBytes, number of bytes of key
//  2. hasEventTime determine whether primary key should record event time for expiration
//  3. initNumBuckets determines the starting number of buckets, setting to 0 to use default
func NewPrimaryKey(keyBytes int, hasEventTime bool, initNumBuckets int,
	hostMemoryManager memCom.HostMemoryManager) memCom.PrimaryKey {
	_ = "STUB: not implemented"
	return *new(memCom.PrimaryKey)
}
