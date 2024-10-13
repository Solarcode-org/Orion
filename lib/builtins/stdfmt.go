package builtins

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Solarcode-org/Orion/lib"
	"github.com/Solarcode-org/Orion/lib/ast"
)

func add_fmt_mod(functions FunctionsType) {
	functions["print"] = fmt_print
	functions["input"] = fmt_input
	functions["fmt/print"] = fmt_print
	functions["fmt/input"] = fmt_input
	functions["fmt/join"] = fmt_join
}

func fmt_print(data ast.DataList) (ast.Data, error) {
	for i := 0; i < len(data); i++ {
		datum := data[i]

		if datum.Type == ast.String {
			fmt.Print(datum.Data)
		} else if data[0].Type == ast.FuncCallType {
			value := lib.RunFunc(data[0].Func, Functions)

			fmt.Print(value.Data)
		}

		fmt.Print(" ")
	}

	fmt.Println()

	return ast.None, nil
}

func fmt_input(data ast.DataList) (ast.Data, error) {
	lib.ExactArgs("input", data, 1)

	if data[0].Type == ast.String {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(data[0].Data)

		ans, err := reader.ReadString('\n')
		if err != nil {
			return ast.None, err
		}

		ans = strings.TrimSpace(ans)

		return ast.Data{
			Data: ans,
			Type: ast.String,
		}, nil
	}

	if data[0].Type == ast.FuncCallType {
		value := lib.RunFunc(data[0].Func, Functions)

		if value.Type == ast.String {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print(value.Data)

			ans, err := reader.ReadString('\n')
			if err != nil {
				return ast.None, err
			}

			ans = strings.TrimSpace(ans)

			return ast.Data{
				Data: ans,
				Type: ast.String,
			}, nil
		}
	}

	return ast.None, fmt.Errorf("expected prompt to be of type string")
}

func fmt_join(data ast.DataList) (ast.Data, error) {
	joined := ""

	for i := 0; i < len(data); i++ {
		datum := data[i]

		if datum.Type == ast.String {
			joined += datum.Data
		} else if data[0].Type == ast.FuncCallType {
			value := lib.RunFunc(data[0].Func, Functions)

			joined += value.Data
		}
	}

	return ast.Data{
		Data: joined,
		Type: ast.String,
	}, nil
}
