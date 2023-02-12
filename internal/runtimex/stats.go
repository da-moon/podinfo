package runtimex

// Stats struct encapulates the runtime statistics.
//
//go:generate gomodifytags -override -file $GOFILE -struct Stats -add-tags json,mapstructure -w -transform snakecase
type Stats struct {
	Memory *MemStats `json:"memory" mapstructure:"memory"`
}

// GetStats is the Stats constructor
func GetStats() *Stats {
	result := &Stats{}
	result.Memory = GetMemoryStats()
	return result
}
