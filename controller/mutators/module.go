package mutators

import (
	"github.com/uber/aresdb/cluster/kvstore"
	"github.com/uber/aresdb/controller/mutators/common"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// Result holds all mutators
type Result struct {
	fx.Out

	JobMutator                 common.JobMutator
	NamespaceMutator           common.NamespaceMutator
	SubscriberMutator          common.SubscriberMutator
	SchemaMutator              common.TableSchemaMutator
	IngestionAssignmentMutator common.IngestionAssignmentMutator
	MembershipMutator          common.MembershipMutator
	PlacementMutator           common.PlacementMutator
	EnumMutator                common.EnumMutator
}

// Params represents params to initialize all mutators
type Params struct {
	fx.In

	EtcdClient *kvstore.EtcdClient
	Logger     *zap.SugaredLogger
}

// InitMutators initialize mutators
func InitMutators(param Params) Result { _ = "STUB: not implemented"; return *new(Result) }

// Module defines business module
var Module = fx.Provide(InitMutators)
