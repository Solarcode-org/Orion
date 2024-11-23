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

// The standard library for Orion.
// It contains all of the built-in functions, variables, types, etc. of Orion.
package builtins

import (
	"fmt"

	"github.com/Solarcode-org/Orion/ast"
)

// A FunctionsType is a map of all functions ([OrionFunction]) in Orion.
type FunctionsType = map[string]OrionFunction

// A VariablesType is a map of all variables ([OrionVar]) in Orion.
type VariablesType = map[string]OrionVar

// An OrionFunction represents a function in Orion.
type OrionFunction = func([]*ast.Expr) (ast.Expr, error)

// An OrionVar represents a variable in Orion.
type OrionVar = ast.Expr

// Functions map containing all functions in Orion.
var Functions FunctionsType

// Variables map containing all variables in Orion.
var Variables VariablesType

// Initialize the functions map for use.
// NOTE: This must be called before any usage of the [Functions] variable.
func MakeFunctions() {
	Functions = make(FunctionsType)

	addFmt(Functions)
	addModGetter(Functions)
	addArithmetic(Functions)
}

// Initialize the variables map for use.
// NOTE: This must be called before any usage of the [Variables] variable.
func MakeVariables(verbose uint8) {
	Variables = make(VariablesType)

	Variables["@verbose"] = ast.Expr{
		Type: ast.Expr_Number,
		Id:   fmt.Sprint(verbose),
	}
}
