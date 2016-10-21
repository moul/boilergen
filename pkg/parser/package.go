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
func (p *Package) check(fs *token.FileSet, astFiles []*ast.File) error {
	p.defs = make(map[*ast.Ident]types.Object)
	config := types.Config{Importer: importer.Default(), FakeImportC: true}
	info := &types.Info{
		Defs: p.defs,
	}
	typesPkg, err := config.Check(p.dir, fs, astFiles, info)
	if err != nil {
		return fmt.Errorf("checking package: %s", err)
	}
	p.typesPkg = typesPkg
	return nil
}

// getters
func (p *Package) GetDir() string                       { return p.dir }
func (p *Package) GetName() string                      { return p.name }
func (p *Package) GetDefs() map[*ast.Ident]types.Object { return p.defs }
func (p *Package) GetFiles() []*File                    { return p.files }
func (p *Package) GetTypesPkg() *types.Package          { return p.typesPkg }
