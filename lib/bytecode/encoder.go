package bytecode

import (
	"encoding/binary"

	"github.com/Solarcode-org/Orion/ast"
	"github.com/shopspring/decimal"
)

// EncodedSyntaxTree returns bytecode instructions for the given
// Abstract Syntax Tree.
func EncodedSyntaxTree(astree []*ast.Expr) (Instructions, error) {
	bytecode := Instructions{}

	for i := 0; i < len(astree); i++ {
		expr := astree[i]

		switch expr.Type {
		case ast.Expr_FuncCall:
			bytecode = append(bytecode, byte(Func))

			bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(len(expr.Id)))
			bytecode = append(bytecode, Instructions(expr.Id)...)

			args, err := EncodedSyntaxTree(expr.Args)
			if err != nil {
				return nil, err
			}

			bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(len(args)))
			bytecode = append(bytecode, args...)

		case ast.Expr_String:
			bytecode = append(bytecode, byte(String))
			bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(len(expr.Id)))
			bytecode = append(bytecode, Instructions(expr.Id)...)

		case ast.Expr_Number:
			bytecode = append(bytecode, byte(Number))

			num, err := decimal.NewFromString(expr.Id)
			if err != nil {
				return nil, err
			}

			binaryNum, err := num.MarshalBinary()
			if err != nil {
				return nil, err
			}

			bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(len(binaryNum)))
			bytecode = append(bytecode, binaryNum...)

		case ast.Expr_Variable:
			bytecode = append(bytecode, byte(Variable))
			bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(len(expr.Id)))
			bytecode = append(bytecode, Instructions(expr.Id)...)

		case ast.Expr_VariableDef:
			bytecode = append(bytecode, byte(VariableDef))

			bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(len(expr.Id)))
			bytecode = append(bytecode, Instructions(expr.Id)...)

			rawVal := expr.Args[0]
			val, err := EncodedSyntaxTree([]*ast.Expr{rawVal})
			if err != nil {
				return nil, err
			}

			bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(len(val)))
			bytecode = append(bytecode, val...)

		case ast.Expr_VariableTypeDef:
			bytecode = append(bytecode, byte(VariableTypeDef))

			bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(len(expr.Id)))
			bytecode = append(bytecode, Instructions(expr.Id)...)

			varType := expr.Args[1].Type
			bytecode = append(bytecode, byte(varType))

			rawVal := expr.Args[0]
			val, err := EncodedSyntaxTree([]*ast.Expr{rawVal})
			if err != nil {
				return nil, err
			}

			bytecode = binary.BigEndian.AppendUint16(bytecode, uint16(len(val)))
			bytecode = append(bytecode, val...)

		default:
			bytecode = append(bytecode, byte(None))
		}
	}

	return bytecode, nil
}
