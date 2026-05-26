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
	"io"
	"net/http"
	"time"

	"github.com/uber/aresdb/client"
	"github.com/uber/aresdb/controller/models"

	metaCom "github.com/uber/aresdb/metastore/common"
)

const (
	// InstanceNameHeaderKey is the key for instance name http header
	InstanceNameHeaderKey = "AresDB-InstanceName"
)

// ControllerClient defines methods to communicate with ares-controller
type ControllerClient interface {
	client.SchemaFetcher

	GetSchemaHash(namespace string) (string, error)
	GetAllSchema(namespace string) ([]metaCom.Table, error)
	GetNamespaces() ([]string, error)
	GetAssignmentHash(jobNamespace, instance string) (string, error)
	GetAssignment(jobNamespace, instance string) (*models.IngestionAssignment, error)
}

// ControllerHTTPClient implements ControllerClient over http
type ControllerHTTPClient struct {
	c         *http.Client
	address   string
	headers   http.Header
	namespace string
}

// NewControllerHTTPClient returns new ControllerHTTPClient
func NewControllerHTTPClient(address string, timeoutSec time.Duration, headers http.Header) *ControllerHTTPClient {
	_ = "STUB: not implemented"
	return nil
}

// buildRequest builds an http.Request with headers.
func (c *ControllerHTTPClient) buildRequest(method, path string, body io.Reader) (req *http.Request, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ControllerHTTPClient) getResponse(request *http.Request) (respBytes []byte, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ControllerHTTPClient) getJSONResponse(request *http.Request, output interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *ControllerHTTPClient) GetSchemaHash(namespace string) (hash string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (c *ControllerHTTPClient) GetAllSchema(namespace string) (tables []metaCom.Table, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c *ControllerHTTPClient) GetNamespaces() (namespaces []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetAssignmentHash get hash code of assignment
func (c *ControllerHTTPClient) GetAssignmentHash(jobNamespace, instance string) (hash string, err error) {
	_ = "STUB: not implemented"
	return "", nil
}

// GetAssignment gets the job assignment of the ares-subscriber
func (c *ControllerHTTPClient) GetAssignment(jobNamespace, instance string) (assignment *models.IngestionAssignment, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// SetNamespace sets the namespace which the ControllerHTTPClient connects to
func (c *ControllerHTTPClient) SetNamespace(namespace string) { _ = "STUB: not implemented"; return }

// FetchAllSchemas fetches all schemas
func (c *ControllerHTTPClient) FetchAllSchemas() (tables []*metaCom.Table, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchSchema fetch one schema for given table
func (c *ControllerHTTPClient) FetchSchema(tableName string) (table *metaCom.Table, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// FetchAllEnums fetches all enums for given table and column
func (c *ControllerHTTPClient) FetchAllEnums(tableName string, columnName string) (enumDictReponse []string, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ExtendEnumCases extends enum cases to given table column
func (c *ControllerHTTPClient) ExtendEnumCases(tableName, columnName string, enumCases []string) (enumIDs []int, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
