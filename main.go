package main

import (
	"log"
	"os"

	"github.com/GerlachSnezka/glsz/rsa"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "rsa",
				Aliases: []string{"r"},
				Usage: "Useful RSA attacks",
				Subcommands: rsa.Commands(),
			},
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}