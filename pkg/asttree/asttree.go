package asttree

import (
	"go/ast"

	"github.com/moul/boilergen/pkg/parser"
)

func FromParserPackage(input *parser.Package) (*Package, error) {
	pkg := NewPackage()
	pkg.Name = input.Name
	pkg.Dir = input.Dir

	for _, f := range input.Files {
		// debug
		//log.Printf("ast.Print(%q)", f.Name)
		//ast.Fileprint(os.Stderr, input.FileS, f.File, ast.NotNilFilter)

		// populate
		file := NewFile()
		file.Package = pkg
		ast.Inspect(f.File, file.populate)
		pkg.appendFile(file)
	}

	return pkg, nil
}
