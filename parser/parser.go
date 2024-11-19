// Package parser is generated by gogll. Do not edit.
package parser

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/Solarcode-org/Orion/lexer"
	"github.com/Solarcode-org/Orion/parser/bsr"
	"github.com/Solarcode-org/Orion/parser/slot"
	"github.com/Solarcode-org/Orion/parser/symbols"
	"github.com/Solarcode-org/Orion/token"
)

type parser struct {
	cI int

	R *descriptors
	U *descriptors

	popped   map[poppedNode]bool
	crf      map[clusterNode][]*crfNode
	crfNodes map[crfNode]*crfNode

	lex         *lexer.Lexer
	parseErrors []*Error

	bsrSet *bsr.Set
}

func newParser(l *lexer.Lexer) *parser {
	return &parser{
		cI:     0,
		lex:    l,
		R:      &descriptors{},
		U:      &descriptors{},
		popped: make(map[poppedNode]bool),
		crf: map[clusterNode][]*crfNode{
			{symbols.NT_Orion, 0}: {},
		},
		crfNodes:    map[crfNode]*crfNode{},
		bsrSet:      bsr.New(symbols.NT_Orion, l),
		parseErrors: nil,
	}
}

// Parse returns the BSR set containing the parse forest.
// If the parse was successfull []*Error is nil
func Parse(l *lexer.Lexer) (*bsr.Set, []*Error) {
	return newParser(l).parse()
}

