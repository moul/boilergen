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

type Type struct {
	Common
	Name              string
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
	tree := Tree{}

	common := Common{}
	common.BasePackage.Dir = input.GetDir()
	common.BasePackage.Name = input.GetName()
	common.BasePackage.TypesPkg = input.GetTypesPkg()
	common.BasePackage.Defs = make([]types.Object, 0)
	for _, def := range input.GetDefs() {
		common.BasePackage.Defs = append(common.BasePackage.Defs, def)
	}
	common.BasePackage.Files = input.GetFiles()
	tree.Common = common

	return tree, nil
}
