
// Package slot is generated by gogll. Do not edit. 
package slot

import(
	"bytes"
	"fmt"
	
	"github.com/Solarcode-org/Orion/parser/symbols"
)

type Label int

const(
	Data0R0 Label = iota
	Data0R1
	Data1R0
	Data1R1
	Data2R0
	Data2R1
	Data3R0
	Data3R1
	DataList0R0
	DataList0R1
	DataList1R0
	DataList1R1
	DataList1R2
	DataList1R3
	FuncCall0R0
	FuncCall0R1
	FuncCall0R2
	FuncCall0R3
	FuncCall0R4
	FuncCall1R0
	FuncCall1R1
	FuncCall1R2
	FuncCall1R3
	Import0R0
	Import0R1
	Import0R2
	Number0R0
	Number0R1
	Number1R0
	Number1R1
	Operation0R0
	Operation0R1
	Operation1R0
	Operation1R1
	Operation1R2
	Operation1R3
	Orion0R0
	Orion0R1
	Orion0R2
	Package0R0
	Package0R1
	Package0R2
	Statement0R0
	Statement0R1
	Statement1R0
	Statement1R1
	Statements0R0
	Statements0R1
	Statements1R0
	Statements1R1
	Statements1R2
	String0R0
	String0R1
)

type Slot struct {
	NT      symbols.NT
	Alt     int
	Pos     int
	Symbols symbols.Symbols
	Label 	Label
}

type Index struct {
	NT      symbols.NT
	Alt     int
	Pos     int
}

func GetAlternates(nt symbols.NT) []Label {
	alts, exist := alternates[nt]
	if !exist {
		panic(fmt.Sprintf("Invalid NT %s", nt))
	}
	return alts
}

func GetLabel(nt symbols.NT, alt, pos int) Label {
	l, exist := slotIndex[Index{nt,alt,pos}]
	if exist {
		return l
	}
	panic(fmt.Sprintf("Error: no slot label for NT=%s, alt=%d, pos=%d", nt, alt, pos))
}

func (l Label) EoR() bool {
	return l.Slot().EoR()
}

func (l Label) Head() symbols.NT {
	return l.Slot().NT
}

func (l Label) Index() Index {
	s := l.Slot()
	return Index{s.NT, s.Alt, s.Pos}
}

func (l Label) Alternate() int {
	return l.Slot().Alt
}

func (l Label) Pos() int {
	return l.Slot().Pos
}

func (l Label) Slot() *Slot {
	s, exist := slots[l]
	if !exist {
		panic(fmt.Sprintf("Invalid slot label %d", l))
	}
	return s
}

func (l Label) String() string {
	return l.Slot().String()
}

func (l Label) Symbols() symbols.Symbols {
	return l.Slot().Symbols
}

func (s *Slot) EoR() bool {
	return s.Pos >= len(s.Symbols)
}

func (s *Slot) String() string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%s : ", s.NT)
	for i, sym := range s.Symbols {
		if i == s.Pos {
			fmt.Fprintf(buf, "∙")
		}
		fmt.Fprintf(buf, "%s ", sym)
	}
	if s.Pos >= len(s.Symbols) {
		fmt.Fprintf(buf, "∙")
	}
	return buf.String()
}

