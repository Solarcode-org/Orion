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

	"github.com/shopspring/decimal"
)

// addArithmetic adds the `math` module and arithmetic functions.
func addArithmetic(functions FunctionsType) {
	functions["Sum"] = math_sum
	functions["Difference"] = math_difference
	functions["Product"] = math_product
	functions["Quotient"] = math_quotient
	functions["math.Sum"] = math_sum
	functions["math.Difference"] = math_difference
	functions["math.Product"] = math_product
	functions["math.Quotient"] = math_quotient
	functions["math.Round"] = math_round
	functions["math.Ceil"] = math_ceil
	functions["math.Floor"] = math_floor
}

func math_sum(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}
	result := decimal.Zero

	for i := 0; i < len(args); i++ {
		datum := args[i]

		if datum.Type != ast.Expr_Number {
			return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
		}

		num, err := decimal.NewFromString(datum.Id)
		if err != nil {
			return ast.Expr{}, err
		}

		result = result.Add(num)
	}

	return ast.Expr{
		Id:   result.String(),
		Type: ast.Expr_Number,
	}, nil
}

func math_difference(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}
	if err = checkIfExactArgs(args, 2); err != nil {
		return ast.Expr{}, err
	}

	if data[0].Type != ast.Expr_Number {
		return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
	}
	result, err := decimal.NewFromString(args[0].Id)
	if err != nil {
		return ast.Expr{}, err
	}

	if data[1].Type != ast.Expr_Number {
		return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
	}
	subtrahend, err := decimal.NewFromString(args[1].Id)
	if err != nil {
		return ast.Expr{}, err
	}

	result = result.Sub(subtrahend)

	return ast.Expr{
		Id:   result.String(),
		Type: ast.Expr_Number,
	}, nil
}

func math_product(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}

	if args[0].Type != ast.Expr_Number {
		return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
	}
	result, err := decimal.NewFromString(args[0].Id)
	if err != nil {
		return ast.Expr{}, err
	}

	for i := 0; i < len(args)-1; i++ {
		datum := args[i+1]

		if datum.Type != ast.Expr_Number {
			return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
		}

		num, err := decimal.NewFromString(datum.Id)
		if err != nil {
			return ast.Expr{}, err
		}

		result = result.Mul(num)
	}

	return ast.Expr{
		Id:   result.String(),
		Type: ast.Expr_Number,
	}, nil
}

func math_quotient(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}
	if err = checkIfExactArgs(args, 2); err != nil {
		return ast.Expr{}, err
	}

	if data[0].Type != ast.Expr_Number {
		return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
	}
	result, err := decimal.NewFromString(args[0].Id)
	if err != nil {
		return ast.Expr{}, err
	}

	if data[1].Type != ast.Expr_Number {
		return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
	}
	divider, err := decimal.NewFromString(args[1].Id)
	if err != nil {
		return ast.Expr{}, err
	}

	result = result.Div(divider)

	return ast.Expr{
		Id:   result.String(),
		Type: ast.Expr_Number,
	}, nil
}

func math_round(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}
	if err = checkIfExactArgs(args, 1); err != nil {
		return ast.Expr{}, err
	}

	if data[0].Type != ast.Expr_Number {
		return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
	}
	num, err := decimal.NewFromString(args[0].Id)
	if err != nil {
		return ast.Expr{}, err
	}

	rounded := num.Round(0)

	return ast.Expr{
		Id:   rounded.String(),
		Type: ast.Expr_Number,
	}, nil
}

func math_ceil(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}
	if err = checkIfExactArgs(args, 1); err != nil {
		return ast.Expr{}, err
	}

	if data[0].Type != ast.Expr_Number {
		return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
	}
	num, err := decimal.NewFromString(args[0].Id)
	if err != nil {
		return ast.Expr{}, err
	}

	rounded := num.RoundCeil(0)

	return ast.Expr{
		Id:   rounded.String(),
		Type: ast.Expr_Number,
	}, nil
}

func math_floor(data []*ast.Expr) (ast.Expr, error) {
	args, err := ParsedArgs(data)
	if err != nil {
		return ast.Expr{}, err
	}
	if err = checkIfExactArgs(args, 1); err != nil {
		return ast.Expr{}, err
	}

	if data[0].Type != ast.Expr_Number {
		return ast.Expr{}, fmt.Errorf("expected arguments to be of type integer or float")
	}
	num, err := decimal.NewFromString(args[0].Id)
	if err != nil {
		return ast.Expr{}, err
	}

	rounded := num.RoundFloor(0)

	return ast.Expr{
		Id:   rounded.String(),
		Type: ast.Expr_Number,
	}, nil
}
