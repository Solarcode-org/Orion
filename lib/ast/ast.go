package ast

import (
	"github.com/Solarcode-org/Orion/lib/token"
)

type (
	FuncCallList []FuncCall

	DataList []Data
	DataType uint32
)

type FuncCall struct {
	Name string
	Args DataList
}

type Data struct {
	Data string
	Type DataType
	Func FuncCall
}

const (
	Int DataType = iota
	Float
	String
	NoneType
	FuncCallType
	Ident
)

var None = Data{
	Data: "None",
	Type: NoneType,
}

func NewFuncCallList(funcCall interface{}) (FuncCallList, error) {
	return FuncCallList{funcCall.(FuncCall)}, nil
}

func AppendFuncCall(funcCallList, funcCall interface{}) (FuncCallList, error) {
	return append(funcCallList.(FuncCallList), funcCall.(FuncCall)), nil
}

// funcName is passed using $T0, funcArgs is passed using $2.
func NewFuncCall(funcName *token.Token, funcArgs interface{}) (FuncCall, error) {
	return FuncCall{
		Name: string(funcName.Lit),
		Args: funcArgs.(DataList),
	}, nil
}

func NewFuncCallOneArg(funcName *token.Token, funcArgs interface{}) (FuncCall, error) {
	return FuncCall{
		Name: string(funcName.Lit),
		Args: []Data{funcArgs.(Data)},
	}, nil
}

func NewFuncCallManyArgs(funcName *token.Token, funcArgs ...interface{}) (FuncCall, error) {
	args := make(DataList, 0, len(funcArgs))

	for _, arg := range funcArgs {
		args = append(args, arg.(Data))
	}

	return FuncCall{
		Name: string(funcName.Lit),
		Args: args,
	}, nil
}

func NewDataList(data interface{}) (DataList, error) {
	return DataList{data.(Data)}, nil
}

func AppendData(dataList, data interface{}) (DataList, error) {
	return append(dataList.(DataList), data.(Data)), nil
}

// data is passed using $T0, dataType is passed using $1.
func NewData(data *token.Token, dataType DataType) (Data, error) {
	dataStr := string(data.Lit)

	if dataType == String {
		return Data{
			Data: dataStr[1 : len(dataStr)-1],
			Type: String,
		}, nil
	}

	return Data{
		Data: dataStr,
		Type: dataType,
	}, nil
}

// data is passed using $0, dataType is passed using $1.
func NewFuncData(data interface{}) (Data, error) {
	return Data{
		Data: "",
		Type: FuncCallType,
		Func: data.(FuncCall),
	}, nil
}
