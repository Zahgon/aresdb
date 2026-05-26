package etcd

import (
	"github.com/m3db/m3/src/cluster/placement"
	"github.com/m3db/m3/src/cluster/services"
	"github.com/uber/aresdb/cluster/kvstore"
	"github.com/uber/aresdb/controller/mutators/common"
)

type placementMutator struct {
	client *kvstore.EtcdClient
}

func (p *placementMutator) getServiceID(namespace string) services.ServiceID {
	_ = "STUB: not implemented"
	return *new(services.ServiceID)
}

func validateAllAvailable(p placement.Placement) error { _ = "STUB: not implemented"; return nil }

// NewPlacementMutator creates mutator for placement
func NewPlacementMutator(client *kvstore.EtcdClient) common.PlacementMutator {
	_ = "STUB: not implemented"
	return *new(common.PlacementMutator)
}

func checkNumShardsIsPowerOfTwo(numShards int) bool { _ = "STUB: not implemented"; return false }

func (p *placementMutator) placementOptions() placement.Options {
	_ = "STUB: not implemented"
	return *new(placement.Options)
}

// if we specify more than one new instance, we want to add them all

// for now we want to make sure replacement does not affect existing instances not being replaced

func (p *placementMutator) BuildInitialPlacement(namespace string, numShards int, numReplica int, instances []placement.Instance) (placement.Placement, error) {
	_ = "STUB: not implemented"
	return *new(placement.Placement), nil
}

func (p *placementMutator) GetCurrentPlacement(namespace string) (placement.Placement, error) {
	_ = "STUB: not implemented"
	return *new(placement.Placement), nil
}

func (p *placementMutator) AddInstance(namespace string, instances []placement.Instance) (placement.Placement, error) {
	_ = "STUB: not implemented"
	return *new(placement.Placement), nil
}

func (p *placementMutator) ReplaceInstance(namespace string, leavingInstances []string, newInstances []placement.Instance) (placement.Placement, error) {
	_ = "STUB: not implemented"
	return *new(placement.Placement), nil
}

func (p *placementMutator) RemoveInstance(namespace string, leavingInstances []string) (placement.Placement, error) {
	_ = "STUB: not implemented"
	return *new(placement.Placement), nil
}

func (p *placementMutator) MarkNamespaceAvailable(namespace string) (placement.Placement, error) {
	_ = "STUB: not implemented"
	return *new(placement.Placement), nil
}

func (p *placementMutator) MarkInstanceAvailable(namespace string, instance string) (placement.Placement, error) {
	_ = "STUB: not implemented"
	return *new(placement.Placement), nil
}

func (p *placementMutator) MarkShardsAvailable(namespace string, instance string, shards []uint32) (placement.Placement, error) {
	_ = "STUB: not implemented"
	return *new(placement.Placement), nil
}
