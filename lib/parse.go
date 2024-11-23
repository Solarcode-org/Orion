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
	"fmt"
	"strconv"

	"github.com/Solarcode-org/Orion/ast"
	"github.com/Solarcode-org/Orion/lexer"
	"github.com/Solarcode-org/Orion/lib/builtins"
	"github.com/Solarcode-org/Orion/parser"
	"github.com/Solarcode-org/Orion/utils"

	log "github.com/sirupsen/logrus"

	"github.com/Solarcode-org/Orion/parser/bsr"
	"github.com/Solarcode-org/Orion/parser/symbols"
)

// ParsedFrom uses the lexer to tokenize the input source and
// returns the AST formed by parsing the tokenized content.
func ParsedFrom(src []byte) ([]*ast.Expr, []*parser.Error, error) {
	log.Tracef("started function `lib.GetAbstractSyntaxTree` with argument `src`=`%s`\n", src)

	lex := lexer.New([]rune(string(src)))
	bsrSet, errs := parser.Parse(lex)

	if len(errs) > 0 {
		return nil, errs, nil
	}

	ast, err := AbstractSyntaxTree(bsrSet.GetRoot())
	if err != nil {
		return nil, nil, err
	}

	log.Traceln("successfully ended function `lib.GetAbstractSyntaxTree`")
	return ast, nil, nil
}

// AbstractSyntaxTree takes a non-terminal (`Orion`) as the `b` argument and builds it into
// an AST.
func AbstractSyntaxTree(b bsr.BSR) ([]*ast.Expr, error) {
	nts := b.GetNTChildI(1).GetAllNTChildren()

	exprs := make([]*ast.Expr, 0, len(nts))
	for i := 0; i < len(nts); i++ {
		nt := nts[i][0]
		statements := Statements(nt)

		for j := 0; j < len(statements); j++ {
			bsr := statements[j].GetNTChildI(0)

			expr, err := Expr(bsr, nil)
			if err != nil {
				return nil, err
			}

			exprs = append(exprs, expr)
		}
	}

	return exprs, nil
}

// Statements converts the `Statement` and `Statements` non-terminals into an array of statements
func Statements(stmt bsr.BSR) []bsr.BSR {
	if stmt.Label.Head() == symbols.NT_Statement {
		return []bsr.BSR{stmt}
	} else {
		statements := []bsr.BSR{}
		children := stmt.GetAllNTChildren()

		for i := 0; i < len(children); i++ {
			child := children[i][0]

			statements = append(statements, Statements(child)...)
		}

		return statements
	}
}

