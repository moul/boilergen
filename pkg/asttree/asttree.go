package asttree

import (
	"go/ast"
	"go/types"
	"log"
	"os"

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

func (t *Tree) populate(node ast.Node) bool {
	log.Print(node)
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
