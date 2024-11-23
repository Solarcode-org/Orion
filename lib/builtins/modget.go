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
	"strings"

	"github.com/Solarcode-org/Orion/ast"
)

// addModGetter adds import functionality to Orion.
func addModGetter(functions FunctionsType) {
	functions["get"] = func(data []*ast.Expr) (ast.Expr, error) {
		for i := 0; i < len(data); i++ {
			module := data[i]

			if module.Type != ast.Expr_String /* && module.Type != ast.Ident */ {
				return ast.Expr{}, fmt.Errorf("expected string as module arguments")
			}

			keys := make([]string, 0, len(functions))
			for k := range functions {
				keys = append(keys, k)
			}

			for _, key := range keys {
				if strings.Split(key, ".")[0] == module.Id {
					functions[strings.Split(key, ".")[1]] = functions[key]
				}
			}
		}

		return ast.Expr{}, nil
	}
}
