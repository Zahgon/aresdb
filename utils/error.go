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

package utils

// APIError represents APIError with error code
type APIError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
	Cause   error  `json:"cause"`
}

func (e APIError) Error() string { _ = "STUB: not implemented"; return "" }

// StackedError contains multiple lines of error messages as well as the stack trace.
type StackedError struct {
	Messages []string `json:"messages"`
	Stack    []string `json:"stack"`
}

func (e *StackedError) Error() string { _ = "STUB: not implemented"; return "" }

// StackError adds one more line of message to err.
// It updates err if it's already a StackedError, otherwise creates a new StackedError
// with the message from err and the stack trace of current goroutine.
func StackError(err error, message string, args ...interface{}) *StackedError {
	_ = "STUB: not implemented"
	return nil
}

// RecoverWrap recover all panics inside the passed in func
func RecoverWrap(call func() error) (err error) { _ = "STUB: not implemented"; return nil }

// find out exactly what the error was and set err
