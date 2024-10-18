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

package lib_test

import (
	"fmt"
	"testing"

	"github.com/Solarcode-org/Orion/lib/ast"
	"github.com/Solarcode-org/Orion/lib/lexer"
	"github.com/Solarcode-org/Orion/lib/parser"
)

func TestPass(t *testing.T) {
	sml, err := test([]byte("Println(\"Hi\")"))
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Printf("output: %v\n", sml)
}

func TestFail(t *testing.T) {
	_, err := test([]byte("Println(\")"))
	if err == nil {
		t.Fatal("expected parse error")
	} else {
		fmt.Printf("Parsing failed as expected: %v\n", err)
	}
}

func test(src []byte) (astree ast.FuncCallList, err error) {
	fmt.Printf("input: %s\n", src)
	s := lexer.NewLexer(src)
	p := parser.NewParser()
	a, err := p.Parse(s)
	if err == nil {
		astree = a.(ast.FuncCallList)
	}
	return
}
