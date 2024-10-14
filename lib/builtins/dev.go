package builtins

import (
	"github.com/Solarcode-org/Orion/lib"
	"github.com/Solarcode-org/Orion/lib/ast"
)

func EvalArgs(data ast.DataList) ast.DataList {
	args := make(ast.DataList, 0, len(data))

	for i := 0; i < len(data); i++ {
		datum := data[i]

		if datum.Type == ast.FuncCallType {
			value := lib.RunFunc(datum.Func, Functions)
			args = append(args, value)

			continue
		} else if datum.Type == ast.Ident {
			// TODO: add variable functionality
			args = append(args, ast.None)

			continue
		}

		args = append(args, datum)
	}

	return args
}
