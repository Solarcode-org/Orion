package builtins

import (
	"fmt"
	"strings"

	"github.com/Solarcode-org/Orion/lib/ast"
)

func add_modgetter(functions FunctionsType) {
	functions["get"] = func(data ast.DataList) (ast.Data, error) {
		for i := 0; i < len(data); i++ {
			module := data[i]

			if module.Type != ast.String {
				return ast.None, fmt.Errorf("get: expected string or identifier as module arguments")
			}

			keys := make([]string, 0, len(functions))
			for k := range functions {
				keys = append(keys, k)
			}

			for _, key := range keys {
				if strings.Split(key, "/")[0] == module.Data {
					functions[strings.Split(key, "/")[1]] = functions[key]
				}
			}
		}

		return ast.None, nil
	}
}
