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

import "github.com/Solarcode-org/Orion/lib/ast"

type FunctionsType = map[string]OrionFunction
type OrionFunction = func(ast.DataList) (ast.Data, error)

var Functions FunctionsType

func MakeFunctions() {
	Functions = make(FunctionsType)

	add_fmt_mod(Functions)
	add_modgetter(Functions)
	add_arithmetic_mod(Functions)
}
