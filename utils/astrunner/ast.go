package astrunner

import (
	"log"

	"github.com/Solarcode-org/Orion/ast"
	"github.com/Solarcode-org/Orion/lib/builtins"
	"github.com/Solarcode-org/Orion/utils"
	"github.com/shopspring/decimal"
)

// RunAST evaluates an Abstract Syntax Tree.
func RunAST(astree []*ast.Expr) {
	for i := 0; i < len(astree); i++ {
		stmt := astree[i]

		switch stmt.Type {
		case ast.Expr_FuncCall:
			_, err := utils.RunFunc(*stmt, builtins.Functions)

			utils.CheckErr(err)
		case ast.Expr_VariableDef:
			if stmt.Args[0].Type == ast.Expr_FuncCall {
				value, err := utils.RunFunc(*stmt.Args[0], builtins.Functions)
				utils.CheckErr(err)

				builtins.Variables[stmt.Id] = value
			} else if stmt.Args[0].Type == ast.Expr_Variable {
				builtins.Variables[stmt.Id] = builtins.Variables[stmt.Args[0].Id]
			} else {
				builtins.Variables[stmt.Id] = *stmt.Args[0]
			}
		case ast.Expr_VariableTypeDef:
			var val ast.Expr

			if stmt.Args[0].Type == ast.Expr_FuncCall {
				funcVal, err := utils.RunFunc(*stmt.Args[0], builtins.Functions)
				utils.CheckErr(err)

				val = funcVal
			} else if stmt.Args[0].Type == ast.Expr_Variable {
				val = builtins.Variables[stmt.Args[0].Id]
			} else {
				val = *stmt.Args[0]
			}

			switch stmt.Args[1].Type {
			case ast.Expr_String:
				builtins.Variables[stmt.Id] = ast.Expr{
					Type: ast.Expr_String,
					Id:   val.Id,
				}
			case ast.Expr_Number:
				if _, err := decimal.NewFromString(val.Id); err != nil {
					log.Fatalln(err)
				}

				builtins.Variables[stmt.Id] = ast.Expr{
					Type: ast.Expr_Number,
					Id:   val.Id,
				}
			}
		}
	}

}
