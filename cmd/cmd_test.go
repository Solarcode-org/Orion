/*
Copyright © 2024 Arnab Phukan <iamarnab.phukan@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package cmd_test is for testing and benchmarking the commands of the Orion CLI.
package cmd_test

import (
	"testing"

	"github.com/Solarcode-org/Orion/cmd"
)

// TestIor tests the help flag
func TestIor(t *testing.T) {
	cmd.RootCmd.SetArgs([]string{"ior", "--help"})
	cmd.Execute()
}

// BenchmarkExample measures the performance of the benchmarking example ([_examples/bench.or])
func BenchmarkExample(b *testing.B) {
	cmd.RootCmd.SetArgs([]string{"ior", "../_examples/bench.or"})

	for i := 0; i < b.N; i++ {
		cmd.Execute()
	}
}
