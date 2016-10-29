package asttree

import (
	"go/ast"
	"log"
	"unicode"
)

type File struct {
	Package *Package

	Path     string
	FileName string

	Defs    []*Def
	Imports []*Import
	Methods []*Method
	Structs []*Struct
	Types   []*Type
}

func NewFile() *File {
	return &File{
		Defs:    make([]*Def, 0),
		Imports: make([]*Import, 0),
		Methods: make([]*Method, 0),
		Structs: make([]*Struct, 0),
		Types:   make([]*Type, 0),
	}
}

func (file *File) populate(node ast.Node) bool {
	switch t := node.(type) {
	case *ast.TypeSpec:
		switch s := t.Type.(type) {
		case *ast.StructType:
			if !t.Name.IsExported() {
				return false
			}
			theStruct := NewStruct()
			theStruct.Name = t.Name.Name
			theType := NewType()
			theType.Name = t.Name.Name
			// theType.IsBasePackage =
			for _, f := range s.Fields.List {
				isPublic := false

				field := NewField()
				field.Struct = theStruct
				field.Type = NewType()
				field.Type.Name = resolveFieldTypes(f.Type, file.Package.Name)

				for _, n := range f.Names {
					field.Names = append(field.Names, n.Name)
					if unicode.IsUpper(rune(n.Name[0])) {
						isPublic = true
					}
				}
				field.Name = field.Names[0]
				if isPublic {
					theStruct.PublicFields = append(theStruct.PublicFields, field)
				} else {
					theStruct.PrivateFields = append(theStruct.PrivateFields, field)
				}
			}
			file.Types = append(file.Types, theType)
			file.Structs = append(file.Structs, theStruct)
			break
		}
		break
	case *ast.FuncDecl, *ast.Ident, *ast.File, *ast.Field, *ast.MapType, *ast.ImportSpec, *ast.StructType, *ast.GenDecl, *ast.ValueSpec, *ast.FieldList, *ast.ArrayType:
		log.Printf("%#v", node)
		return true
	case *ast.CommentGroup, *ast.Comment, nil:
		return false
	default:
		log.Printf("Unknown node type: %#v", t)
	}
	return true
}
