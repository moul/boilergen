package asttree

import "github.com/moul/boilergen/pkg/parser"

type Tree struct {
	Common
	Imports []Import
	Method  []Method
}

type Common struct {
	BasePackage       string
	BasePackageImport string
	BasePackageName   string
	InterfaceName     string
}
type Param struct{}
type Import struct{}
type Method struct{}

func FromParserPackage(input *parser.Package) (Tree, error) {
	tree := Tree{}

	common := Common{
		BasePackageName: input.GetName(),
	}
	tree.Common = common

	return tree, nil
}