func (p *parser) parse() (*bsr.Set, []*Error) {
	var L slot.Label
	m, cU := len(p.lex.Tokens)-1, 0
	p.ntAdd(symbols.NT_Orion, 0)
	// p.DumpDescriptors()
	for !p.R.empty() {
		L, cU, p.cI = p.R.remove()

		// fmt.Println()
		// fmt.Printf("L:%s, cI:%d, I[p.cI]:%s, cU:%d\n", L, p.cI, p.lex.Tokens[p.cI], cU)
		// p.DumpDescriptors()

		switch L {
		case slot.Data0R0: // Data : ∙String

			p.call(slot.Data0R1, cU, p.cI)
		case slot.Data0R1: // Data : String ∙

			if p.follow(symbols.NT_Data) {
				p.rtn(symbols.NT_Data, cU, p.cI)
			} else {
				p.parseError(slot.Data0R0, p.cI, followSets[symbols.NT_Data])
			}
		case slot.Data1R0: // Data : ∙FuncCall

			p.call(slot.Data1R1, cU, p.cI)
		case slot.Data1R1: // Data : FuncCall ∙

			if p.follow(symbols.NT_Data) {
				p.rtn(symbols.NT_Data, cU, p.cI)
			} else {
				p.parseError(slot.Data1R0, p.cI, followSets[symbols.NT_Data])
			}
		case slot.Data2R0: // Data : ∙Number

			p.call(slot.Data2R1, cU, p.cI)
		case slot.Data2R1: // Data : Number ∙

			if p.follow(symbols.NT_Data) {
				p.rtn(symbols.NT_Data, cU, p.cI)
			} else {
				p.parseError(slot.Data2R0, p.cI, followSets[symbols.NT_Data])
			}
		case slot.Data3R0: // Data : ∙Operation

			p.call(slot.Data3R1, cU, p.cI)
		case slot.Data3R1: // Data : Operation ∙

			if p.follow(symbols.NT_Data) {
				p.rtn(symbols.NT_Data, cU, p.cI)
			} else {
				p.parseError(slot.Data3R0, p.cI, followSets[symbols.NT_Data])
			}
		case slot.DataList0R0: // DataList : ∙Data

			p.call(slot.DataList0R1, cU, p.cI)
		case slot.DataList0R1: // DataList : Data ∙

			if p.follow(symbols.NT_DataList) {
				p.rtn(symbols.NT_DataList, cU, p.cI)
			} else {
				p.parseError(slot.DataList0R0, p.cI, followSets[symbols.NT_DataList])
			}
		case slot.DataList1R0: // DataList : ∙DataList , Data

			p.call(slot.DataList1R1, cU, p.cI)
		case slot.DataList1R1: // DataList : DataList ∙, Data

			if !p.testSelect(slot.DataList1R1) {
				p.parseError(slot.DataList1R1, p.cI, first[slot.DataList1R1])
				break
			}

			p.bsrSet.Add(slot.DataList1R2, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.DataList1R2) {
				p.parseError(slot.DataList1R2, p.cI, first[slot.DataList1R2])
				break
			}

			p.call(slot.DataList1R3, cU, p.cI)
		case slot.DataList1R3: // DataList : DataList , Data ∙

			if p.follow(symbols.NT_DataList) {
				p.rtn(symbols.NT_DataList, cU, p.cI)
			} else {
				p.parseError(slot.DataList1R0, p.cI, followSets[symbols.NT_DataList])
			}
		case slot.FuncCall0R0: // FuncCall : ∙ident ( DataList )

			p.bsrSet.Add(slot.FuncCall0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.FuncCall0R1) {
				p.parseError(slot.FuncCall0R1, p.cI, first[slot.FuncCall0R1])
				break
			}

			p.bsrSet.Add(slot.FuncCall0R2, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.FuncCall0R2) {
				p.parseError(slot.FuncCall0R2, p.cI, first[slot.FuncCall0R2])
				break
			}

			p.call(slot.FuncCall0R3, cU, p.cI)
		case slot.FuncCall0R3: // FuncCall : ident ( DataList ∙)

			if !p.testSelect(slot.FuncCall0R3) {
				p.parseError(slot.FuncCall0R3, p.cI, first[slot.FuncCall0R3])
				break
			}

			p.bsrSet.Add(slot.FuncCall0R4, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_FuncCall) {
				p.rtn(symbols.NT_FuncCall, cU, p.cI)
			} else {
				p.parseError(slot.FuncCall0R0, p.cI, followSets[symbols.NT_FuncCall])
			}
		case slot.FuncCall1R0: // FuncCall : ∙ident ( )

			p.bsrSet.Add(slot.FuncCall1R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.FuncCall1R1) {
				p.parseError(slot.FuncCall1R1, p.cI, first[slot.FuncCall1R1])
				break
			}

			p.bsrSet.Add(slot.FuncCall1R2, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.FuncCall1R2) {
				p.parseError(slot.FuncCall1R2, p.cI, first[slot.FuncCall1R2])
				break
			}

			p.bsrSet.Add(slot.FuncCall1R3, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_FuncCall) {
				p.rtn(symbols.NT_FuncCall, cU, p.cI)
			} else {
				p.parseError(slot.FuncCall1R0, p.cI, followSets[symbols.NT_FuncCall])
			}
		case slot.Import0R0: // Import : ∙get DataList

			p.bsrSet.Add(slot.Import0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.Import0R1) {
				p.parseError(slot.Import0R1, p.cI, first[slot.Import0R1])
				break
			}

			p.call(slot.Import0R2, cU, p.cI)
		case slot.Import0R2: // Import : get DataList ∙

			if p.follow(symbols.NT_Import) {
				p.rtn(symbols.NT_Import, cU, p.cI)
			} else {
				p.parseError(slot.Import0R0, p.cI, followSets[symbols.NT_Import])
			}
		case slot.Number0R0: // Number : ∙integer

			p.bsrSet.Add(slot.Number0R1, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_Number) {
				p.rtn(symbols.NT_Number, cU, p.cI)
			} else {
				p.parseError(slot.Number0R0, p.cI, followSets[symbols.NT_Number])
			}
		case slot.Number1R0: // Number : ∙float

			p.bsrSet.Add(slot.Number1R1, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_Number) {
				p.rtn(symbols.NT_Number, cU, p.cI)
			} else {
				p.parseError(slot.Number1R0, p.cI, followSets[symbols.NT_Number])
			}
		case slot.Operation0R0: // Operation : ∙Number

			p.call(slot.Operation0R1, cU, p.cI)
		case slot.Operation0R1: // Operation : Number ∙

			if p.follow(symbols.NT_Operation) {
				p.rtn(symbols.NT_Operation, cU, p.cI)
			} else {
				p.parseError(slot.Operation0R0, p.cI, followSets[symbols.NT_Operation])
			}
		case slot.Operation1R0: // Operation : ∙Operation op Number

			p.call(slot.Operation1R1, cU, p.cI)
		case slot.Operation1R1: // Operation : Operation ∙op Number

			if !p.testSelect(slot.Operation1R1) {
				p.parseError(slot.Operation1R1, p.cI, first[slot.Operation1R1])
				break
			}

			p.bsrSet.Add(slot.Operation1R2, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.Operation1R2) {
				p.parseError(slot.Operation1R2, p.cI, first[slot.Operation1R2])
				break
			}

			p.call(slot.Operation1R3, cU, p.cI)
		case slot.Operation1R3: // Operation : Operation op Number ∙

			if p.follow(symbols.NT_Operation) {
				p.rtn(symbols.NT_Operation, cU, p.cI)
			} else {
				p.parseError(slot.Operation1R0, p.cI, followSets[symbols.NT_Operation])
			}
		case slot.Orion0R0: // Orion : ∙Package Statements

			p.call(slot.Orion0R1, cU, p.cI)
		case slot.Orion0R1: // Orion : Package ∙Statements

			if !p.testSelect(slot.Orion0R1) {
				p.parseError(slot.Orion0R1, p.cI, first[slot.Orion0R1])
				break
			}

			p.call(slot.Orion0R2, cU, p.cI)
		case slot.Orion0R2: // Orion : Package Statements ∙

			if p.follow(symbols.NT_Orion) {
				p.rtn(symbols.NT_Orion, cU, p.cI)
			} else {
				p.parseError(slot.Orion0R0, p.cI, followSets[symbols.NT_Orion])
			}
		case slot.Package0R0: // Package : ∙package string_lit

			p.bsrSet.Add(slot.Package0R1, cU, p.cI, p.cI+1)
			p.cI++
			if !p.testSelect(slot.Package0R1) {
				p.parseError(slot.Package0R1, p.cI, first[slot.Package0R1])
				break
			}

			p.bsrSet.Add(slot.Package0R2, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_Package) {
				p.rtn(symbols.NT_Package, cU, p.cI)
			} else {
				p.parseError(slot.Package0R0, p.cI, followSets[symbols.NT_Package])
			}
		case slot.Statement0R0: // Statement : ∙FuncCall

			p.call(slot.Statement0R1, cU, p.cI)
		case slot.Statement0R1: // Statement : FuncCall ∙

			if p.follow(symbols.NT_Statement) {
				p.rtn(symbols.NT_Statement, cU, p.cI)
			} else {
				p.parseError(slot.Statement0R0, p.cI, followSets[symbols.NT_Statement])
			}
		case slot.Statement1R0: // Statement : ∙Import

			p.call(slot.Statement1R1, cU, p.cI)
		case slot.Statement1R1: // Statement : Import ∙

			if p.follow(symbols.NT_Statement) {
				p.rtn(symbols.NT_Statement, cU, p.cI)
			} else {
				p.parseError(slot.Statement1R0, p.cI, followSets[symbols.NT_Statement])
			}
		case slot.Statements0R0: // Statements : ∙Statement

			p.call(slot.Statements0R1, cU, p.cI)
		case slot.Statements0R1: // Statements : Statement ∙

			if p.follow(symbols.NT_Statements) {
				p.rtn(symbols.NT_Statements, cU, p.cI)
			} else {
				p.parseError(slot.Statements0R0, p.cI, followSets[symbols.NT_Statements])
			}
		case slot.Statements1R0: // Statements : ∙Statements Statement

			p.call(slot.Statements1R1, cU, p.cI)
		case slot.Statements1R1: // Statements : Statements ∙Statement

			if !p.testSelect(slot.Statements1R1) {
				p.parseError(slot.Statements1R1, p.cI, first[slot.Statements1R1])
				break
			}

			p.call(slot.Statements1R2, cU, p.cI)
		case slot.Statements1R2: // Statements : Statements Statement ∙

			if p.follow(symbols.NT_Statements) {
				p.rtn(symbols.NT_Statements, cU, p.cI)
			} else {
				p.parseError(slot.Statements1R0, p.cI, followSets[symbols.NT_Statements])
			}
		case slot.String0R0: // String : ∙string_lit

			p.bsrSet.Add(slot.String0R1, cU, p.cI, p.cI+1)
			p.cI++
			if p.follow(symbols.NT_String) {
				p.rtn(symbols.NT_String, cU, p.cI)
			} else {
				p.parseError(slot.String0R0, p.cI, followSets[symbols.NT_String])
			}

		default:
			panic("This must not happen")
		}
	}
	if !p.bsrSet.Contain(symbols.NT_Orion, 0, m) {
		p.sortParseErrors()
		return nil, p.parseErrors
	}
	return p.bsrSet, nil
}

