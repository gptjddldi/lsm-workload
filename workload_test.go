package lsm_workload

import (
	"testing"
)

func TestLsmWorkloadGenerate(t *testing.T) {
	// Create a new LsmWorkload
	lw := NewLsmWorkload("number", 50000000, 5000, 100, 0.5)

	lw.Generate()
}
