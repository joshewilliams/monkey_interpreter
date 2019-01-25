package object

import "fmt"

// ObjectType string
type ObjectType string

const (
	INTEGER_OBJ      = "INTEGER"
	BOOLEAN_OBJ      = "BOOLEAN"
	NULL_OBJ         = "NULL"
	RETURN_VALUE_OBJ = "RETURN_VALUE"
	ERROR_OBJ        = "ERROR"
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

// NewEnvironment environment
func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s}
}

// Environment struct
type Environment struct {
	store map[string]Object
}

// Get ter for environment
func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	return obj, ok
}

// Set ter for environment
func (e *Environment) Set(name string, val Object) Object {
	e.store[name] = val
	return val
}
