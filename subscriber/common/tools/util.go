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

package tools

var srcPkg = "github.com/uber/aresdb/subscriber"

// GetModulePath will return the source path of the given module
// under the github.com/uber/aresdb/subscriber.
// This function uses the default GOPATH/GOENV resolution to find
// the given module
func GetModulePath(module string) string {
	_ = "STUB: not implemented"
	// use the default search path for searching the package, and thus the empty
	// string for the srcDir argument
	return ""
}

// return the source directory of the package

// ToJSON converts input to JSON format
func ToJSON(data interface{}) string { _ = "STUB: not implemented"; return "" }
