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

import (
	"net/http"
	"time"

	"github.com/uber-go/tally"
	"github.com/uber/aresdb/common"
)

const (
	HTTPContentTypeHeaderKey     = "Content-Type"
	HTTPAcceptTypeHeaderKey      = "Accept"
	HTTPAcceptEncodingHeaderKey  = "Accept-Encoding"
	HTTPContentEncodingHeaderKey = "Content-Encoding"

	HTTPContentTypeApplicationJson = "application/json"
	HTTPContentTypeApplicationGRPC = "application/grpc"
	// HTTPContentTypeUpsertBatch defines the upsert data content type.
	HTTPContentTypeUpsertBatch = "application/upsert-data"
	// HTTPContentTypeHyperLogLog defines the hyperloglog query result content type.
	HTTPContentTypeHyperLogLog = "application/hll"
	HTTPContentEncodingGzip    = "gzip"

	// CompressionThreshold is the min number of bytes beyond which we will compress json payload
	CompressionThreshold = 1 << 10
)

var epoch = time.Unix(0, 0).Format(time.RFC1123)

var noCacheHeaders = map[string]string{
	"Expires":         epoch,
	"Cache-Control":   "no-cache, private, max-age=0",
	"Pragma":          "no-cache",
	"X-Accel-Expires": "0",
}

var etagHeaders = []string{
	"ETag",
	"If-Modified-Since",
	"If-Match",
	"If-None-Match",
	"If-Range",
	"If-Unmodified-Since",
}

// NoCache sets no cache headers and removes any ETag headers that may have been set.
func NoCache(h http.Handler) http.Handler { _ = "STUB: not implemented"; return *new(http.Handler) }

// GetOrigin returns the caller of the request.
func GetOrigin(r *http.Request) string { _ = "STUB: not implemented"; return "" }

// LimitServe will start a http server on the port with the handler and at most maxConnection concurrent connections.
func LimitServe(port int, handler http.Handler, httpCfg common.HTTPConfig) {
	_ = "STUB: not implemented"
	return
}

// LimitServeAsync will start a http server on the port with the handler and at most maxConnection concurrent connections.
func LimitServeAsync(port int, handler http.Handler, httpCfg common.HTTPConfig) (chan error, *http.Server) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HandlerFunc defines http handler function
type HandlerFunc func(rw *ResponseWriter, r *http.Request)

// HTTPHandlerWrapper wraps http handler function
type HTTPHandlerWrapper func(handler HandlerFunc) HandlerFunc

// ApplyHTTPWrappers apply wrappers according to the order
func ApplyHTTPWrappers(handler HandlerFunc, wrappers ...HTTPHandlerWrapper) http.HandlerFunc {
	_ = "STUB: not implemented"
	return *new(http.HandlerFunc)
}

// MetricsLoggingMiddleWareProvider provides middleware for metrics and logger for http requests
type MetricsLoggingMiddleWareProvider struct {
	scope  tally.Scope
	logger common.Logger
}

// NewMetricsLoggingMiddleWareProvider creates metrics and logging middleware provider
func NewMetricsLoggingMiddleWareProvider(scope tally.Scope, logger common.Logger) MetricsLoggingMiddleWareProvider {
	_ = "STUB: not implemented"
	return *new(MetricsLoggingMiddleWareProvider)
}

// WithMetrics plug in metrics middleware
func (p *MetricsLoggingMiddleWareProvider) WithMetrics(next HandlerFunc) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

// WithLogging plug in logging middleware
func (p *MetricsLoggingMiddleWareProvider) WithLogging(next HandlerFunc) HandlerFunc {
	_ = "STUB: not implemented"
	return *new(HandlerFunc)
}

func setCommonHeaders(w http.ResponseWriter) { _ = "STUB: not implemented"; return }

// ErrorResponse represents error response.
// swagger:response errorResponse
type ErrorResponse struct {
	//in: body
	Body APIError
}

// ResponseWriter decorates http.ResponseWriter
type ResponseWriter struct {
	http.ResponseWriter
	statusCode int
	req        interface{}
	err        error
}

// NewResponseWriter returns response writer with status code 200
func NewResponseWriter(rw http.ResponseWriter) *ResponseWriter {
	_ = "STUB: not implemented"
	return nil
}

// SetRequest set unmarshalled request body to response writer for logging purpose
func (s *ResponseWriter) SetRequest(req interface{}) {
	_ = "STUB: not implemented"

	// WriteHeader implements http.ResponseWriter WriteHeader for write status code
	return
}

func (s *ResponseWriter) WriteHeader(code int) { _ = "STUB: not implemented"; return }

// WriteBytes implements http.ResponseWriter Write for write bytes
func (s *ResponseWriter) WriteBytes(bts []byte) { _ = "STUB: not implemented"; return }

// WriteBytesWithCode writes bytes with code
func (s *ResponseWriter) WriteBytesWithCode(code int, bts []byte) {
	_ = "STUB: not implemented"
	return
}

// WriteJSONBytes write json bytes with default status ok
func (s *ResponseWriter) WriteJSONBytes(jsonBytes []byte, marshalErr error) {
	_ = "STUB: not implemented"
	return
}

// WriteJSONBytesWithCode write json bytes and marshal error to response
func (s *ResponseWriter) WriteJSONBytesWithCode(code int, jsonBytes []byte, marshalErr error) {
	_ = "STUB: not implemented"
	return
}

// ignore this error since this should not happen

// try best effort write with gzip compression

// default to normal json response

// WriteObject write json object to response
func (s *ResponseWriter) WriteObject(obj interface{}) { _ = "STUB: not implemented"; return }

// WriteObjectWithCode serialize object and write code
func (s *ResponseWriter) WriteObjectWithCode(code int, obj interface{}) {
	_ = "STUB: not implemented"
	return
}

// WriteErrorWithCode writes error with specific code
func (s *ResponseWriter) WriteErrorWithCode(code int, err error) { _ = "STUB: not implemented"; return }

// WriteError write error to response
func (s *ResponseWriter) WriteError(err error) { _ = "STUB: not implemented"; return }
