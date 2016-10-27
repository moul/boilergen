package asttree

import (
	"go/ast"
	"go/types"
	"log"
	"os"
	"unicode"

	"github.com/moul/boilergen/pkg/parser"
)

func (tree *Tree) populate(node ast.Node) bool {
	switch t := node.(type) {
	case *ast.TypeSpec:
		switch s := t.Type.(type) {
		case *ast.StructType:
			if !t.Name.IsExported() {
				return false
			}
			theStruct := Struct{
				Name: t.Name.Name,
			}
			theType := Type{
				Name: t.Name.Name,
				// IsBasePackage=
			}
			for _, f := range s.Fields.List {
				isPublic := false
				field := Field{
					Struct: &theStruct,
					Names:  make([]string, 0),
					Type: Type{
						Name: resolveFieldTypes(f.Type, tree.Common.BasePackage.Name),
					},
				}
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
			tree.Types = append(tree.Types, theType)
			tree.Structs = append(tree.Structs, theStruct)
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

func FromParserPackage(input *parser.Package) (Tree, error) {
	// common
	common := Common{}
	common.BasePackage.Dir = input.Dir
	common.BasePackage.Name = input.Name
	common.BasePackage.TypesPkg = input.TypesPkg
	common.BasePackage.Defs = make([]types.Object, 0)
	for _, def := range input.Defs {
		common.BasePackage.Defs = append(common.BasePackage.Defs, def)
	}
	common.BasePackage.Files = input.Files

	// tree
	tree := Tree{}
	tree.Common = common
	tree.Types = make([]Type, 0)
	tree.Defs = make([]Def, 0)

	for _, file := range input.Files {
		// debug
		log.Printf("ast.Print(%q)", file.Name)
		ast.Fprint(os.Stderr, input.FS, file.File, ast.NotNilFilter)

		// populate
		ast.Inspect(file.File, tree.populate)
	}

	return tree, nil
}
