package asttree

type Package struct {
	Files []*File

	BuildCommand string
	Name         string
	Dir          string

	Defs    []*Def
	Imports []*Import
	Methods []*Method
	Structs []*Struct
	Types   []*Type
}

func NewPackage() *Package {
	return &Package{
		Defs:    make([]*Def, 0),
		Imports: make([]*Import, 0),
		Methods: make([]*Method, 0),
		Structs: make([]*Struct, 0),
		Types:   make([]*Type, 0),
	}
}

func (pkg *Package) appendFile(file *File) {
	pkg.Files = append(pkg.Files)

	pkg.Defs = append(pkg.Defs, file.Defs...)
	pkg.Imports = append(pkg.Imports, file.Imports...)
	pkg.Methods = append(pkg.Methods, file.Methods...)
	pkg.Structs = append(pkg.Structs, file.Structs...)
	pkg.Types = append(pkg.Types, file.Types...)
}