var slots = map[Label]*Slot{ 
	Data0R0: {
		symbols.NT_Data, 0, 0, 
		symbols.Symbols{  
			symbols.NT_String,
		}, 
		Data0R0, 
	},
	Data0R1: {
		symbols.NT_Data, 0, 1, 
		symbols.Symbols{  
			symbols.NT_String,
		}, 
		Data0R1, 
	},
	Data1R0: {
		symbols.NT_Data, 1, 0, 
		symbols.Symbols{  
			symbols.NT_FuncCall,
		}, 
		Data1R0, 
	},
	Data1R1: {
		symbols.NT_Data, 1, 1, 
		symbols.Symbols{  
			symbols.NT_FuncCall,
		}, 
		Data1R1, 
	},
	Data2R0: {
		symbols.NT_Data, 2, 0, 
		symbols.Symbols{  
			symbols.NT_Number,
		}, 
		Data2R0, 
	},
	Data2R1: {
		symbols.NT_Data, 2, 1, 
		symbols.Symbols{  
			symbols.NT_Number,
		}, 
		Data2R1, 
	},
	Data3R0: {
		symbols.NT_Data, 3, 0, 
		symbols.Symbols{  
			symbols.NT_Operation,
		}, 
		Data3R0, 
	},
	Data3R1: {
		symbols.NT_Data, 3, 1, 
		symbols.Symbols{  
			symbols.NT_Operation,
		}, 
		Data3R1, 
	},
	DataList0R0: {
		symbols.NT_DataList, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Data,
		}, 
		DataList0R0, 
	},
	DataList0R1: {
		symbols.NT_DataList, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Data,
		}, 
		DataList0R1, 
	},
	DataList1R0: {
		symbols.NT_DataList, 1, 0, 
		symbols.Symbols{  
			symbols.NT_DataList, 
			symbols.T_2, 
			symbols.NT_Data,
		}, 
		DataList1R0, 
	},
	DataList1R1: {
		symbols.NT_DataList, 1, 1, 
		symbols.Symbols{  
			symbols.NT_DataList, 
			symbols.T_2, 
			symbols.NT_Data,
		}, 
		DataList1R1, 
	},
	DataList1R2: {
		symbols.NT_DataList, 1, 2, 
		symbols.Symbols{  
			symbols.NT_DataList, 
			symbols.T_2, 
			symbols.NT_Data,
		}, 
		DataList1R2, 
	},
	DataList1R3: {
		symbols.NT_DataList, 1, 3, 
		symbols.Symbols{  
			symbols.NT_DataList, 
			symbols.T_2, 
			symbols.NT_Data,
		}, 
		DataList1R3, 
	},
	FuncCall0R0: {
		symbols.NT_FuncCall, 0, 0, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.T_0, 
			symbols.NT_DataList, 
			symbols.T_1,
		}, 
		FuncCall0R0, 
	},
	FuncCall0R1: {
		symbols.NT_FuncCall, 0, 1, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.T_0, 
			symbols.NT_DataList, 
			symbols.T_1,
		}, 
		FuncCall0R1, 
	},
	FuncCall0R2: {
		symbols.NT_FuncCall, 0, 2, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.T_0, 
			symbols.NT_DataList, 
			symbols.T_1,
		}, 
		FuncCall0R2, 
	},
	FuncCall0R3: {
		symbols.NT_FuncCall, 0, 3, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.T_0, 
			symbols.NT_DataList, 
			symbols.T_1,
		}, 
		FuncCall0R3, 
	},
	FuncCall0R4: {
		symbols.NT_FuncCall, 0, 4, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.T_0, 
			symbols.NT_DataList, 
			symbols.T_1,
		}, 
		FuncCall0R4, 
	},
	FuncCall1R0: {
		symbols.NT_FuncCall, 1, 0, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.T_0, 
			symbols.T_1,
		}, 
		FuncCall1R0, 
	},
	FuncCall1R1: {
		symbols.NT_FuncCall, 1, 1, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.T_0, 
			symbols.T_1,
		}, 
		FuncCall1R1, 
	},
	FuncCall1R2: {
		symbols.NT_FuncCall, 1, 2, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.T_0, 
			symbols.T_1,
		}, 
		FuncCall1R2, 
	},
	FuncCall1R3: {
		symbols.NT_FuncCall, 1, 3, 
		symbols.Symbols{  
			symbols.T_6, 
			symbols.T_0, 
			symbols.T_1,
		}, 
		FuncCall1R3, 
	},
	Import0R0: {
		symbols.NT_Import, 0, 0, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_DataList,
		}, 
		Import0R0, 
	},
	Import0R1: {
		symbols.NT_Import, 0, 1, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_DataList,
		}, 
		Import0R1, 
	},
	Import0R2: {
		symbols.NT_Import, 0, 2, 
		symbols.Symbols{  
			symbols.T_5, 
			symbols.NT_DataList,
		}, 
		Import0R2, 
	},
	Number0R0: {
		symbols.NT_Number, 0, 0, 
		symbols.Symbols{  
			symbols.T_7,
		}, 
		Number0R0, 
	},
	Number0R1: {
		symbols.NT_Number, 0, 1, 
		symbols.Symbols{  
			symbols.T_7,
		}, 
		Number0R1, 
	},
	Number1R0: {
		symbols.NT_Number, 1, 0, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		Number1R0, 
	},
	Number1R1: {
		symbols.NT_Number, 1, 1, 
		symbols.Symbols{  
			symbols.T_4,
		}, 
		Number1R1, 
	},
	Operation0R0: {
		symbols.NT_Operation, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Number,
		}, 
		Operation0R0, 
	},
	Operation0R1: {
		symbols.NT_Operation, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Number,
		}, 
		Operation0R1, 
	},
	Operation1R0: {
		symbols.NT_Operation, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Operation, 
			symbols.T_9, 
			symbols.NT_Number,
		}, 
		Operation1R0, 
	},
	Operation1R1: {
		symbols.NT_Operation, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Operation, 
			symbols.T_9, 
			symbols.NT_Number,
		}, 
		Operation1R1, 
	},
	Operation1R2: {
		symbols.NT_Operation, 1, 2, 
		symbols.Symbols{  
			symbols.NT_Operation, 
			symbols.T_9, 
			symbols.NT_Number,
		}, 
		Operation1R2, 
	},
	Operation1R3: {
		symbols.NT_Operation, 1, 3, 
		symbols.Symbols{  
			symbols.NT_Operation, 
			symbols.T_9, 
			symbols.NT_Number,
		}, 
		Operation1R3, 
	},
	Orion0R0: {
		symbols.NT_Orion, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Package, 
			symbols.NT_Statements,
		}, 
		Orion0R0, 
	},
	Orion0R1: {
		symbols.NT_Orion, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Package, 
			symbols.NT_Statements,
		}, 
		Orion0R1, 
	},
	Orion0R2: {
		symbols.NT_Orion, 0, 2, 
		symbols.Symbols{  
			symbols.NT_Package, 
			symbols.NT_Statements,
		}, 
		Orion0R2, 
	},
	Package0R0: {
		symbols.NT_Package, 0, 0, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_11,
		}, 
		Package0R0, 
	},
	Package0R1: {
		symbols.NT_Package, 0, 1, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_11,
		}, 
		Package0R1, 
	},
	Package0R2: {
		symbols.NT_Package, 0, 2, 
		symbols.Symbols{  
			symbols.T_10, 
			symbols.T_11,
		}, 
		Package0R2, 
	},
	Statement0R0: {
		symbols.NT_Statement, 0, 0, 
		symbols.Symbols{  
			symbols.NT_FuncCall,
		}, 
		Statement0R0, 
	},
	Statement0R1: {
		symbols.NT_Statement, 0, 1, 
		symbols.Symbols{  
			symbols.NT_FuncCall,
		}, 
		Statement0R1, 
	},
	Statement1R0: {
		symbols.NT_Statement, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Import,
		}, 
		Statement1R0, 
	},
	Statement1R1: {
		symbols.NT_Statement, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Import,
		}, 
		Statement1R1, 
	},
	Statements0R0: {
		symbols.NT_Statements, 0, 0, 
		symbols.Symbols{  
			symbols.NT_Statement,
		}, 
		Statements0R0, 
	},
	Statements0R1: {
		symbols.NT_Statements, 0, 1, 
		symbols.Symbols{  
			symbols.NT_Statement,
		}, 
		Statements0R1, 
	},
	Statements1R0: {
		symbols.NT_Statements, 1, 0, 
		symbols.Symbols{  
			symbols.NT_Statements, 
			symbols.NT_Statement,
		}, 
		Statements1R0, 
	},
	Statements1R1: {
		symbols.NT_Statements, 1, 1, 
		symbols.Symbols{  
			symbols.NT_Statements, 
			symbols.NT_Statement,
		}, 
		Statements1R1, 
	},
	Statements1R2: {
		symbols.NT_Statements, 1, 2, 
		symbols.Symbols{  
			symbols.NT_Statements, 
			symbols.NT_Statement,
		}, 
		Statements1R2, 
	},
	String0R0: {
		symbols.NT_String, 0, 0, 
		symbols.Symbols{  
			symbols.T_11,
		}, 
		String0R0, 
	},
	String0R1: {
		symbols.NT_String, 0, 1, 
		symbols.Symbols{  
			symbols.T_11,
		}, 
		String0R1, 
	},
}

