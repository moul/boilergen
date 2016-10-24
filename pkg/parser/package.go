package parser

import (
	"fmt"
	"go/ast"
	"go/importer"
	"go/token"
	"go/types"
)

type Package struct {
	Dir      string
	Name     string
	Defs     map[*ast.Ident]types.Object
	Files    []*File
	TypesPkg *types.Package
	Info     *types.Info
	FS       *token.FileSet
}

// check type-checks the package. The package must be OK to proceed.
func (p *Package) check(fs *token.FileSet, astFiles []*ast.File) error {
	p.Defs = make(map[*ast.Ident]types.Object)
	config := types.Config{Importer: importer.Default(), FakeImportC: true}
	p.Info = &types.Info{
		Defs: p.Defs,
	}
	typesPkg, err := config.Check(p.Dir, fs, astFiles, p.Info)
	if err != nil {
		return fmt.Errorf("checking package: %s", err)
	}
	p.TypesPkg = typesPkg
	return nil
}
