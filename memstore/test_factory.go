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

// Package memstore has to put test factory here since otherwise we will have a
// memstore -> utils -> memstore import cycle.
package memstore

import (
	"sync"

	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/memstore/tests"
	"github.com/uber/aresdb/memstore/vectors"
	"github.com/uber/aresdb/utils"
)

const (
	vpValueTestDelimiter = ","
)

var (
	testFactory = TestFactoryT{
		TestFactoryBase: tests.TestFactoryBase{
			RootPath:             "../testing/data",
			FileSystem:           utils.OSFileSystem{},
			ToArchiveVectorParty: toArchiveVectorParty,
			ToLiveVectorParty:    toLiveVectorParty,
			ToVectorParty:        toVectorParty,
		},
	}
)

// TestFactoryT creates memstore test objects from text file
type TestFactoryT struct {
	tests.TestFactoryBase
}

// NewMockMemStore returns a new memstore with mocked diskstore and metastore.
func (t TestFactoryT) NewMockMemStore() *memStoreImpl { _ = "STUB: not implemented"; return nil }

func toArchiveVectorParty(vp memCom.VectorParty, locker sync.Locker) memCom.ArchiveVectorParty {
	_ = "STUB: not implemented"
	return *new(memCom.ArchiveVectorParty)
}

func toLiveVectorParty(vp memCom.VectorParty) memCom.LiveVectorParty {
	_ = "STUB: not implemented"
	return *new(memCom.LiveVectorParty)
}

func toVectorParty(rvp *tests.RawVectorParty, forLiveVP bool) (memCom.VectorParty, error) {
	_ = "STUB: not implemented"
	return *new(memCom.VectorParty), nil
}

func setDataValue(v *vectors.Vector, idx int, val memCom.DataValue) {
	_ = "STUB: not implemented"
	return
}

func GetFactory() TestFactoryT { _ = "STUB: not implemented"; return *new(TestFactoryT) }
