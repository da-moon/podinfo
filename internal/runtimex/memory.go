package runtimex

import (
	"runtime"
	"strconv"
)

// MemStats represents memory stats.
//
//go:generate gomodifytags -override -file $GOFILE -struct MemStats -add-tags json,mapstructure -w -transform snakecase
type MemStats struct {
	// measures the virtual address space reserved by the Go
	// runtime for the heap, stacks, and other internal data structures
	Alloc int `json:"alloc" mapstructure:"alloc"`
	// total_alloc increases as heap objects are allocated, but
	// unlike alloc it does not decrease when objects are freed.
	TotalAlloc int `json:"total_alloc" mapstructure:"total_alloc"`
	// measures the virtual address space reserved by the Go
	// runtime for the heap, stacks, and other internal data
	Sys   int `json:"sys" mapstructure:"sys"`
	NumGc int `json:"num_gc" mapstructure:"num_gc"`
}

// GetMemoryStats returns current memory usage
func GetMemoryStats() *MemStats {
	bToMb := func(b uint64) int {
		return int(b / 1024 / 1024)
	}
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	result := &MemStats{}
	result.Alloc = bToMb(m.Alloc)
	result.TotalAlloc = bToMb(m.TotalAlloc)
	result.Sys = bToMb(m.Sys)
	result.NumGc = int(m.NumGC)
	return result
}

// String returns a string representation of the memory stats.
func (m *MemStats) String() string {
	return "Alloc: " + strconv.Itoa(m.Alloc) + " MB, TotalAlloc: " + strconv.Itoa(m.TotalAlloc) + " MB, Sys: " + strconv.Itoa(m.Sys) + " MB, NumGC: " + strconv.Itoa(m.NumGc)
}
