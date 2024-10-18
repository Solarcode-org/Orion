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

package lib

import (
	"github.com/Solarcode-org/Orion/lib/ast"
	"github.com/Solarcode-org/Orion/lib/lexer"
	"github.com/Solarcode-org/Orion/lib/parser"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// GetAbstractSyntaxTree returns the AST formed by the input source.
func GetAbstractSyntaxTree(src []byte) (astree ast.FuncCallList, err error) {
	log.Tracef("started function `lib.GetAbstractSyntaxTree` with argument `src`=`%s`\n", src)

	s := lexer.NewLexer(src)
	p := parser.NewParser()

	a, err := p.Parse(s)

	if err != nil {
		log.Traceln("could not parse `src`")
		return nil, err
	}

	log.Traceln("successfully ended function `lib.GetAbstractSyntaxTree`")
	return a.(ast.FuncCallList), nil
}

// HandleFatal checks if an error value is not nil and runs `log.Fatalln` for the error.
func HandleFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func NoArgs(funcname string, data ast.DataList) {
	if len(data) > 0 {
		log.Fatalf("%s: expected no arguments, got %d\n", funcname, len(data))
	}
}

func ExactArgs(funcname string, data ast.DataList, amount int) {
	if len(data) != amount {
		log.Fatalf("%s: expected exactly %d arguments, got %d\n", funcname, amount, len(data))
	}
}

func RunFunc(funcCall ast.FuncCall, functions map[string]func(ast.DataList) (ast.Data, error)) ast.Data {
	if function, ok := functions[funcCall.Name]; ok {
		value, err := function(funcCall.Args)
		if err != nil {
			log.Fatalf("%s: %s\n", funcCall.Name, err)
		}

		return value
	}

	caser := cases.Title(language.AmericanEnglish)

	if _, ok := functions[caser.String(funcCall.Name)]; ok {
		log.Fatalf("Could not find function: %s\nDid you mean: %s?\n", funcCall.Name, caser.String(funcCall.Name))
	}

	log.Fatalf("Could not find function: %s\nMaybe you forgot to add a module prefix?\n", funcCall.Name)
	return ast.None
}
