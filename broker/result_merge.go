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

package broker

import (
	"github.com/uber/aresdb/broker/common"
	queryCom "github.com/uber/aresdb/query/common"
)

func newResultMergeContext(aggType common.AggType) resultMergeContext {
	_ = "STUB: not implemented"
	return *new(resultMergeContext)
}

// resultMergeContext is the context for merging results
// caller should check for err after calling
type resultMergeContext struct {
	agg    common.AggType
	parent queryCom.AQLQueryResult
	path   []string
	err    error
}

// run merges results from rhs to lhs in place
func (c *resultMergeContext) run(lhs, rhs queryCom.AQLQueryResult) queryCom.AQLQueryResult {
	_ = "STUB: not implemented"
	return *new(queryCom.AQLQueryResult)
}

func (c *resultMergeContext) mergeResultsRecursive(lhs, rhs interface{}) {
	_ = "STUB: not implemented"
	return
}

// should not happen
