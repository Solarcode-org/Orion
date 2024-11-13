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
	"github.com/Solarcode-org/Orion/ast"
	"github.com/Solarcode-org/Orion/lib"
)

// EvalArgs converts all the arguments of a function into usable values.
// For example: evaluating function values befor use.
func EvalArgs(data []*ast.Expr) []*ast.Expr {
	args := make([]*ast.Expr, 0, len(data))

	for i := 0; i < len(data); i++ {
		datum := data[i]

		if datum.Type == ast.Expr_FuncCall {
			value := lib.RunFunc(*datum, Functions)
			args = append(args, &value)

			continue
		}

		args = append(args, datum)
	}

	return args
}
