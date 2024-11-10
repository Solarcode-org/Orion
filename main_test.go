package main_test

import (
	"testing"

	"github.com/Solarcode-org/Orion/cmd"
)

func TestMainHelp(t *testing.T) {
	cmd.Execute()
}

func BenchmarkMainHelp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cmd.Execute()
	}
}
