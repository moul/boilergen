package asttree

import (
	"go/types"

	"github.com/moul/boilergen/pkg/parser"
)

type Tree struct {
	Common
	Imports []Import
	Method  []Method
}

type Common struct {
	BasePackage struct {
		Dir      string `json:"Dir"`
		Name     string `json:"Name"`
		Defs     []types.Object
		Files    []*parser.File
		TypesPkg *types.Package
	} `json:"BasePackage"`
}
type Param struct{}
type Import struct{}
type Method struct{}

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
