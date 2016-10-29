package asttree

type Field struct {
	Package *Package

	Name  string
	Names []string

	Struct *Struct
	Type   *Type
}

func NewField() *Field {
	return &Field{
		Names: make([]string, 0),
	}
}
