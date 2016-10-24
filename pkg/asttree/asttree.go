package asttree

import (
	"go/types"

	"github.com/moul/boilergen/pkg/parser"
)

type Tree struct {
	Common
	Imports []Import
	Methods []Method
	Types   []Type
	Defs    []Def
}

type Common struct {
	BuildCommand string
	BasePackage  struct {
		Dir      string `json:"Dir"`
		Name     string `json:"Name"`
		Defs     []types.Object
		Files    []*parser.File
		TypesPkg *types.Package
	} `json:"BasePackage"`
}

type Def struct {
	Name       string
	ObjectName string
	Decl       interface{}
	Data       interface{}
	Type       interface{}
}

type Type struct {
	Common
	Name              string
	Value             string
	ValueExact        string
	PrivateProperties []Property
	PublicProperties  []Property
	PrivateMethods    []Method
	PublicMethods     []Method
}

type Property struct {
	Common
	Name string
}

type Import struct {
	Common
	Name string
}

type Method struct {
	Common
	Name string
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

	// ast.Print(input.FS, input.Files[0].File)

	for def := range input.Info.Defs {
		obj := Def{
			Name: def.Name,
		}
		if def.Obj != nil {
			obj.ObjectName = def.Obj.Name
			// log.Printf("decl: %v", def.Obj.Decl)
			// obj.Decl = def.Obj.Decl
			obj.Data = def.Obj.Data
			obj.Type = def.Obj.Type
		}
		tree.Defs = append(tree.Defs, obj)
	}
	for _, typ := range input.Info.Types {
		tree.Types = append(tree.Types, Type{
			Name:       typ.Type.String(),
			Value:      typ.Value.String(),
			ValueExact: typ.Value.ExactString(),
		})
	}

	return tree, nil
}
