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

// Package ast is the Abstract Syntax Tree module for Orion.
package ast

// An ExprType conveys the expression type of an [Expr].
type ExprType int

// The type of an [Expr] is one of these types.
const (
	Expr_String ExprType = iota
	Expr_FuncCall
	Expr_Number
	Expr_Variable
	Expr_VariableDef
	Expr_VariableTypeDef
)

// An Expr represents an expression in Orion.
// The zero value for Expr is an empty expression i.e the "nil" value.
type Expr struct {
	Type ExprType `json:"type,omitempty"` // The type of this expression [ExprType].
	Id   string   `json:"id,omitempty"`   // Either the value of any data or the name of a function.
	Args []*Expr  `json:"args,omitempty"` // If this is a function or argument list, this is filled.
}
