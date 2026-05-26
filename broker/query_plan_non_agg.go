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
	"context"
	"net/http"

	"github.com/uber/aresdb/cluster/topology"
	dataCli "github.com/uber/aresdb/datanode/client"
)

// StreamingScanNode implements StreamingPlanNode
type StreamingScanNode struct {
	qc             QueryContext
	host           topology.Host
	dataNodeClient dataCli.DataNodeQueryClient
	topo           topology.HealthTrackingDynamicTopoloy
}

func (ssn *StreamingScanNode) Execute(ctx context.Context) (bs []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// context cancelled or expired, no need to query

// check ctx.Err() in case context expired during client call

// retry

func NewNonAggQueryPlan(qc *QueryContext, topo topology.HealthTrackingDynamicTopoloy, client dataCli.DataNodeQueryClient) (plan *NonAggQueryPlan, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// get rewritten query after compilation

// make a deep copy

type streamingScanNoderesult struct {
	data []byte
	err  error
}

// NonAggQueryPlan implements QueryPlan
type NonAggQueryPlan struct {
	qc         *QueryContext
	resultChan chan streamingScanNoderesult
	doneChan   chan struct{}
	headers    []string
	nodes      []*StreamingScanNode
	// number of rows flushed
	flushed int
}

func (nqp *NonAggQueryPlan) Execute(ctx context.Context, w http.ResponseWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// the first result

// only log time waited for the fastest datanode for now

// write rows

// no limit, nor need to translate enums, flush data directly

// we have to deserialize

// translate enum

// strip brackets

func (nqp *NonAggQueryPlan) getRowsWanted() int { _ = "STUB: not implemented"; return 0 }
