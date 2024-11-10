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
	"fmt"
	"os"

	"github.com/Solarcode-org/Orion/ast"
	"github.com/Solarcode-org/Orion/lexer"
	"github.com/Solarcode-org/Orion/parser"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// GetAbstractSyntaxTree returns the AST formed by the input source.
func GetAbstractSyntaxTree(src []byte) ([]*ast.Expr, []*parser.Error) {
	log.Tracef("started function `lib.GetAbstractSyntaxTree` with argument `src`=`%s`\n", src)

	lex := lexer.New([]rune(string(src)))
	bsrSet, errs := parser.Parse(lex)

	if len(errs) > 0 {
		return nil, errs
	}

	ast := buildRoot(bsrSet.GetRoot())

	log.Traceln("successfully ended function `lib.GetAbstractSyntaxTree`")
	return ast, nil
}

// HandleFatal checks if an error value is not nil and runs `log.Fatalln` for the error.
func HandleFatal(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func NoArgs(funcname string, data []*ast.Expr) {
	if len(data) > 0 {
		log.Fatalf("%s: expected no arguments, got %d\n", funcname, len(data))
	}
}

func ExactArgs(funcname string, data []*ast.Expr, amount int) {
	if len(data) != amount {
		log.Fatalf("%s: expected exactly %d arguments, got %d\n", funcname, amount, len(data))
	}
}

func RunFunc(funcCall ast.Expr, functions map[string]func([]*ast.Expr) (ast.Expr, error)) ast.Expr {
	if function, ok := functions[funcCall.Id]; ok {
		value, err := function(funcCall.Args)
		if err != nil {
			log.Fatalf("%s: %s\n", funcCall.Id, err)
		}

		return value
	}

	caser := cases.Title(language.AmericanEnglish)

	if _, ok := functions[caser.String(funcCall.Id)]; ok {
		log.Fatalf("Could not find function: %s\nDid you mean: %s?\n", funcCall.Id, caser.String(funcCall.Id))
	}

	log.Fatalf("Could not find function: %s\nMaybe you forgot to add a module prefix?\n", funcCall.Id)
	return ast.Expr{}
}

// Print all the errors with the same line number as errs[0] and exit(1)
func FailParse(errs []*parser.Error) {
	fmt.Println("Parse Errors:")
	ln := errs[0].Line
	for _, err := range errs {
		if err.Line == ln {
			fmt.Println("  ", err)
		}
	}
	os.Exit(1)
}
