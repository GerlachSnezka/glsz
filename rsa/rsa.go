package rsa

import (
	"log"
	"math/big"

	"github.com/GerlachSnezka/glsz/rsa/attacks"
	"github.com/urfave/cli/v2"
)

type Rsa struct {
	log *log.Logger
}

func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name: "ifa",
			Aliases: []string{"i"},
			Usage: "Perform Integer Factorization Attack",
			Action: IFA,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name: "n",
					Usage: "Modulus",
					Required: true,
				},
				&cli.StringFlag{
					Name: "e",
					Usage: "Public exponent",
					Required: true,
				},
				&cli.StringFlag{
					Name: "c",
					Usage: "Ciphertext",
					Required: true,
				},
				&cli.BoolFlag{
					Name: "verbose",
					Usage: "Verbose output",
					Value: false,
				},
			},
		},
	}
}

func IFA(ctx *cli.Context) error {
	argN := ctx.String("n")
	argE := ctx.String("e")
	argC := ctx.String("c")

	n, _ := big.NewInt(0).SetString(argN, 10)
	e, _ := big.NewInt(0).SetString(argE, 10)
	c, _ := big.NewInt(0).SetString(argC, 10)

	attacks.New(ctx.Bool("verbose")).IfaAttack(n, e, c)

	return nil
}