package lib_test

import (
	"testing"

	"github.com/Solarcode-org/Orion/lib"
)

func TestParseEmpty(t *testing.T) {
	contents := []byte{}

	_, errs := lib.GetAbstractSyntaxTree(contents)
	if len(errs) == 0 {
		t.Fatal("expected failure on empty file")
	}
}

func TestParsePackageOnly(t *testing.T) {
	contents := []byte("package \"__testing/untitled\"")

	_, errs := lib.GetAbstractSyntaxTree(contents)
	if len(errs) == 0 {
		t.Fatal("expected failure on file with only package name.")
	}
}

var validFile = []byte(`
package "__testing/untitled"

Println("Hello,", "world")
`)

func TestParseValid(t *testing.T) {
	_, errs := lib.GetAbstractSyntaxTree(validFile)
	if len(errs) > 0 {
		t.Log("Parse errors:")

		ln := errs[0].Line
		for _, err := range errs {
			if err.Line == ln {
				t.Log("  ", err)
			}
		}
		t.FailNow()
	}
}

func BenchmarkParseValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		lib.GetAbstractSyntaxTree(validFile)
	}
}
