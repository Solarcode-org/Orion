package bytecode_test

import (
	"os"
	"path"
	"testing"

	"github.com/Solarcode-org/Orion/lib"
	"github.com/Solarcode-org/Orion/lib/bytecode"
)

func TestEncoder(t *testing.T) {
	contents, err := os.ReadFile(path.Join("..", "..", "_examples", "bench.or"))
	if err != nil {
		t.Error(err)
	}

	ast, parseErrs, err := lib.ParsedFrom(contents)
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

	if _, err := bytecode.EncodedSyntaxTree(ast); err != nil {
		t.Error(err)
	}
}

func BenchmarkEncoder(b *testing.B) {
	contents, _ := os.ReadFile(path.Join("..", "..", "_examples", "bench.or"))
	ast, _, _ := lib.ParsedFrom(contents)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bytecode.EncodedSyntaxTree(ast)
	}
}

func TestDecoder(t *testing.T) {
	contents, err := os.ReadFile(path.Join("..", "..", "_examples", "bench.or.ast"))
	if err != nil {
		t.Error(err)
	}

	if _, err := bytecode.DecodedSyntaxTree(contents); err != nil {
		t.Error(err)
	}
}

func BenchmarkDecoder(b *testing.B) {
	contents, _ := os.ReadFile(path.Join("..", "..", "_examples", "bench.or.ast"))
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		bytecode.DecodedSyntaxTree(contents)
	}
}
