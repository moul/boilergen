package parser

import "go/ast"

// File holds a single parsed file and associated data.
type File struct {
	Name string    // Filename
	Pkg  *Package  // Package to which this file belongs
	File *ast.File // Parsed AST
}
