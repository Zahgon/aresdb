package main

import (
	"github.com/spf13/cobra"
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/subscriber/common/job"
	"github.com/uber/aresdb/subscriber/common/rules"
	"github.com/uber/aresdb/subscriber/config"
	"github.com/uber/aresdb/utils"
	cfgfx "go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Params struct {
	fx.In

	Lifecycle fx.Lifecycle
}

type Result struct {
	fx.Out

	Environment utils.EnvironmentContext
	Provider    cfgfx.Provider
	Logger      *zap.Logger
	Scope       tally.Scope
	Consumer    job.NewConsumer
	Decoder     job.NewDecoder
	Sink        job.NewSink
}

func main() {
	module := fx.Provide(Init)
	Execute(module, config.Module, rules.Module, job.Module)
}

func Execute(opts ...fx.Option) { _ = "STUB: not implemented"; return }

func addFlags(cmd *cobra.Command) { _ = "STUB: not implemented"; return }

func newDefaultEnvironmentCtx() utils.EnvironmentContext {
	_ = "STUB: not implemented"
	return *new(utils.EnvironmentContext)
}

func newDefaultConfig() cfgfx.Provider { _ = "STUB: not implemented"; return *new(cfgfx.Provider) }

func newDefaultLogger(params Params, env utils.EnvironmentContext) *zap.Logger {
	_ = "STUB: not implemented"
	return nil
}

func newDefaultScope() tally.Scope { _ = "STUB: not implemented"; return *new(tally.Scope) }

func Init(params Params) Result { _ = "STUB: not implemented"; return *new(Result) }
