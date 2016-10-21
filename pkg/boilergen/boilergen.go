package boilergen

import (
	"html/template"
	"os"

	"github.com/kr/fs"
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

	walker := fs.Walk(b.templatesDir)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			return err
		}
		if walker.Stat().IsDir() {
			continue
		}
		tmpl, err := template.ParseFiles(walker.Path())
		if err != nil {
			return err
		}
		if err := tmpl.Execute(os.Stdout, tree); err != nil {
			return err
		}
	}
	return nil
}
