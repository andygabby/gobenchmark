package main

import (
	"testing"
)

func TestRandInt(t *testing.T) {
	got := RandInt()
	if got % 1 != 0 {
		t.Errorf("RandInt() = %d; want 0", got)
	}
}

func BenchmarkRandInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		RandInt()
	}
}
