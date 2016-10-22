package main

import (
	"log"
	"os"

	"github.com/moul/boilergen/pkg/boilergen"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "boilergen"
	app.Author = "Manfred Touron"
	app.Email = "https://github.com/moul/boilergen"
	app.Version = boilergen.VERSION
	app.Usage = app.Name

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "t,templates-directory",
			Usage: "templates directory",
			Value: ".",
		},
		cli.StringFlag{
			Name:  "o,output-directory",
			Usage: "output directory",
			Value: ".",
		},
	}

	app.Action = generate

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Runtime error: %v", err)
	}
}

func generate(c *cli.Context) error {
	directory := "."
	if len(c.Args()) > 0 {
		directory = c.Args()[0]
	}

	boiler := boilergen.New()
	boiler.SetOutputDirectory(c.String("output-directory"))
	boiler.SetTemplatesDirectory(c.String("templates-directory"))
	if err := boiler.ParsePackageDir(directory); err != nil {
		log.Fatalf("Parse error: %v", err)
	}
	if err := boiler.Generate(); err != nil {
		log.Fatalf("Generate error: %v", err)
	}
	return nil
}
