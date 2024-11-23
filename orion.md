<!-- 
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
-->

# Orion

This is the Grammar File for the Orion language.

```
package "github.com/Solarcode-org/Orion"
```

## Lexical Rules

### Identifier:
```
ident : [letter|'@'] {letter|number|'_'|'?'|'.'} ;
```

### Comments:
```
!line_comment : ';' {not "\n"} ;
!block_comment : '/''*' {not "*" | '*' not "/"} '*''/' ;
```

### String Literal:
```
string_lit : '"' <not "\\\"" | '\\' any "\\\"nrt"> '"' ;
```

### Integers and Floats:
```
integer : number {number} ;
float : integer ['.' integer] ;
```

### Operation Symbols:
```
op : ('+' | '-' | '*' | '/') ;
```

## The Orion Rule

It is the root rule for all statements in Orion.
An Orion program starts with the [package directive](#the-package-rule),
which is followed by all the [statements](#the-statements-plural-rule).

```
Orion : Package Statements ;
```

## The Package Rule

> Syntax: package NAME

The package name must be a string.

```
Package : "package" string_lit ;
```

## The Statements (plural) Rule

It represents the collection of all the [statements](#the-statement-singular-rule) in the program.

```
Statements
        : Statement
        | Statements Statement
        ;
```

## The Statement (singular) Rule

It represents four types of statements:
  1. [Function definitions](#todo-the-function-definition-rule) (_TODO_)
  2. [Function calls](#the-function-call-rule)
  3. [Variable definitions](#the-variable-definition-rule)
  4. [Imports](#the-import-rule)

```
Statement : FuncCall | Import | VariableDef ; 
```

## (_TODO_): The Function Definition Rule

It represents a function definition.

## The Function Call Rule

It represent a function call.
The arguments are in the form of a [DataList](#the-datalist-rule).

```
FuncCall
        : ident "(" DataList ")"
        | ident "(" ")"
        ;
```

## The Import Rule

It represent an import of a module or package
The arguments are in the form of a [DataList](#the-datalist-rule).

```
Import : "get" DataList ;
```

## The Variable Definition Rule

It represents a variable definition.
A variable declaration statement is in the form:
<blockquote>variableName := value</blockquote>
or
<blockquote>variableName: type = value</blockquote>

```
VariableDef
        : ident ":=" Data
        | ident ":" ident "=" Data
        ;
```

## The DataList Rule

It represents a collection of [Data](#the-data-rule).

```
DataList
        : Data
        | DataList "," Data
        ;
```

## The Data Rule

It represents all data types in Orion: string, numbers, etc.
It also represents variables and operations.

```
Data : String | FuncCall | Number | Operation | Variable ;
```

### The String Rule

It represents all quoted strings in Orion.

```
String : string_lit ;
```

### The Number Rule

It represents all integers and floating-point numbers in Orion.

```
Number : integer | float ;
```

### The Variable Rule

It represents all variables in Orion.

```
Variable : ident ;
```

### The Operation Rule

It represents all the operations in Orion.

```
Operation 
        : Data op Data
        ;
```
