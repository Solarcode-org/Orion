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
	functions["Println"] = fmt_println
	functions["Print"] = fmt_print
	functions["Input"] = fmt_input
	functions["fmt/Println"] = fmt_println
	functions["fmt/Print"] = fmt_print
	functions["fmt/Input"] = fmt_input
	functions["fmt/Join"] = fmt_join
}

func fmt_print(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)

	for i := 0; i < len(args); i++ {
		datum := args[i]
		fmt.Print(datum.Data)

		fmt.Print(" ")
	}

	return ast.None, nil
}

func fmt_println(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)

	for i := 0; i < len(args); i++ {
		datum := args[i]
		fmt.Print(datum.Data)

		fmt.Print(" ")
	}

	fmt.Println()

	return ast.None, nil
}

func fmt_input(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)
	lib.ExactArgs("input", args, 1)

	if args[0].Type == ast.String {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(args[0].Data)

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

	return ast.None, fmt.Errorf("input: expected prompt to be of type string")
}

func fmt_join(data ast.DataList) (ast.Data, error) {
	args := EvalArgs(data)
	joined := ""

	for i := 0; i < len(data); i++ {
		datum := args[i]

		joined += datum.Data
	}

	return ast.Data{
		Data: joined,
		Type: ast.String,
	}, nil
}
