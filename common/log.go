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
	"go.uber.org/zap"
)

// Logger is a general logger interface
type Logger interface {
	// Log at debug level
	Debug(args ...interface{})

	// Log at debug level with fmt.Printf-like formatting
	Debugf(format string, args ...interface{})

	// Log at info level
	Info(args ...interface{})

	// Log at info level with fmt.Printf-like formatting
	Infof(format string, args ...interface{})

	// Log at warning level
	Warn(args ...interface{})

	// Log at warning level with fmt.Printf-like formatting
	Warnf(format string, args ...interface{})

	// Log at error level
	Error(args ...interface{})

	// Log at error level with fmt.Printf-like formatting
	Errorf(format string, args ...interface{})

	// Log at fatal level, then terminate process (irrecoverable)
	Fatal(args ...interface{})

	// Log at fatal level with fmt.Printf-like formatting, then terminate process (irrecoverable)
	Fatalf(format string, args ...interface{})

	// Log at panic level, then panic (recoverable)
	Panic(args ...interface{})

	// Log at panic level with fmt.Printf-like formatting, then panic (recoverable)
	Panicf(format string, args ...interface{})

	// Return a logger with the specified key-value pair set, to be logged in a subsequent normal logging call
	With(args ...interface{}) Logger
}

// LoggerFactory defines the log factory ares needs.
type LoggerFactory interface {
	// GetDefaultLogger returns the default logger.
	GetDefaultLogger() Logger
	// GetLogger returns logger given the logger name.
	GetLogger(name string) Logger
}

// ZapLoggerFactory is the stdlog implementation of LoggerFactory
type ZapLoggerFactory struct {
	logger *ZapLogger
}

// NewZapLoggerFactory creates ZapLoggerFactory
func NewZapLoggerFactory(logger *zap.SugaredLogger) LoggerFactory {
	_ = "STUB: not implemented"
	return *new(LoggerFactory)
}

// NewLoggerFactory creates a default zap LoggerFactory implementation.
func NewLoggerFactory() LoggerFactory { _ = "STUB: not implemented"; return *new(LoggerFactory) }

// GetDefaultLogger returns the default zap logger.
func (r *ZapLoggerFactory) GetDefaultLogger() Logger {
	_ = "STUB: not implemented"

	// GetLogger of ZapLoggerFactory ignores the given name and just return the default logger.
	return *new(Logger)
}

func (r *ZapLoggerFactory) GetLogger(name string) Logger {
	_ = "STUB: not implemented"

	// ZapLogger is wrapper of zap
	return *new(Logger)
}

type ZapLogger struct {
	sugaredLogger *zap.SugaredLogger
}

// Debug is log at debug level
func (z *ZapLogger) Debug(args ...interface{}) { _ = "STUB: not implemented"; return }

// Debugf is log at debug level with fmt.Printf-like formatting
func (z *ZapLogger) Debugf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Info is log at info level
func (z *ZapLogger) Info(args ...interface{}) { _ = "STUB: not implemented"; return }

// Infof is log at info level with fmt.Printf-like formatting
func (z *ZapLogger) Infof(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Warn is log at warning level
func (z *ZapLogger) Warn(args ...interface{}) { _ = "STUB: not implemented"; return }

// Warnf is log at warning level with fmt.Printf-like formatting
func (z *ZapLogger) Warnf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Error is log at error level
func (z *ZapLogger) Error(args ...interface{}) { _ = "STUB: not implemented"; return }

// Errorf is log at error level with fmt.Printf-like formatting
func (z *ZapLogger) Errorf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Fatal is log at fatal level, then terminate process (irrecoverable)
func (z *ZapLogger) Fatal(args ...interface{}) { _ = "STUB: not implemented"; return }

// Fatalf is log at fatal level with fmt.Printf-like formatting, then terminate process (irrecoverable)
func (z *ZapLogger) Fatalf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// Panic is log at panic level, then panic (recoverable)
func (z *ZapLogger) Panic(args ...interface{}) { _ = "STUB: not implemented"; return }

// Panicf is log at panic level with fmt.Printf-like formatting, then panic (recoverable)
func (z *ZapLogger) Panicf(format string, args ...interface{}) { _ = "STUB: not implemented"; return }

// With returns a logger with the specified key-value pair set, to be logged in a subsequent normal logging call
func (z *ZapLogger) With(args ...interface{}) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}

// NoopLogger is wrapper of noop logger
type NoopLogger struct{}

// Debug is log at debug level
func (z *NoopLogger) Debug(args ...interface{}) {
	_ = "STUB: not implemented"

	// Debugf is log at debug level with fmt.Printf-like formatting
	return
}

func (z *NoopLogger) Debugf(format string, args ...interface{}) {
	_ = "STUB: not implemented"

	// Info is log at info level
	return
}

func (z *NoopLogger) Info(args ...interface{}) {
	_ = "STUB: not implemented"

	// Infof is log at info level with fmt.Printf-like formatting
	return
}

func (z *NoopLogger) Infof(format string, args ...interface{}) {
	_ = "STUB: not implemented"

	// Warn is log at warning level
	return
}

func (z *NoopLogger) Warn(args ...interface{}) {
	_ = "STUB: not implemented"

	// Warnf is log at warning level with fmt.Printf-like formatting
	return
}

func (z *NoopLogger) Warnf(format string, args ...interface{}) {
	_ = "STUB: not implemented"

	// Error is log at error level
	return
}

func (z *NoopLogger) Error(args ...interface{}) {
	_ = "STUB: not implemented"

	// Errorf is log at error level with fmt.Printf-like formatting
	return
}

func (z *NoopLogger) Errorf(format string, args ...interface{}) {
	_ = "STUB: not implemented"

	// Fatal is log at fatal level, then terminate process (irrecoverable)
	return
}

func (z *NoopLogger) Fatal(args ...interface{}) {
	_ = "STUB: not implemented"

	// Fatalf is log at fatal level with fmt.Printf-like formatting, then terminate process (irrecoverable)
	return
}

func (z *NoopLogger) Fatalf(format string, args ...interface{}) {
	_ = "STUB: not implemented"

	// Panic is log at panic level, then panic (recoverable)
	return
}

func (z *NoopLogger) Panic(args ...interface{}) {
	_ = "STUB: not implemented"

	// Panicf is log at panic level with fmt.Printf-like formatting, then panic (recoverable)
	return
}

func (z *NoopLogger) Panicf(format string, args ...interface{}) {
	_ = "STUB: not implemented"

	// With returns a logger with the specified key-value pair set, to be logged in a subsequent normal logging call
	return
}

func (z *NoopLogger) With(args ...interface{}) Logger {
	_ = "STUB: not implemented"
	return *new(Logger)
}
