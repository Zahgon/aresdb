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

import (
	"sync"

	memCom "github.com/uber/aresdb/memstore/common"
	"github.com/uber/aresdb/memstore/tests"
	"github.com/uber/aresdb/utils"
)

var (
	testFactory = TestFactoryT{
		TestFactoryBase: tests.TestFactoryBase{
			RootPath:             "../../testing/data",
			FileSystem:           utils.OSFileSystem{},
			ToArchiveVectorParty: ToArrayArchiveVectorParty,
			ToLiveVectorParty:    ToArrayLiveVectorParty,
			ToVectorParty:        ToArrayVectorParty,
		},
	}
)

// TestFactoryT creates test objects from text file
type TestFactoryT struct {
	tests.TestFactoryBase
}

func GetFactory() TestFactoryT { _ = "STUB: not implemented"; return *new(TestFactoryT) }

func ToArrayArchiveVectorParty(vp memCom.VectorParty, locker sync.Locker) memCom.ArchiveVectorParty {
	_ = "STUB: not implemented"
	return *new(memCom.ArchiveVectorParty)
}

func ToArrayLiveVectorParty(vp memCom.VectorParty) memCom.LiveVectorParty {
	_ = "STUB: not implemented"
	return *new(memCom.LiveVectorParty)
}

func ToArrayVectorParty(rvp *tests.RawVectorParty, forLiveVP bool) (vp memCom.VectorParty, err error) {
	_ = "STUB: not implemented"
	return *new(memCom.VectorParty), nil
}

// array live party

// array archive party
