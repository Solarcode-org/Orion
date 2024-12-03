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

// Package lib_test is for testing and benchmarking the parsing capability of Orion.
package lib_test

import (
	"testing"

	"github.com/Solarcode-org/Orion/lib"
)

// TestParseEmpty tries to parse an empty file. This test *must* fail.
func TestParseEmpty(t *testing.T) {
	contents := []byte{}

	_, errs, _ := lib.ParsedFrom(contents)
	if len(errs) == 0 {
		t.Fatal("expected failure on empty file")
	}
}

// TestParsePackageOnly tries to parse a file with only the package declaration line. This test *must* fail.
func TestParsePackageOnly(t *testing.T) {
	contents := []byte("package \"__testing/untitled\"")

	_, parseErrs, _ := lib.ParsedFrom(contents)
	if len(parseErrs) == 0 {
		t.Fatal("expected failure on file with only package name.")
	}
}

// validFile contains the contents for a perfectly valid but minimal file.
var validFile = []byte(`
package "__testing/untitled"
world := "world"
one: number = "1"
Println("Hello", world)
Println(one + 1)
Println(one / 1)
`)

// TestParseValid tries to parse [validFile].
func TestParseValid(t *testing.T) {
	_, parseErrs, err := lib.ParsedFrom(validFile)
	if len(parseErrs) > 0 {
		t.Log("parse errors:")

		ln := parseErrs[0].Line
		for _, err := range parseErrs {
			if err.Line == ln {
				t.Log("  ", err)
			}
		}

		t.Fail()
	}
	if err != nil {
		t.Error(err)
	}
}

// BenchmarkParseValid measures the performance of the parsing capability
// by trying to parse [validFile].
func BenchmarkParseValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib.ParsedFrom(validFile)
	}
}
