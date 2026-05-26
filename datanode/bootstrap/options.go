package bootstrap

import (
	"time"
)

const (
	defaultBootstrapSessionTTL              = int64(5 * time.Minute)
	defaultMaxConcurrentTableShards         = 8
	defaultMaxCocurrentSessionPerTableShard = 2
)

// options implements bootstrap Options
type options struct {
	maxConcurrentTableShards          int
	maxConcurrentStreamsPerTableShard int
	bootstrapSessionTTL               int64
}

func (o *options) MaxConcurrentTableShards() int { _ = "STUB: not implemented"; return 0 }

func (o *options) SetMaxConcurrentShards(numShards int) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

func (o *options) MaxConcurrentStreamsPerTableShards() int { _ = "STUB: not implemented"; return 0 }

func (o *options) SetMaxConcurrentStreamsPerTableShards(numStreams int) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

// BootstrapSessionTTL returns the ttl for bootstrap session
func (o *options) BootstrapSessionTTL() int64 { _ = "STUB: not implemented"; return 0 }

// SetBootstrapSessionTTL sets the session ttl for bootstrap session
func (o *options) SetBootstrapSessionTTL(ttl int64) Options {
	_ = "STUB: not implemented"
	return *new(Options)
}

// NewOptions returns bootstrap default options
func NewOptions() Options { _ = "STUB: not implemented"; return *new(Options) }
