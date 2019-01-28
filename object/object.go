package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"monkey_interpreter/ast"
	"strings"
)

// ObjectType string
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
	FUNCTION_OBJ     = "FUNCTION"
	STRING_OBJ       = "STRING"
	BUILTIN_OBJ      = "BUILTIN"
	ARRAY_OBJ        = "ARRAY"
	HASH_OBJ         = "HASH"
)

// Object interface
type Object interface {
	Type() ObjectType
	Inspect() string
}

// Integer struct
type Integer struct {
	Value int64
}

// Inspect method for integer type
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type method returns integer ObjectType
func (i *Integer) Type() ObjectType { return INTEGER_OBJ }

// Boolean struct
type Boolean struct {
	Value bool
}

// Type method returns boolean ObjectType
func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ }

// Inspect method for boolean type
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Null struct
type Null struct{}

// Type method returns null ObjectType
func (n *Null) Type() ObjectType { return NULL_OBJ }

// Inspect method for null type
func (n *Null) Inspect() string { return "null" }

// ReturnValue struct
type ReturnValue struct {
	Value Object
}

// Type method returns return ObjectType
func (rv *ReturnValue) Type() ObjectType { return RETURN_VALUE_OBJ }

// Inspect method for ReturnValue type
func (rv *ReturnValue) Inspect() string { return rv.Value.Inspect() }

// Error struct
type Error struct {
	Message string
}

// Type method returns Error ObjectType
func (e *Error) Type() ObjectType { return ERROR_OBJ }

// Inspect method for Error type
func (e *Error) Inspect() string { return "ERROR: " + e.Message }

// Function struct
type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

// Type returns function ObjType
func (f *Function) Type() ObjectType { return FUNCTION_OBJ }

// Inspect method for function type
func (f *Function) Inspect() string {
	var out bytes.Buffer

	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// String struct
type String struct {
	Value string
}

// Type returns String object type
func (s *String) Type() ObjectType { return STRING_OBJ }

// Inspect method for String type
func (s *String) Inspect() string { return s.Value }

// Builtin struct
type Builtin struct {
	Fn BuiltinFunction
}

// Type returns Builtin Object Type
func (b *Builtin) Type() ObjectType { return BUILTIN_OBJ }

// Inspect method for Builtin Type
func (b *Builtin) Inspect() string { return "builtin function" }

// BuiltinFunction function
type BuiltinFunction func(args ...Object) Object

// Array struct
type Array struct {
	Elements []Object
}

// Type returns Array object type
func (ao *Array) Type() ObjectType { return ARRAY_OBJ }

// Inspect method for Array type
func (ao *Array) Inspect() string {
	var out bytes.Buffer

	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

// HashKey struct
type HashKey struct {
	Type  ObjectType
	Value uint64
}

// HashKey function booleans
func (b *Boolean) HashKey() HashKey {
	var value uint64

	if b.Value {
		value = 1
	} else {
		value = 0
	}

	return HashKey{Type: b.Type(), Value: value}
}

// HashKey function Ints
func (i *Integer) HashKey() HashKey {
	return HashKey{Type: i.Type(), Value: uint64(i.Value)}
}

// HashKey function Strings
func (s *String) HashKey() HashKey {
	h := fnv.New64a()
	h.Write([]byte(s.Value))

	return HashKey{Type: s.Type(), Value: h.Sum64()}
}

// HashPair struct
type HashPair struct {
	Key   Object
	Value Object
}

// Hash struct
type Hash struct {
	Pairs map[HashKey]HashPair
}

// Type function returns Hash ObjectType
func (h *Hash) Type() ObjectType { return HASH_OBJ }

// Inspect method for Hash objects
func (h *Hash) Inspect() string {
	var out bytes.Buffer

	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s", pair.Key.Inspect(), pair.Value.Inspect()))
	}

	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")

	return out.String()
}

// Hashable interface
type Hashable interface {
	HashKey() HashKey
}
