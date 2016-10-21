package parser

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

// ParsePackageDir parses the package residing in the directory.
func ParsePackageDir(directory string) (*Package, error) {
	pkg, err := build.Default.ImportDir(directory, 0)
	if err != nil {
		return nil, fmt.Errorf("cannot process directory %s: %s", directory, err)
	}
	var names []string
	names = append(names, pkg.GoFiles...)
	names = append(names, pkg.CgoFiles...)
	// TODO: Need to think about constants in test files. Maybe write type_string_test.go
	// in a separate pass? For later.
	// names = append(names, pkg.TestGoFiles...) // These are also in the "foo" package.
	names = append(names, pkg.SFiles...)
	names = prefixDirectory(directory, names)
	return ParsePackage(directory, names, nil)
}

// ParsePackageFiles parses the package occupying the named files.
func ParsePackageFiles(names []string) (*Package, error) {
	return ParsePackage(".", names, nil)
}

// prefixDirectory places the directory name on the beginning of each name in the list.
func prefixDirectory(directory string, names []string) []string {
	if directory == "." {
		return names
	}
	ret := make([]string, len(names))
	for i, name := range names {
		ret[i] = filepath.Join(directory, name)
	}
	return ret
}

// ParsePackage analyzes the single package constructed from the named files.
// If text is non-nil, it is a string to be used instead of the content of the file,
// to be used for testing. parsePackage exits if there is an error.
func ParsePackage(directory string, names []string, text interface{}) (*Package, error) {
	var files []*File
	var astFiles []*ast.File
	pkg := new(Package)
	fs := token.NewFileSet()
	for _, name := range names {
		if !strings.HasSuffix(name, ".go") {
			continue
		}
		parsedFile, err := parser.ParseFile(fs, name, text, 0)
		if err != nil {
			return nil, fmt.Errorf("parsing package: %s: %s", name, err)
		}
		astFiles = append(astFiles, parsedFile)
		files = append(files, &File{
			file: parsedFile,
			pkg:  pkg,
		})
	}
	if len(astFiles) == 0 {
		return nil, fmt.Errorf("%s: no buildable Go files", directory)
	}
	pkg.name = astFiles[0].Name.Name
	pkg.files = files
	pkg.dir = directory
	// Type check the package.
	return pkg, pkg.check(fs, astFiles)
}
