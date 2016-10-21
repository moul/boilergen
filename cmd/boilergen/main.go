package main

import (
	"fmt"
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

	app.Action = generate

	if err := app.Run(os.Args); err != nil {
		log.Fatalf("Runtime error: %v", err)
	}
}

func generate(c *cli.Context) error {
	fmt.Println("Hello world")
	return nil
}