func (p *parser) ntAdd(nt symbols.NT, j int) {
	// fmt.Printf("p.ntAdd(%s, %d)\n", nt, j)
	failed := true
	expected := map[token.Type]string{}
	for _, l := range slot.GetAlternates(nt) {
		if p.testSelect(l) {
			p.dscAdd(l, j, j)
			failed = false
		} else {
			for k, v := range first[l] {
				expected[k] = v
			}
		}
	}
	if failed {
		for _, l := range slot.GetAlternates(nt) {
			p.parseError(l, j, expected)
		}
	}
}

/*** Call Return Forest ***/

type poppedNode struct {
	X    symbols.NT
	k, j int
}

type clusterNode struct {
	X symbols.NT
	k int
}

type crfNode struct {
	L slot.Label
	i int
}

/*
suppose that L is Y ::=αX ·β
if there is no CRF node labelled (L,i)

	create one let u be the CRF node labelled (L,i)

if there is no CRF node labelled (X, j) {

		create a CRF node v labelled (X, j)
		create an edge from v to u
		ntAdd(X, j)
	} else {

		let v be the CRF node labelled (X, j)
		if there is not an edge from v to u {
			create an edge from v to u
			for all ((X, j,h)∈P) {
				dscAdd(L, i, h);
				bsrAdd(L, i, j, h)
			}
		}
	}
*/
func (p *parser) call(L slot.Label, i, j int) {
	// fmt.Printf("p.call(%s,%d,%d)\n", L,i,j)
	u, exist := p.crfNodes[crfNode{L, i}]
	// fmt.Printf("  u exist=%t\n", exist)
	if !exist {
		u = &crfNode{L, i}
		p.crfNodes[*u] = u
	}
	X := L.Symbols()[L.Pos()-1].(symbols.NT)
	ndV := clusterNode{X, j}
	v, exist := p.crf[ndV]
	if !exist {
		// fmt.Println("  v !exist")
		p.crf[ndV] = []*crfNode{u}
		p.ntAdd(X, j)
	} else {
		// fmt.Println("  v exist")
		if !existEdge(v, u) {
			// fmt.Printf("  !existEdge(%v)\n", u)
			p.crf[ndV] = append(v, u)
			// fmt.Printf("|popped|=%d\n", len(popped))
			for pnd := range p.popped {
				if pnd.X == X && pnd.k == j {
					p.dscAdd(L, i, pnd.j)
					p.bsrSet.Add(L, i, j, pnd.j)
				}
			}
		}
	}
}

