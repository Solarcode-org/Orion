package bytecode

import (
	"encoding/binary"

	"github.com/Solarcode-org/Orion/ast"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
)

// DecodedSyntaxTree returns an Abstract Syntax Tree by decoding the
// bytecode instructions.
func DecodedSyntaxTree(instructions Instructions) ([]*ast.Expr, error) {
	astree := []*ast.Expr{}
	skip := 0

	for i := 0; i < len(instructions); i++ {
		if skip > 0 {
			skip--
			continue
		}

		var expr *ast.Expr
		var err error

		expr, skip, err = decodedExpr(instructions, i)
		if err != nil {
			return nil, err
		}
		if expr == nil {
			continue
		}

		astree = append(astree, expr)
	}

	return astree, nil
}

// decodedExpr converts instructions into expressions.
func decodedExpr(instructions Instructions, i int) (expr *ast.Expr, skip int, err error) {
	switch Instruction(instructions[i]) {
	case Func:
		funcNameLenEncoded := []byte{instructions[i+1], instructions[i+2]}
		funcNameLen := binary.BigEndian.Uint16(funcNameLenEncoded)

		i += 2

		funcName := instructions[i+1 : i+int(funcNameLen)+1]

		i += int(funcNameLen)

		funcArgsLenEncoded := []byte{instructions[i+1], instructions[i+2]}
		funcArgsLen := binary.BigEndian.Uint16(funcArgsLenEncoded)

		i += 3

		funcArgs := instructions[i : i+int(funcArgsLen)]

		args, err := DecodedSyntaxTree(funcArgs)
		if err != nil {
			return nil, 0, err
		}

		log.Tracef("Function %+#v with args %+#v", string(funcName), args)

		return &ast.Expr{
			Type: ast.Expr_FuncCall,
			Id:   string(funcName),
			Args: args,
		}, 4 + int(funcNameLen) + int(funcArgsLen), nil

	case String:
		stringLenEncoded := []byte{instructions[i+1], instructions[i+2]}
		stringLen := binary.BigEndian.Uint16(stringLenEncoded)

		i += 2

		stringVal := instructions[i+1 : i+int(stringLen)+1]

		log.Tracef("String %+#v", string(stringVal))

		return &ast.Expr{
			Type: ast.Expr_String,
			Id:   string(stringVal),
		}, 2 + int(stringLen), nil

	case Number:
		numLenEncoded := []byte{instructions[i+1], instructions[i+2]}
		numLen := binary.BigEndian.Uint16(numLenEncoded)

		i += 2

		rawNum := instructions[i+1 : i+int(numLen)+1]

		num := decimal.Zero
		if err := num.UnmarshalBinary(rawNum); err != nil {
			return nil, 0, err
		}

		log.Tracef("Number %+#v", num.String())

		return &ast.Expr{
			Type: ast.Expr_Number,
			Id:   num.String(),
		}, 2 + int(numLen), nil

	case Variable:
		varNameLenEncoded := []byte{instructions[i+1], instructions[i+2]}
		varNameLen := binary.BigEndian.Uint16(varNameLenEncoded)

		i += 2

		varName := instructions[i+1 : i+int(varNameLen)+1]

		log.Tracef("Variable %+#v", string(varName))

		return &ast.Expr{
			Type: ast.Expr_String,
			Id:   string(varName),
		}, 2 + int(varNameLen), nil

	case VariableDef:
		varNameLenEncoded := []byte{instructions[i+1], instructions[i+2]}
		varNameLen := binary.BigEndian.Uint16(varNameLenEncoded)

		i += 2

		varName := instructions[i+1 : i+int(varNameLen)+1]

		i += int(varNameLen)

		valLenEncoded := []byte{instructions[i+1], instructions[i+2]}
		valLen := binary.BigEndian.Uint16(valLenEncoded)

		i += 3

		rawVal := instructions[i : i+int(valLen)]
		val, err := DecodedSyntaxTree(rawVal)
		if err != nil {
			return nil, 0, err
		}

		log.Tracef("Variable definition (untyped) %+#v = %+#v", string(varName), val)

		return &ast.Expr{
			Type: ast.Expr_VariableDef,
			Id:   string(varName),
			Args: val,
		}, 4 + int(varNameLen) + int(valLen), nil

	case VariableTypeDef:
		varNameLenEncoded := []byte{instructions[i+1], instructions[i+2]}
		varNameLen := binary.BigEndian.Uint16(varNameLenEncoded)

		i += 2

		varName := instructions[i+1 : i+int(varNameLen)+1]

		i += int(varNameLen)

		varType := instructions[i+1]

		i++

		valLenEncoded := []byte{instructions[i+1], instructions[i+2]}
		valLen := binary.BigEndian.Uint16(valLenEncoded)

		i += 3

		rawVal := instructions[i : i+int(valLen)]
		val, err := DecodedSyntaxTree(rawVal)
		if err != nil {
			return nil, 0, err
		}

		args := []*ast.Expr{}
		args = append(args, val...)
		args = append(args, &ast.Expr{Type: ast.ExprType(varType)})

		log.Tracef("Variable definition (typed) %+#v: %+#v = %+#v", string(varName), varType, val)

		return &ast.Expr{
			Type: ast.Expr_VariableDef,
			Id:   string(varName),
			Args: args,
		}, 5 + int(varNameLen) + int(valLen), nil
	}

	return nil, 0, nil
}
