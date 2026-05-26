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

package client

import (
	"context"
	. "io/ioutil"
	"net/http"

	"github.com/pkg/errors"
	"github.com/uber/aresdb/cluster/topology"
	"github.com/uber/aresdb/common"
	queryCom "github.com/uber/aresdb/query/common"
)

const (
	requestIDHeaderKey = "RequestID"
)

// ErrFailedToConnect represents error to connect to datanode
var ErrFailedToConnect = errors.New("Datanode query client failed to connect")

// NewDataNodeQueryClient creates query client to datanode
func NewDataNodeQueryClient(logger common.Logger) DataNodeQueryClient {
	_ = "STUB: not implemented"
	return *new(DataNodeQueryClient)
}

type dataNodeQueryClientImpl struct {
	client http.Client
	logger common.Logger
}

type aqlRequestBody struct {
	Queries []queryCom.AQLQuery `json:"queries"`
}

type aqlRespBody struct {
	Results []queryCom.AQLQueryResult `json:"results"`
}

func (dc *dataNodeQueryClientImpl) Query(ctx context.Context, requestID string, host topology.Host, query queryCom.AQLQuery, hll bool) (result queryCom.AQLQueryResult, err error) {
	_ = "STUB: not implemented"
	return *new(queryCom.AQLQueryResult), nil
}

func (dc *dataNodeQueryClientImpl) QueryRaw(ctx context.Context, requestID string, host topology.Host, query queryCom.AQLQuery) (bs []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (dc *dataNodeQueryClientImpl) queryRaw(ctx context.Context, requestID string, host topology.Host, query queryCom.AQLQuery, hll bool) (bs []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