func existEdge(nds []*crfNode, nd *crfNode) bool {
	for _, nd1 := range nds {
		if nd1 == nd {
			return true
		}
	}
	return false
}

func (p *parser) rtn(X symbols.NT, k, j int) {
	// fmt.Printf("p.rtn(%s,%d,%d)\n", X,k,j)
	pn := poppedNode{X, k, j}
	if _, exist := p.popped[pn]; !exist {
		p.popped[pn] = true
		for _, nd := range p.crf[clusterNode{X, k}] {
			p.dscAdd(nd.L, nd.i, j)
			p.bsrSet.Add(nd.L, nd.i, k, j)
		}
	}
}

// func CRFString() string {
// 	buf := new(bytes.Buffer)
// 	buf.WriteString("CRF: {")
// 	for cn, nds := range crf{
// 		for _, nd := range nds {
// 			fmt.Fprintf(buf, "%s->%s, ", cn, nd)
// 		}
// 	}
// 	buf.WriteString("}")
// 	return buf.String()
// }

func (cn clusterNode) String() string {
	return fmt.Sprintf("(%s,%d)", cn.X, cn.k)
}

func (n crfNode) String() string {
	return fmt.Sprintf("(%s,%d)", n.L.String(), n.i)
}

// func PoppedString() string {
// 	buf := new(bytes.Buffer)
// 	buf.WriteString("Popped: {")
// 	for p, _ := range popped {
// 		fmt.Fprintf(buf, "(%s,%d,%d) ", p.X, p.k, p.j)
// 	}
// 	buf.WriteString("}")
// 	return buf.String()
// }

/*** descriptors ***/

type descriptors struct {
	set []*descriptor
}

func (ds *descriptors) contain(d *descriptor) bool {
	for _, d1 := range ds.set {
		if d1 == d {
			return true
		}
	}
	return false
}

