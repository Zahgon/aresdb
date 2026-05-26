package tasks

import (
	"github.com/uber-go/tally"
	"github.com/uber/aresdb/cluster/kvstore"
	mutatorCom "github.com/uber/aresdb/controller/mutators/common"
	"go.uber.org/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// EtcdTaskParams wraps etcd params with fx
type EtcdTaskParams struct {
	fx.In

	ConfigProvider config.Provider
	Logger         *zap.SugaredLogger
	Scope          tally.Scope

	EtcdClient         *kvstore.EtcdClient
	NamespaceMutator   mutatorCom.NamespaceMutator
	JobMutator         mutatorCom.JobMutator
	SchemaMutator      mutatorCom.TableSchemaMutator
	SubscriberMutator  mutatorCom.SubscriberMutator
	AssignmentsMutator mutatorCom.IngestionAssignmentMutator
}

// InvokeEtcdTask invoke etcd based ingestion assignment task
func InvokeEtcdTask(params EtcdTaskParams, Lifecycle fx.Lifecycle) {
	_ = "STUB: not implemented"
	return
}

var Module = fx.Invoke(InvokeEtcdTask)
