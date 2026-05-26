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

	"github.com/uber/aresdb/broker/common"
	"github.com/uber/aresdb/cluster/topology"
	dataCli "github.com/uber/aresdb/datanode/client"
	queryCom "github.com/uber/aresdb/query/common"
)

const (
	rpcRetries = 2
)

type blockingPlanNodeImpl struct {
	children []common.BlockingPlanNode
}

func (bpn *blockingPlanNodeImpl) Children() []common.BlockingPlanNode {
	_ = "STUB: not implemented"
	return nil
}

func (bpn *blockingPlanNodeImpl) Add(nodes ...common.BlockingPlanNode) {
	_ = "STUB: not implemented"
	return
}

func NewMergeNode(agg common.AggType) common.MergeNode {
	_ = "STUB: not implemented"
	return *new(common.MergeNode)
}

type mergeNodeImpl struct {
	blockingPlanNodeImpl
	// MeasureType decides merge behaviour
	aggType common.AggType
}

func (mn *mergeNodeImpl) AggType() common.AggType {
	_ = "STUB: not implemented"
	return *new(common.AggType)
}

func (mn *mergeNodeImpl) Execute(ctx context.Context) (result queryCom.AQLQueryResult, err error) {
	_ = "STUB: not implemented"
	return *new(queryCom.AQLQueryResult), nil
}

// checks before fan out

// err means downstream retry failed

// TODO early merge before all results come back

// BlockingScanNode is a BlockingPlanNode that handles rpc calls to fetch data from datanode
type BlockingScanNode struct {
	blockingPlanNodeImpl

	qc             QueryContext
	host           topology.Host
	dataNodeClient dataCli.DataNodeQueryClient
	topo           topology.HealthTrackingDynamicTopoloy
}

func (sn *BlockingScanNode) Execute(ctx context.Context) (result queryCom.AQLQueryResult, err error) {
	_ = "STUB: not implemented"
	return *new(queryCom.AQLQueryResult), nil
}

// context cancelled or expired, cancel node execution

// check ctx.Err() in case context expired during client call

// retry

// AggQueryPlan is the plan for aggregate queries
type AggQueryPlan struct {
	aggType common.AggType
	qc      *QueryContext
	root    common.BlockingPlanNode
}

// NewAggQueryPlan creates a new agg query plan
func NewAggQueryPlan(qc *QueryContext, topo topology.HealthTrackingDynamicTopoloy, client dataCli.DataNodeQueryClient) (plan *AggQueryPlan, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// compiler already checked that only 1 measure exists, which is a expr.Call

// TODO revisit how to implement AVG. maybe add rollingAvg to datanode so only 1 call per shard needed

func (ap *AggQueryPlan) postProcess(results queryCom.AQLQueryResult, execErr error, w http.ResponseWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (ap *AggQueryPlan) translateEnum(results queryCom.AQLQueryResult) (rewritten interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// helper function that traverse AQLQueryResult, translate enum rank to value
func traverseRecursive(dimIndex int, curr interface{}, dimReverseDict map[int][]string) (rewritten interface{}, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// visit curr and translate

func getResultSizeRecursive(res interface{}) int { _ = "STUB: not implemented"; return 0 }

// convert HLL to binary format
func (ap *AggQueryPlan) postProcessHLLBinary(res queryCom.AQLQueryResult, execErr error) (data []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// build dataTypes and enumDicts

// build HLL binary

// write hll binary to hll result, and return

func (ap *AggQueryPlan) Execute(ctx context.Context, w http.ResponseWriter) (err error) {
	_ = "STUB: not implemented"
	return nil
}

// splitAvgQuery to sum and count queries
func splitAvgQuery(qc QueryContext) (sumqc QueryContext, countqc QueryContext) {
	_ = "STUB: not implemented"
	return *new(QueryContext), *new(QueryContext)
}

func buildSubPlan(agg common.AggType, qc QueryContext, assignments map[topology.Host][]uint32, topo topology.HealthTrackingDynamicTopoloy, client dataCli.DataNodeQueryClient) common.MergeNode {
	_ = "STUB: not implemented"
	return *new(common.MergeNode)
}

// make deep copy
