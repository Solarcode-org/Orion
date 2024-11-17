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

package lib

import (
	"strconv"

	"github.com/Solarcode-org/Orion/ast"

	"github.com/Solarcode-org/Orion/parser/bsr"
	"github.com/Solarcode-org/Orion/parser/symbols"
	log "github.com/sirupsen/logrus"
)

// buildAST takes a non-terminal (`Orion`) as the `b` argument and builds it into
// an Abstract Syntax tree.
func buildAST(b bsr.BSR) []*ast.Expr {
	nts := b.GetNTChildI(1).GetAllNTChildren()

	exprs := make([]*ast.Expr, 0, len(nts))
	for i := 0; i < len(nts); i++ {
		nt := nts[i][0]
		statements := getStatements(nt)

		for j := 0; j < len(statements); j++ {
			bsr := statements[j].GetNTChildI(0)

			expr := buildExpr(bsr, nil)

			exprs = append(exprs, expr)
		}
	}

	return exprs
}

// getStatements converts the `Statement` and `Statements` non-terminals into an array of statements
func getStatements(stmt bsr.BSR) []bsr.BSR {
	if stmt.Label.Head() == symbols.NT_Statement {
		return []bsr.BSR{stmt}
	} else {
		statements := []bsr.BSR{}
		children := stmt.GetAllNTChildren()

		for i := 0; i < len(children); i++ {
			child := children[i][0]

			statements = append(statements, getStatements(child)...)
		}

		return statements
	}
}

// buildExpr takes any expression and builds it into an AST Node ([*ast.Expr])
func buildExpr(b bsr.BSR, passed_args []*ast.Expr) *ast.Expr {
	switch b.Label.Head() {
	case symbols.NT_FuncCall:
		rawArgs := b.GetNTChildrenI(2)[0].GetAllNTChildren()

		args := make([]*ast.Expr, 0, len(rawArgs))

		for i := 0; i < len(rawArgs); i++ {
			bsr := rawArgs[i][0]

			built := buildExpr(bsr, args)

			if bsr.Label.Head() == symbols.NT_DataList {
				args = built.Args
			} else {
				args = append(args, built)
			}
		}

		return &ast.Expr{
			Type: ast.Expr_FuncCall,
			Id:   b.GetTChildI(0).LiteralString(),
			Args: args,
		}

	case symbols.NT_Import:
		rawArgs := b.GetNTChildrenI(1)[0].GetAllNTChildren()

		args := make([]*ast.Expr, 0, len(rawArgs))

		for i := 0; i < len(rawArgs); i++ {
			bsr := rawArgs[i][0]

			built := buildExpr(bsr, args)

			if bsr.Label.Head() == symbols.NT_DataList {
				args = built.Args
			} else {
				args = append(args, built)
			}
		}

		return &ast.Expr{
			Type: ast.Expr_FuncCall,
			Id:   "get",
			Args: args,
		}

	case symbols.NT_String:
		quoted := b.GetTChildI(0).LiteralString()
		str, err := strconv.Unquote(quoted)
		CheckErr(err)

		return &ast.Expr{
			Type: ast.Expr_String,
			Id:   str,
		}

	case symbols.NT_Number:
		number := b.GetTChildI(0).LiteralString()

		return &ast.Expr{
			Type: ast.Expr_Number,
			Id:   number,
		}

	case symbols.NT_Operation:
		operation := b.GetNTChildI(0)
		num1 := buildExpr(operation, nil)

		if operation.Label.Head() == symbols.NT_Number {
			return num1
		}

		op := b.GetTChildI(1)
		num2 := buildExpr(b.GetNTChildI(2), nil)

		switch op.LiteralString() {
		case "+":
			return &ast.Expr{
				Type: ast.Expr_FuncCall,
				Id:   "Sum",
				Args: []*ast.Expr{num1, num2},
			}
		case "-":
			return &ast.Expr{
				Type: ast.Expr_FuncCall,
				Id:   "Difference",
				Args: []*ast.Expr{num1, num2},
			}
		case "*":
			return &ast.Expr{
				Type: ast.Expr_FuncCall,
				Id:   "Product",
				Args: []*ast.Expr{num1, num2},
			}
		case "/":
			return &ast.Expr{
				Type: ast.Expr_FuncCall,
				Id:   "Quotient",
				Args: []*ast.Expr{num1, num2},
			}

		default:
			log.Fatalln("reached invalid operation", op.LiteralString())
			return &ast.Expr{}
		}

	case symbols.NT_Data:
		child := b.GetNTChildI(0)

		return buildExpr(child, nil)

	case symbols.NT_DataList:
		children := b.GetAllNTChildren()

		args := passed_args

		if len(children) == 2 {
			datalist := children[0][0].GetAllNTChildren()
			data := children[1][0]

			for i := 0; i < len(datalist); i++ {
				datum := datalist[i][0]

				arg := buildExpr(datum, args)

				if datum.Label.Head() == symbols.NT_DataList {
					args = append(args, arg.Args...)
				} else {
					args = append(args, arg)
				}
			}

			args = append(args, buildExpr(
				data, args))

		} else {
			args = append(args, buildExpr(children[0][0], args))
		}

		return &ast.Expr{
			Args: args,
		}

	default:
		log.Fatalln("reached invalid parse", b.Label.Head())
		return nil
	}
}