func (ds *descriptors) empty() bool {
	return len(ds.set) == 0
}

func (ds *descriptors) String() string {
	buf := new(bytes.Buffer)
	buf.WriteString("{")
	for i, d := range ds.set {
		if i > 0 {
			buf.WriteString("; ")
		}
		fmt.Fprintf(buf, "%s", d)
	}
	buf.WriteString("}")
	return buf.String()
}

type descriptor struct {
	L slot.Label
	k int
	i int
}

func (d *descriptor) String() string {
	return fmt.Sprintf("%s,%d,%d", d.L, d.k, d.i)
}

func (p *parser) dscAdd(L slot.Label, k, i int) {
	// fmt.Printf("p.dscAdd(%s,%d,%d)\n", L, k, i)
	d := &descriptor{L, k, i}
	if !p.U.contain(d) {
		p.R.set = append(p.R.set, d)
		p.U.set = append(p.U.set, d)
	}
}

func (ds *descriptors) remove() (L slot.Label, k, i int) {
	d := ds.set[len(ds.set)-1]
	ds.set = ds.set[:len(ds.set)-1]
	// fmt.Printf("remove: %s,%d,%d\n", d.L, d.k, d.i)
	return d.L, d.k, d.i
}

func (p *parser) DumpDescriptors() {
	p.DumpR()
	p.DumpU()
}

func (p *parser) DumpR() {
	fmt.Println("R:")
	for _, d := range p.R.set {
		fmt.Printf(" %s\n", d)
	}
}

func (p *parser) DumpU() {
	fmt.Println("U:")
	for _, d := range p.U.set {
		fmt.Printf(" %s\n", d)
	}
}

/*** TestSelect ***/

func (p *parser) follow(nt symbols.NT) bool {
	_, exist := followSets[nt][p.lex.Tokens[p.cI].Type()]
	return exist
}

func (p *parser) testSelect(l slot.Label) bool {
	_, exist := first[l][p.lex.Tokens[p.cI].Type()]
	// fmt.Printf("testSelect(%s) = %t\n", l, exist)
	return exist
}

