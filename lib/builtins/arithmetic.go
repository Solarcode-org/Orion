package builtins

import (
	"fmt"

	"github.com/Solarcode-org/Orion/lib"
	"github.com/Solarcode-org/Orion/lib/ast"
	"github.com/shopspring/decimal"
)

func add_arithmetic_mod(functions FunctionsType) {
	functions["Sum"] = math_sum
	functions["Difference"] = math_difference
	functions["Product"] = math_product
	functions["Quotient"] = math_quotient
	functions["math/Sum"] = math_sum
	functions["math/Difference"] = math_difference
	functions["math/Product"] = math_product
	functions["math/Quotient"] = math_quotient
	functions["math/Round"] = math_round
	functions["math/Ceil"] = math_ceil
	functions["math/Floor"] = math_floor
}

func math_sum(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)
	result := decimal.Zero

	for i := 0; i < len(args); i++ {
		datum := args[i]

		if datum.Type != ast.Int && datum.Type != ast.Float {
			return ast.None, fmt.Errorf("sum: expected arguments to be of type integer or float")
		}

		num, err := decimal.NewFromString(datum.Data)
		lib.HandleFatal(err)

		result = result.Add(num)
	}

	return ast.Data{
		Data: fmt.Sprint(result),
		Type: ast.Int,
	}, nil
}

func math_difference(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)
	lib.ExactArgs("difference", args, 2)

	if data[0].Type != ast.Int && data[0].Type != ast.Float {
		return ast.None, fmt.Errorf("difference: expected arguments to be of type integer or float")
	}
	result, err := decimal.NewFromString(args[0].Data)
	lib.HandleFatal(err)

	if data[1].Type != ast.Int && data[1].Type != ast.Float {
		return ast.None, fmt.Errorf("difference: expected arguments to be of type integer or float")
	}
	subtrahend, err := decimal.NewFromString(args[1].Data)
	lib.HandleFatal(err)

	result = result.Sub(subtrahend)

	return ast.Data{
		Data: fmt.Sprint(result),
		Type: ast.Int,
	}, nil
}

func math_product(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)

	if args[0].Type != ast.Int && args[0].Type != ast.Float {
		return ast.None, fmt.Errorf("product: expected arguments to be of type integer or float")
	}
	result, err := decimal.NewFromString(args[0].Data)
	lib.HandleFatal(err)

	for i := 0; i < len(args)-1; i++ {
		datum := args[i+1]

		if datum.Type != ast.Int && datum.Type != ast.Float {
			return ast.None, fmt.Errorf("product: expected arguments to be of type integer or float")
		}

		num, err := decimal.NewFromString(datum.Data)
		lib.HandleFatal(err)

		result = result.Mul(num)
	}

	return ast.Data{
		Data: fmt.Sprint(result),
		Type: ast.Int,
	}, nil
}

func math_quotient(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)
	lib.ExactArgs("quotient", args, 2)

	if data[0].Type != ast.Int && data[0].Type != ast.Float {
		return ast.None, fmt.Errorf("quotient: expected arguments to be of type integer or float")
	}
	result, err := decimal.NewFromString(args[0].Data)
	lib.HandleFatal(err)

	if data[1].Type != ast.Int && data[1].Type != ast.Float {
		return ast.None, fmt.Errorf("quotient: expected arguments to be of type integer or float")
	}
	divider, err := decimal.NewFromString(args[1].Data)
	lib.HandleFatal(err)

	result = result.Div(divider)

	return ast.Data{
		Data: fmt.Sprint(result),
		Type: ast.Int,
	}, nil
}

func math_round(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)
	lib.ExactArgs("math/round", args, 1)

	if data[0].Type != ast.Int && data[0].Type != ast.Float {
		return ast.None, fmt.Errorf("math/round: expected arguments to be of type integer or float")
	}

	return ast.Data{
		Data: "0",
		Type: ast.Int,
	}, nil
}

func math_ceil(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)
	lib.ExactArgs("math/ceil", args, 1)

	if data[0].Type != ast.Int && data[0].Type != ast.Float {
		return ast.None, fmt.Errorf("math/ceil: expected arguments to be of type integer or float")
	}

	return ast.Data{
		Data: "0",
		Type: ast.Int,
	}, nil
}

func math_floor(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)
	lib.ExactArgs("math/floor", args, 1)

	if data[0].Type != ast.Int && data[0].Type != ast.Float {
		return ast.None, fmt.Errorf("math/floor: expected arguments to be of type integer or float")
	}

	return ast.Data{
		Data: "0",
		Type: ast.Int,
	}, nil
}
