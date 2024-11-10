package lib

import (
	"strconv"

	"github.com/Solarcode-org/Orion/ast"

	"github.com/Solarcode-org/Orion/parser/bsr"
	"github.com/Solarcode-org/Orion/parser/symbols"
	log "github.com/sirupsen/logrus"
)

func buildRoot(b bsr.BSR) []*ast.Expr {
	nts := b.GetNTChildI(1).GetAllNTChildren()

	exprs := make([]*ast.Expr, 0, len(nts))
	for i := 0; i < len(nts); i++ {
		nt := nts[i][0]
		if nt.Label.Head() == symbols.NT_Statements {
			child := nt.GetNTChildI(0)

			if child.Label.Head() == symbols.NT_Statements && child.GetNTChildI(0).Label.Head() == symbols.NT_Statements {
				// log.Fatalln("todo: Statements = Statements Statement")
				statements := child.GetNTChildI(0).GetAllNTChildren()

				for j := 0; j < len(statements); j++ {
					statement := statements[j][0].GetNTChildI(0)

					expr := buildExpr(statement, nil)

					exprs = append(exprs, expr)
				}

				nt = child.GetNTChildI(1)
			} else if child.Label.Head() == symbols.NT_Statements {
				nt = child.GetNTChildI(0)
			} else {
				nt = child
			}
		}

		ntChild := nt.GetNTChildI(0)

		expr := buildExpr(ntChild, nil)

		exprs = append(exprs, expr)
	}

	return exprs
}

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

	case symbols.NT_Data:
		if len(b.GetAllNTChildren()) > 0 {
			child := b.GetNTChildI(0)

			return buildExpr(child, nil)

		} else {
			quoted := b.GetTChildI(0).LiteralString()
			str, err := strconv.Unquote(quoted)
			if err != nil {
				log.Fatalln(err)
			}

			return &ast.Expr{
				Type: ast.Expr_String,
				Id:   str,
			}
		}

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
		return &ast.Expr{}
	}
}
