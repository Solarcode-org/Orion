package bytecode

// The Instructions encoded from an Abstract Syntax Tree.
type Instructions []byte

// An Instruction is a singular byte used to represent any
// bytecode instruction.
//
// Multiple of these can be used to encode numbers, strings
// and other data types
type Instruction byte

// Bytecode Instructions.
const (
	// Func marks that the instructions for a function call will
	// follow it.
	//
	// Instructions for a function are in the form:
	//
	// [Instruction Func]
	// (length of function name) (encoded function name)
	// (length of encoded args) (encoded args)
	Func Instruction = iota

	// String marks that the instructions for a string will
	// follow it.
	//
	// Instructions for a string are in the form:
	//
	// [Instruction String]
	// (length of string) (encoded string)
	String

	// Number marks that the instructions for a number will
	// follow it.
	//
	// As numbers in Orion are stored as Decimal objects for
	// precision, it is stored as a binary string.
	//
	// Instructions for a number are in the form:
	//
	// [Instruction Number]
	// (length of decimal number) (encoded decimal number)
	Number

	// Variable marks that the instructions for a variable will
	// follow it.
	//
	// Instructions for a variable are in the form:
	//
	// [Instruction Variable]
	// (length of variable name) (encoded variable name)
	Variable

	// VariableDef marks that the instructions for an untyped
	// variable definition will follow it.
	//
	// Instructions for an untyped variable definition are in
	// the form:
	//
	// [Instruction Variable]
	// (length of variable name) (encoded variable name)
	// (length of encoded variable value) (encoded variable value)
	VariableDef

	// VariableDef marks that the instructions for a typed
	// variable definition will follow it.
	//
	// Instructions for a typed variable definition are in
	// the form:
	//
	// [Instruction Variable]
	// (encoded variable type {which is only one byte})
	// (length of variable name) (encoded variable name)
	// (length of encoded variable value) (encoded variable value)
	VariableTypeDef

	// The nil instruction.
	//
	// It has no particular use.
	None
)
