package cmd_test

import (
	"testing"

	"github.com/Solarcode-org/Orion/cmd"
)

func TestIor(t *testing.T) {
	cmd.RootCmd.SetArgs([]string{"ior", "--help"})
	cmd.Execute()
}

func BenchmarkMainExample(b *testing.B) {
	cmd.RootCmd.SetArgs([]string{"ior", "../_examples/bench.or"})

	for i := 0; i < b.N; i++ {
		cmd.Execute()
	}
}
