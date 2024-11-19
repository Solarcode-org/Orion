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

	"github.com/Solarcode-org/Orion/ast"
	"github.com/Solarcode-org/Orion/lexer"
	"github.com/Solarcode-org/Orion/parser"
	log "github.com/sirupsen/logrus"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

// ParsedFrom uses the lexer to tokenize the input source and
// returns the AST formed by parsing the tokenized content.
func ParsedFrom(src []byte) ([]*ast.Expr, []*parser.Error, error) {
	log.Tracef("started function `lib.GetAbstractSyntaxTree` with argument `src`=`%s`\n", src)

	lex := lexer.New([]rune(string(src)))
	bsrSet, errs := parser.Parse(lex)

	if len(errs) > 0 {
		return nil, errs, nil
	}

	ast, err := AbstractSyntaxTree(bsrSet.GetRoot())
	if err != nil {
		return nil, nil, err
	}

	log.Traceln("successfully ended function `lib.GetAbstractSyntaxTree`")
	return ast, nil, nil
}

// CheckErr checks if an error value is not nil and runs [log.Fatalln] for the error.
func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// CheckIfNoArgs checks if the list of args is empty. If not, it returns an error.
func CheckIfNoArgs(data []*ast.Expr) error {
	if len(data) > 0 {
		return fmt.Errorf("expected no arguments, got %d", len(data))
	}

	return nil
}

// CheckIfExactArgs checks if the number of args matches `amount`. If not, it returns an error
func CheckIfExactArgs(data []*ast.Expr, amount int) error {
	if len(data) != amount {
		return fmt.Errorf("expected exactly %d arguments, got %d", amount, len(data))
	}

	return nil
}

// RunFunc contains the functionality for running an Orion function.
func RunFunc(funcCall ast.Expr, functions map[string]func([]*ast.Expr) (ast.Expr, error)) (ast.Expr, error) {
	if function, ok := functions[funcCall.Id]; ok {
		value, err := function(funcCall.Args)
		if err != nil {
			return ast.Expr{}, fmt.Errorf("%s: %w", funcCall.Id, err)
		}

		return value, nil
	}

	caser := cases.Title(language.AmericanEnglish)

	if _, ok := functions[caser.String(funcCall.Id)]; ok {
		return ast.Expr{}, fmt.Errorf("could not find function: %s, did you mean: %s", funcCall.Id, caser.String(funcCall.Id))
	}

	return ast.Expr{}, fmt.Errorf("could not find function: %s, you may have forgotten to add a module prefix", funcCall.Id)
}

// FailParse prints all the errors with the same line number as `errs[0]` and exits with code 1.
func FailParse(errs []*parser.Error) {
	errors := "parse errors:\n"

	ln := errs[0].Line
	for _, err := range errs {
		if err.Line == ln {
			errors += fmt.Sprintln("  ", err)
		}
	}

	log.Fatal(errors)
}
