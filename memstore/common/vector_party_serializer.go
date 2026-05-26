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
	"github.com/uber/aresdb/diskstore"
)

// VectorPartyHeader is the magic header written into the beginning of each vector party file.
const VectorPartyHeader uint32 = 0xFADEFACE

// VectorPartyBaseSerializer is the base class contains basic data to read/write VectorParty
type vectorPartyBaseSerializer struct {
	shard, columnID, batchID int
	batchVersion             uint32
	seqNum                   uint32
	table                    string
	diskstore                diskstore.DiskStore
	hostMemoryManager        HostMemoryManager
}

// CheckVectorPartySerializable check if the archive VectorParty is serializable
func (s *vectorPartyBaseSerializer) CheckVectorPartySerializable(vp VectorParty) error {
	_ = "STUB: not implemented"
	return nil
}

// VectorPartyArchiveSerializer is the class to read/write archive VectorParty
type vectorPartyArchiveSerializer struct {
	vectorPartyBaseSerializer
}

// VectorPartyArchiveSerializer is the class to read/write snapshot VectorParty
type vectorPartySnapshotSerializer struct {
	vectorPartyBaseSerializer
	redoLogFile int64
	offset      uint32
}

// NewVectorPartyArchiveSerializer returns a new VectorPartySerializer
func NewVectorPartyArchiveSerializer(hostMemManager HostMemoryManager, diskStore diskstore.DiskStore, table string, shardID int,
	columnID int, batchID int, batchVersion uint32, seqNum uint32) VectorPartySerializer {
	_ = "STUB: not implemented"
	return *new(VectorPartySerializer)
}

// NewVectorPartySnapshotSerializer returns a new VectorPartySerializer
func NewVectorPartySnapshotSerializer(hostMemeManager HostMemoryManager, diskStore diskstore.DiskStore, table string, shardID int,
	columnID, batchID int, batchVersion uint32, seqNum uint32, redoLogFile int64, offset uint32) VectorPartySerializer {
	_ = "STUB: not implemented"
	return *new(VectorPartySerializer)
}

// ReadVectorParty reads vector party from disk and set fields in passed-in vp.
func (s *vectorPartyArchiveSerializer) ReadVectorParty(vp VectorParty) error {
	_ = "STUB: not implemented"
	return nil
}

// WriteVectorParty writes vector party to disk
func (s *vectorPartyArchiveSerializer) WriteVectorParty(vp VectorParty) error {
	_ = "STUB: not implemented"
	return nil
}

// ReportVectorPartyMemoryUsage report memory usage according to underneath VectorParty property
func (s *vectorPartyArchiveSerializer) ReportVectorPartyMemoryUsage(bytes int64) {
	_ = "STUB: not implemented"
	return
}

// WriteVectorParty writes snapshot vector party to disk
func (s *vectorPartySnapshotSerializer) WriteVectorParty(vp VectorParty) error {
	_ = "STUB: not implemented"
	return nil
}

// ReadVectorParty reads snapshot vector party from disk
func (s *vectorPartySnapshotSerializer) ReadVectorParty(vp VectorParty) error {
	_ = "STUB: not implemented"
	return nil
}

// CheckVectorPartySerializable check if the snapshot VectorParty is serializable, which is always true for now
func (s *vectorPartySnapshotSerializer) CheckVectorPartySerializable(vp VectorParty) error {
	_ = "STUB: not implemented"

	// ReportVectorPartyMemoryUsage report memory usage according to underneath VectorParty property
	return nil
}

func (s *vectorPartySnapshotSerializer) ReportVectorPartyMemoryUsage(bytes int64) {
	_ = "STUB: not implemented"
	return
}
