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

package builtins

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Solarcode-org/Orion/ast"
)

// addFmt adds the standard input/output module ("fmt"), along with the
// Print, Println and Input functions.
func addFmt(functions FunctionsType) {
	functions["Println"] = fmt_println
	functions["Print"] = fmt_print
	functions["Input"] = fmt_input
	functions["fmt.Println"] = fmt_println
	functions["fmt.Print"] = fmt_print
	functions["fmt.Input"] = fmt_input
	functions["fmt.Join"] = fmt_join
}

func fmt_print(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}

	for i := 0; i < len(args); i++ {
		datum := args[i]
		fmt.Print(datum.Id)

		fmt.Print(" ")
	}

	return ast.Expr{}, nil
}

func fmt_println(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}

	for i := 0; i < len(args); i++ {
		datum := args[i]

		fmt.Print(datum.Id)

		fmt.Print(" ")
	}

	fmt.Println()

	return ast.Expr{}, nil
}

func fmt_input(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}
	if err = checkIfExactArgs(args, 1); err != nil {
		return ast.Expr{}, err
	}

	if args[0].Type == ast.Expr_String {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(args[0].Id)

		ans, err := reader.ReadString('\n')
		if err != nil {
			return ast.Expr{}, err
		}

		ans = strings.TrimSpace(ans)

		return ast.Expr{
			Id:   ans,
			Type: ast.Expr_String,
		}, nil
	}

	return ast.Expr{}, fmt.Errorf("expected prompt to be of type string")
}

func fmt_join(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}
	joined := ""

	for i := 0; i < len(data); i++ {
		datum := args[i]

		joined += datum.Id
	}

	return ast.Expr{
		Id:   joined,
		Type: ast.Expr_String,
	}, nil
}
