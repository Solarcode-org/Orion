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
