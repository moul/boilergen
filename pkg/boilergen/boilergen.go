package boilergen

import (
	"fmt"

	"github.com/moul/boilergen/pkg/asttree"
	boilerparser "github.com/moul/boilergen/pkg/parser"
)

var VERSION = "0.1.0"

type Boilergen struct {
	outputDir    string
	templatesDir string

	pkg *boilerparser.Package
}

func New() Boilergen {
	return Boilergen{}
}

func (b *Boilergen) ParsePackageDir(directory string) error {
	pkg, err := boilerparser.ParsePackageDir(directory)
	if err != nil {
		return err
	}
	b.pkg = pkg
	return nil
}

func (b *Boilergen) SetOutputDirectory(path string)    { b.outputDir = path }
func (b *Boilergen) SetTemplatesDirectory(path string) { b.templatesDir = path }

func (b *Boilergen) Generate() error {
	tree, err := asttree.FromParserPackage(b.pkg)
	if err != nil {
		return err
	}
	fmt.Println(tree)
	return nil
}
