package boilergen

var VERSION = "0.1.0"

type Boilergen struct {
	outputDir    string
	templatesDir string
}

func New() Boilergen {
	return Boilergen{}
}

func (b *Boilergen) ParsePackageDir(path string) error {
	return nil
}

func (b *Boilergen) SetOutputDirectory(path string)    { b.outputDir = path }
func (b *Boilergen) SetTemplatesDirectory(path string) { b.templatesDir = path }

func (b *Boilergen) Generate() error {
	return nil
}
