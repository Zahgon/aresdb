package bootstrap

import (
	"sync"
)

type status int

const (
	notApplicable status = iota
	copyNeeded
	copyFinished
)

type bootstrapDetails struct {
	*sync.RWMutex

	Source         string         `json:"source"`
	BootstrapStage BootstrapStage `json:"stage"`
	StartedAt      int64          `json:"startedAt"`
	NumColumns     int            `json:"numColumns"`
	// map from batch id to slice of all columns
	Batches map[int32][]status `json:"batches"`
}

func NewBootstrapDetails() BootstrapDetails {
	_ = "STUB: not implemented"
	return *new(BootstrapDetails)
}

func (b *bootstrapDetails) SetSource(source string) { _ = "STUB: not implemented"; return }

func (b *bootstrapDetails) SetNumColumns(numColumns int) { _ = "STUB: not implemented"; return }

func (b *bootstrapDetails) AddVPToCopy(batchID int32, columnID uint32) {
	_ = "STUB: not implemented"
	return
}

func (b *bootstrapDetails) SetBootstrapStage(stage BootstrapStage) {
	_ = "STUB: not implemented"
	return
}

func (b *bootstrapDetails) MarkVPFinished(batchID int32, columnID uint32) {
	_ = "STUB: not implemented"
	return
}

func (b *bootstrapDetails) Clear() { _ = "STUB: not implemented"; return }

func (b *bootstrapDetails) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
