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

package query

import (
	memCom "github.com/uber/aresdb/memstore/common"
)

// SerializeHLL allocates buffer based on the metadata and then serializes hll data into the buffer.
func (qc *AQLQueryContext) SerializeHLL(dataTypes []memCom.DataType,
	enumDicts map[int][]string, timeDimensions []int) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Copy dim values vector from device.

// Fix time dimension by substracting the timezone.

// length is equal to length of timeDimensions

// We don't need to do anything for null.

// Don't need to check type of time dimension, they should be guaranteed by AQL Compiler.
