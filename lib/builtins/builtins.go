package builtins

import "github.com/Solarcode-org/Orion/lib/ast"

type FunctionsType = map[string]OrionFunction
type OrionFunction = func(ast.DataList) (ast.Data, error)

var Functions FunctionsType

func MakeFunctions() {
	Functions = make(FunctionsType)

	add_fmt_mod(Functions)
	add_modgetter(Functions)
}
