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
	"net/http"
)

// WithRequest defines function callback with parsed request
type WithRequest func(request interface{})

// ReadRequest reads request.
// obj passed into this method has to be a pointer to a struct of request object
// Each request object will have path params tagged as `path:""` if needed
// and post body tagged as `body:""` if needed
// path tag must have parameter name, which will be used to read path param
// body tag field has to be a struct.
// eg.
//
//	type AddEnumCaseRequest struct {
//		TableName string `path:"table"`
//		ColumnName string `path:"column"`
//		Body struct {
//			EnumCase string `json:"enumCase"`
//		} `body:""`
//	}
func ReadRequest(r *http.Request, obj interface{}, withRequests ...WithRequest) error {
	_ = "STUB: not implemented"
	return nil
}

// If it's anonymous field, we apply ReadRequest to this struct directly.

// Only string and int is supported in request path fields.
