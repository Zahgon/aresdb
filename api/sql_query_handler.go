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

package api

import (
	"net/http"

	"github.com/uber/aresdb/utils"
)

// HandleSQL swagger:route POST /query/sql querySQL
// query in SQL
//
// Consumes:
//   - application/json
//   - application/hll
//
// Produces:
//   - application/json
//
// Responses:
//
//	default: errorResponse
//	    200: aqlResponse
//	    400: aqlResponse
func (handler *QueryHandler) HandleSQL(w *utils.ResponseWriter, r *http.Request) {
	_ = "STUB: not implemented"
	return
}
