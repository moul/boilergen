package boilergen

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/kr/fs"
	"github.com/moul/boilergen/pkg/asttree"
	boilerparser "github.com/moul/boilergen/pkg/parser"
	"github.com/moul/funcmap"
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

	stat, err := os.Stat(b.outputDir)
	if os.IsNotExist(err) {
		os.MkdirAll(b.outputDir, 0755)
	} else if !stat.IsDir() {
		return fmt.Errorf("%s: not a directory", b.outputDir)
	}

	walker := fs.Walk(b.templatesDir)
	for walker.Step() {
		if err := walker.Err(); err != nil {
			return err
		}

		if walker.Stat().IsDir() {
			continue
		}

		if filepath.Ext(walker.Path()) != ".tmpl" {
			continue
		}

		basename := path.Base(walker.Path())
		tmpl, err := template.
			New(basename).
			Funcs(funcmap.FuncMap).
			ParseFiles(walker.Path())
		if err != nil {
			return err
		}

		rel, err := filepath.Rel(b.templatesDir, walker.Path())
		if err != nil {
			return err
		}

		nameWithoutExtension := rel[0 : len(rel)-len(".tmpl")]
		destFile, err := os.Create(path.Join(b.outputDir, nameWithoutExtension))
		if err != nil {
			return err
		}
		defer destFile.Close()

		if err := tmpl.Execute(destFile, tree); err != nil {
			return err
		}
	}
	return nil
}
