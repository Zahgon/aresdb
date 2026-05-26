package utils

const (
	assetFileName = "asset-manifest.json"
)

// ResourceMapper maps path to a static resource built by react
type ResourceMapper struct {
	assetMap map[string]string
}

// NewResourceMapper creates a new resource mapper
func NewResourceMapper(buildPath string) (mapper ResourceMapper) {
	_ = "STUB: not implemented"
	return *new(ResourceMapper)
}

// keep quiet for now

// keep quiet for now

// keep quiet for now

// Map returns physical path of resource
func (m ResourceMapper) Map(assetName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
