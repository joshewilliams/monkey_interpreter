package object

import "fmt"

// ObjectType string
type ObjectType string

const (
	INTEGER_OBJ = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
	NULL_OBJ    = "NULL"
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
