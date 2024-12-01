package bytecode

type Instructions []byte
type Instruction byte

const (
	Func Instruction = iota
	String
	Number
	None
	Variable
	VariableDef
	VariableTypeDef
)
