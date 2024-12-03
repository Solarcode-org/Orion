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
	"fmt"

	"github.com/Solarcode-org/Orion/ast"
	"github.com/Solarcode-org/Orion/utils"
)

// ParsedArgs converts all the arguments of a function into usable values.
// For example: evaluating function values before use.
func ParsedArgs(data []*ast.Expr) ([]*ast.Expr, error) {
	args := make([]*ast.Expr, 0, len(data))

	for i := 0; i < len(data); i++ {
		datum := data[i]

		if datum.Type == ast.Expr_FuncCall {
			value, err := utils.RunFunc(*datum, Functions)
			if err != nil {
				return nil, err
			}

			args = append(args, &value)

			continue
		}

		if datum.Type == ast.Expr_Variable {
			if value, ok := Variables[datum.Id]; ok {
				args = append(args, &value)
			} else {
				return nil, fmt.Errorf("could not find name: %s", datum.Id)
			}

			continue
		}

		args = append(args, datum)
	}

	return args, nil
}

// checkIfNoArgs checks if the list of args is empty. If not, it returns an error.
func checkIfNoArgs(data []*ast.Expr) error {
	if len(data) > 0 {
		return fmt.Errorf("expected no arguments, got %d", len(data))
	}

	return nil
}

// checkIfExactArgs checks if the number of args matches `amount`. If not, it returns an error
func checkIfExactArgs(data []*ast.Expr, amount int) error {
	if len(data) != amount {
		return fmt.Errorf("expected exactly %d arguments, got %d", amount, len(data))
	}

	return nil
}
