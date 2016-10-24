package parser

import "go/ast"

// File holds a single parsed file and associated data.
type File struct {
	Pkg  *Package  // Package to which this file belongs.
	File *ast.File // Parsed AST.
	// These fields are reset for each type being generated.
	TypeName string  // Name of the constant type.
	Values   []Value // Accumulator for constant values of that type.
}