// Expr takes any expression and builds it into an AST Node ([*ast.Expr])
func Expr(b bsr.BSR, passed_args []*ast.Expr) (*ast.Expr, error) {
	switch b.Label.Head() {
	case symbols.NT_FuncCall:
		var rawArgs [][]bsr.BSR

		if b.Alternate() == 0 {
			rawArgs = b.GetNTChildI(2).GetAllNTChildren()
		}

		args := make([]*ast.Expr, 0, len(rawArgs))

		for i := 0; i < len(rawArgs); i++ {
			bsr := rawArgs[i][0]

			expr, err := Expr(bsr, args)
			if err != nil {
				return nil, err
			}

			if bsr.Label.Head() == symbols.NT_DataList {
				args = expr.Args
			} else {
				args = append(args, expr)
			}
		}

		return &ast.Expr{
			Type: ast.Expr_FuncCall,
			Id:   b.GetTChildI(0).LiteralString(),
			Args: args,
		}, nil

	case symbols.NT_Import:
		rawArgs := b.GetNTChildI(1).GetAllNTChildren()

		args := make([]*ast.Expr, 0, len(rawArgs))

		for i := 0; i < len(rawArgs); i++ {
			bsr := rawArgs[i][0]

			expr, err := Expr(bsr, args)
			if err != nil {
				return nil, err
			}

			if bsr.Label.Head() == symbols.NT_DataList {
				args = expr.Args
			} else {
				args = append(args, expr)
			}
		}

		return &ast.Expr{
			Type: ast.Expr_FuncCall,
			Id:   "get",
			Args: args,
		}, nil

	case symbols.NT_String:
		quoted := b.GetTChildI(0).LiteralString()
		str, err := strconv.Unquote(quoted)
		utils.CheckErr(err)

		return &ast.Expr{
			Type: ast.Expr_String,
			Id:   str,
		}, nil

	case symbols.NT_Number:
		number := b.GetTChildI(0).LiteralString()

		return &ast.Expr{
			Type: ast.Expr_Number,
			Id:   number,
		}, nil

	case symbols.NT_Operation:
		operation := b.GetNTChildI(0)

		num1, err := Expr(operation, nil)
		if err != nil {
			return nil, err
		}

		if operation.Label.Head() == symbols.NT_Number {
			return num1, nil
		}

		op := b.GetTChildI(1)

		num2, err := Expr(b.GetNTChildI(2), nil)
		if err != nil {
			return nil, err
		}

		switch op.LiteralString() {
		case "+":
			return &ast.Expr{
				Type: ast.Expr_FuncCall,
				Id:   "Sum",
				Args: []*ast.Expr{num1, num2},
			}, nil
		case "-":
			return &ast.Expr{
				Type: ast.Expr_FuncCall,
				Id:   "Difference",
				Args: []*ast.Expr{num1, num2},
			}, nil
		case "*":
			return &ast.Expr{
				Type: ast.Expr_FuncCall,
				Id:   "Product",
				Args: []*ast.Expr{num1, num2},
			}, nil
		case "/":
			return &ast.Expr{
				Type: ast.Expr_FuncCall,
				Id:   "Quotient",
				Args: []*ast.Expr{num1, num2},
			}, nil

		default:
			return nil, fmt.Errorf("reached invalid operation: %s", op.LiteralString())
		}

	case symbols.NT_Data:
		child := b.GetNTChildI(0)

		return Expr(child, nil)

	case symbols.NT_DataList:
		args := passed_args

		if b.Alternate() == 1 {
			datalist := b.GetNTChildI(0).GetAllNTChildren()
			data := b.GetNTChildI(1)

			for i := 0; i < len(datalist); i++ {
				datum := datalist[i][0]

				arg, err := Expr(datum, args)
				if err != nil {
					return nil, err
				}

				if datum.Label.Head() == symbols.NT_DataList {
					args = append(args, arg.Args...)
				} else {
					args = append(args, arg)
				}
			}

			arg, err := Expr(data, args)
			if err != nil {
				return nil, err
			}

			args = append(args, arg)

		} else {
			arg, err := Expr(b.GetNTChildrenI(0)[0], args)
			if err != nil {
				return nil, err
			}

			args = append(args, arg)
		}

		return &ast.Expr{
			Args: args,
		}, nil

	case symbols.NT_VariableDef:
		varName := b.GetTChildI(0).LiteralString()

		if b.Alternate() == 1 {
			value, err := Expr(b.GetNTChildI(4), nil)
			if err != nil {
				return nil, err
			}

			varType := b.GetTChildI(2).LiteralString()

			switch varType {
			case "string":
				builtins.Variables[varName] = ast.Expr{
					Type: ast.Expr_String,
					Id:   value.Id,
				}

			case "number":
				if _, err := strconv.Atoi(value.Id); err != nil {
					return nil, err
				}

				builtins.Variables[varName] = ast.Expr{
					Type: ast.Expr_Number,
					Id:   value.Id,
				}
			}

			return &ast.Expr{}, nil
		}

		value := b.GetNTChildI(2)

		valueParsed, err := Expr(value, nil)
		if err != nil {
			return nil, err
		}

		builtins.Variables[varName] = *valueParsed

		return &ast.Expr{}, nil

	case symbols.NT_Variable:
		varName := b.GetTChildI(0).LiteralString()

		if value, ok := builtins.Variables[varName]; ok {
			return &value, nil
		} else {
			return nil, fmt.Errorf("could not find name: %s", varName)
		}

	default:
		return nil, fmt.Errorf("reached invalid parse: %s", b.Label.Head().String())
	}
}
