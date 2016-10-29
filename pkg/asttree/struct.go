package asttree

type Struct struct {
	Package *Package

	Name       string
	Value      string
	ValueExact string

	Types          []*Type
	PrivateFields  []*Field
	PublicFields   []*Field
	PrivateMethods []*Method
	PublicMethods  []*Method
}

func NewStruct() *Struct {
	return &Struct{
		Types:          make([]*Type, 0),
		PrivateFields:  make([]*Field, 0),
		PublicFields:   make([]*Field, 0),
		PrivateMethods: make([]*Method, 0),
		PublicMethods:  make([]*Method, 0),
	}
}
