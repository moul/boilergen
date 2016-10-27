package asttree

import (
	"go/types"

	"github.com/moul/boilergen/pkg/parser"
)

type Tree struct {
	Common
	Imports []Import
	Methods []Method
	Types   []Type
	Defs    []Def
	Structs []Struct
}

type Common struct {
	BuildCommand string
	BasePackage  struct {
		Dir      string `json:"Dir"`
		Name     string `json:"Name"`
		Defs     []types.Object
		Files    []*parser.File
		TypesPkg *types.Package
	} `json:"BasePackage"`
}

type Def struct {
	Name       string
	ObjectName string
	Decl       interface{}
	Data       interface{}
	Type       interface{}
}

type Type struct {
	Name          string
	IsBasePackage bool
}

type Struct struct {
	Common
	Types          []Type
	Name           string
	Value          string
	ValueExact     string
	PrivateFields  []Field
	PublicFields   []Field
	PrivateMethods []Method
	PublicMethods  []Method
}

type Field struct {
	Common
	Struct *Struct
	Name   string
	Names  []string
	Type   Type
}

type Interface struct {
	Common
	Name string
}

type Import struct {
	Common
	Name string
}

type Method struct {
	Common
	Name string
}