var slotIndex = map[Index]Label { 
	Index{ symbols.NT_Data,0,0 }: Data0R0,
	Index{ symbols.NT_Data,0,1 }: Data0R1,
	Index{ symbols.NT_Data,1,0 }: Data1R0,
	Index{ symbols.NT_Data,1,1 }: Data1R1,
	Index{ symbols.NT_Data,2,0 }: Data2R0,
	Index{ symbols.NT_Data,2,1 }: Data2R1,
	Index{ symbols.NT_Data,3,0 }: Data3R0,
	Index{ symbols.NT_Data,3,1 }: Data3R1,
	Index{ symbols.NT_DataList,0,0 }: DataList0R0,
	Index{ symbols.NT_DataList,0,1 }: DataList0R1,
	Index{ symbols.NT_DataList,1,0 }: DataList1R0,
	Index{ symbols.NT_DataList,1,1 }: DataList1R1,
	Index{ symbols.NT_DataList,1,2 }: DataList1R2,
	Index{ symbols.NT_DataList,1,3 }: DataList1R3,
	Index{ symbols.NT_FuncCall,0,0 }: FuncCall0R0,
	Index{ symbols.NT_FuncCall,0,1 }: FuncCall0R1,
	Index{ symbols.NT_FuncCall,0,2 }: FuncCall0R2,
	Index{ symbols.NT_FuncCall,0,3 }: FuncCall0R3,
	Index{ symbols.NT_FuncCall,0,4 }: FuncCall0R4,
	Index{ symbols.NT_FuncCall,1,0 }: FuncCall1R0,
	Index{ symbols.NT_FuncCall,1,1 }: FuncCall1R1,
	Index{ symbols.NT_FuncCall,1,2 }: FuncCall1R2,
	Index{ symbols.NT_FuncCall,1,3 }: FuncCall1R3,
	Index{ symbols.NT_Import,0,0 }: Import0R0,
	Index{ symbols.NT_Import,0,1 }: Import0R1,
	Index{ symbols.NT_Import,0,2 }: Import0R2,
	Index{ symbols.NT_Number,0,0 }: Number0R0,
	Index{ symbols.NT_Number,0,1 }: Number0R1,
	Index{ symbols.NT_Number,1,0 }: Number1R0,
	Index{ symbols.NT_Number,1,1 }: Number1R1,
	Index{ symbols.NT_Operation,0,0 }: Operation0R0,
	Index{ symbols.NT_Operation,0,1 }: Operation0R1,
	Index{ symbols.NT_Operation,1,0 }: Operation1R0,
	Index{ symbols.NT_Operation,1,1 }: Operation1R1,
	Index{ symbols.NT_Operation,1,2 }: Operation1R2,
	Index{ symbols.NT_Operation,1,3 }: Operation1R3,
	Index{ symbols.NT_Orion,0,0 }: Orion0R0,
	Index{ symbols.NT_Orion,0,1 }: Orion0R1,
	Index{ symbols.NT_Orion,0,2 }: Orion0R2,
	Index{ symbols.NT_Package,0,0 }: Package0R0,
	Index{ symbols.NT_Package,0,1 }: Package0R1,
	Index{ symbols.NT_Package,0,2 }: Package0R2,
	Index{ symbols.NT_Statement,0,0 }: Statement0R0,
	Index{ symbols.NT_Statement,0,1 }: Statement0R1,
	Index{ symbols.NT_Statement,1,0 }: Statement1R0,
	Index{ symbols.NT_Statement,1,1 }: Statement1R1,
	Index{ symbols.NT_Statements,0,0 }: Statements0R0,
	Index{ symbols.NT_Statements,0,1 }: Statements0R1,
	Index{ symbols.NT_Statements,1,0 }: Statements1R0,
	Index{ symbols.NT_Statements,1,1 }: Statements1R1,
	Index{ symbols.NT_Statements,1,2 }: Statements1R2,
	Index{ symbols.NT_String,0,0 }: String0R0,
	Index{ symbols.NT_String,0,1 }: String0R1,
}

var alternates = map[symbols.NT][]Label{ 
	symbols.NT_Orion:[]Label{ Orion0R0 },
	symbols.NT_Package:[]Label{ Package0R0 },
	symbols.NT_Statements:[]Label{ Statements0R0,Statements1R0 },
	symbols.NT_Statement:[]Label{ Statement0R0,Statement1R0 },
	symbols.NT_FuncCall:[]Label{ FuncCall0R0,FuncCall1R0 },
	symbols.NT_Import:[]Label{ Import0R0 },
	symbols.NT_DataList:[]Label{ DataList0R0,DataList1R0 },
	symbols.NT_Data:[]Label{ Data0R0,Data1R0,Data2R0,Data3R0 },
	symbols.NT_String:[]Label{ String0R0 },
	symbols.NT_Number:[]Label{ Number0R0,Number1R0 },
	symbols.NT_Operation:[]Label{ Operation0R0,Operation1R0 },
}