var first = []map[token.Type]string{
	// Data : ∙String
	{
		token.T_11: "string_lit",
	},
	// Data : String ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Data : ∙FuncCall
	{
		token.T_6: "ident",
	},
	// Data : FuncCall ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Data : ∙Number
	{
		token.T_4: "float",
		token.T_7: "integer",
	},
	// Data : Number ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Data : ∙Operation
	{
		token.T_4: "float",
		token.T_7: "integer",
	},
	// Data : Operation ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// DataList : ∙Data
	{
		token.T_4:  "float",
		token.T_6:  "ident",
		token.T_7:  "integer",
		token.T_11: "string_lit",
	},
	// DataList : Data ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// DataList : ∙DataList , Data
	{
		token.T_4:  "float",
		token.T_6:  "ident",
		token.T_7:  "integer",
		token.T_11: "string_lit",
	},
	// DataList : DataList ∙, Data
	{
		token.T_2: ",",
	},
	// DataList : DataList , ∙Data
	{
		token.T_4:  "float",
		token.T_6:  "ident",
		token.T_7:  "integer",
		token.T_11: "string_lit",
	},
	// DataList : DataList , Data ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// FuncCall : ∙ident ( DataList )
	{
		token.T_6: "ident",
	},
	// FuncCall : ident ∙( DataList )
	{
		token.T_0: "(",
	},
	// FuncCall : ident ( ∙DataList )
	{
		token.T_4:  "float",
		token.T_6:  "ident",
		token.T_7:  "integer",
		token.T_11: "string_lit",
	},
	// FuncCall : ident ( DataList ∙)
	{
		token.T_1: ")",
	},
	// FuncCall : ident ( DataList ) ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// FuncCall : ∙ident ( )
	{
		token.T_6: "ident",
	},
	// FuncCall : ident ∙( )
	{
		token.T_0: "(",
	},
	// FuncCall : ident ( ∙)
	{
		token.T_1: ")",
	},
	// FuncCall : ident ( ) ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Import : ∙get DataList
	{
		token.T_5: "get",
	},
	// Import : get ∙DataList
	{
		token.T_4:  "float",
		token.T_6:  "ident",
		token.T_7:  "integer",
		token.T_11: "string_lit",
	},
	// Import : get DataList ∙
	{
		token.EOF: "$",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Number : ∙integer
	{
		token.T_7: "integer",
	},
	// Number : integer ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
		token.T_9: "op",
	},
	// Number : ∙float
	{
		token.T_4: "float",
	},
	// Number : float ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
		token.T_9: "op",
	},
	// Operation : ∙Number
	{
		token.T_4: "float",
		token.T_7: "integer",
	},
	// Operation : Number ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
		token.T_9: "op",
	},
	// Operation : ∙Operation op Number
	{
		token.T_4: "float",
		token.T_7: "integer",
	},
	// Operation : Operation ∙op Number
	{
		token.T_9: "op",
	},
	// Operation : Operation op ∙Number
	{
		token.T_4: "float",
		token.T_7: "integer",
	},
	// Operation : Operation op Number ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
		token.T_9: "op",
	},
	// Orion : ∙Package Statements
	{
		token.T_10: "package",
	},
	// Orion : Package ∙Statements
	{
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Orion : Package Statements ∙
	{
		token.EOF: "$",
	},
	// Package : ∙package string_lit
	{
		token.T_10: "package",
	},
	// Package : package ∙string_lit
	{
		token.T_11: "string_lit",
	},
	// Package : package string_lit ∙
	{
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Statement : ∙FuncCall
	{
		token.T_6: "ident",
	},
	// Statement : FuncCall ∙
	{
		token.EOF: "$",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Statement : ∙Import
	{
		token.T_5: "get",
	},
	// Statement : Import ∙
	{
		token.EOF: "$",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Statements : ∙Statement
	{
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Statements : Statement ∙
	{
		token.EOF: "$",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Statements : ∙Statements Statement
	{
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Statements : Statements ∙Statement
	{
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Statements : Statements Statement ∙
	{
		token.EOF: "$",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// String : ∙string_lit
	{
		token.T_11: "string_lit",
	},
	// String : string_lit ∙
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
}

var followSets = []map[token.Type]string{
	// Data
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// DataList
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// FuncCall
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Import
	{
		token.EOF: "$",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Number
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
		token.T_9: "op",
	},
	// Operation
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
		token.T_9: "op",
	},
	// Orion
	{
		token.EOF: "$",
	},
	// Package
	{
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Statement
	{
		token.EOF: "$",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// Statements
	{
		token.EOF: "$",
		token.T_5: "get",
		token.T_6: "ident",
	},
	// String
	{
		token.EOF: "$",
		token.T_1: ")",
		token.T_2: ",",
		token.T_5: "get",
		token.T_6: "ident",
	},
}

/*** Errors ***/

/*
Error is returned by Parse at every point at which the parser fails to parse
a grammar production. For non-LL-1 grammars there will be an error for each
alternate attempted by the parser.

The errors are sorted in descending order of input position (index of token in
the stream of tokens).

Normally the error of interest is the one that has parsed the largest number of
tokens.
*/
type Error struct {
	// Index of token that caused the error.
	cI int

	// Grammar slot at which the error occured.
	Slot slot.Label

	// The token at which the error occurred.
	Token *token.Token

	// The line and column in the input text at which the error occurred
	Line, Column int

	// The tokens expected at the point where the error occurred
	Expected map[token.Type]string
}

func (pe *Error) String() string {
	w := new(bytes.Buffer)
	fmt.Fprintf(w, "Parse Error: %s I[%d]=%s at line %d col %d\n",
		pe.Slot, pe.cI, pe.Token, pe.Line, pe.Column)
	exp := []string{}
	for _, e := range pe.Expected {
		exp = append(exp, e)
	}
	fmt.Fprintf(w, "Expected one of: [%s]", strings.Join(exp, ","))
	return w.String()
}

func (p *parser) parseError(slot slot.Label, i int, expected map[token.Type]string) {
	pe := &Error{cI: i, Slot: slot, Token: p.lex.Tokens[i], Expected: expected}
	p.parseErrors = append(p.parseErrors, pe)
}

func (p *parser) sortParseErrors() {
	sort.Slice(p.parseErrors,
		func(i, j int) bool {
			return p.parseErrors[j].Token.Lext() < p.parseErrors[i].Token.Lext()
		})
	for _, pe := range p.parseErrors {
		pe.Line, pe.Column = p.lex.GetLineColumn(pe.Token.Lext())
	}
}
