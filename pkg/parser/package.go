package parser

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/token"
	"go/types"
)

type Package struct {
	dir      string
	name     string
	defs     map[*ast.Ident]types.Object
	files    []*File
	typesPkg *types.Package
}

// check type-checks the package. The package must be OK to proceed.
func (pkg *Package) check(fs *token.FileSet, astFiles []*ast.File) error {
	pkg.defs = make(map[*ast.Ident]types.Object)
	config := types.Config{Importer: importer.Default(), FakeImportC: true}
	info := &types.Info{
		Defs: pkg.defs,
	}
	typesPkg, err := config.Check(pkg.dir, fs, astFiles, info)
	if err != nil {
		return fmt.Errorf("checking package: %s", err)
	}
	pkg.typesPkg = typesPkg
	return nil
}
