package rsa

import (
	"fmt"
	"math/big"

	"github.com/GerlachSnezka/glsz/rsa/attacks"
	"github.com/charmbracelet/log"
	"github.com/urfave/cli/v2"
)

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
		{
			Name: "ifam",
			Aliases: []string{"im"},
			Usage: "Perform Integer Factorization Attack on multiple n', e', c' values",
			Action: IFAM,
			ArgsUsage: "n1 e1 c1 n2 e2 c2 ...",
		},
	}
}

func IFA(ctx *cli.Context) error {
	argN := ctx.String("n")
	argE := ctx.String("e")
	argC := ctx.String("c")
	verbose := ctx.Bool("verbose")

	n, _ := big.NewInt(0).SetString(argN, 10)
	e, _ := big.NewInt(0).SetString(argE, 10)
	c, _ := big.NewInt(0).SetString(argC, 10)

	ifa := attacks.NewIfa(verbose)
	
	ifa.Print(ifa.Attack(n, e, c))

	return nil
}

func IFAM(ctx *cli.Context) error {
	args := ctx.Args()
	if args.Len() % 3 != 0 {
		log.Fatalf("Invalid number of arguments: %d", args.Len())
	}

	ifa := attacks.NewIfa(true)

	var message string

	for i := 0; i < args.Len(); i += 3 {
		argN := args.Get(i)
		argE := args.Get(i + 1)
		argC := args.Get(i + 2)

		log.Info("Attacking", "pair", i / 3 + 1)

		n, _ := big.NewInt(0).SetString(argN, 10)
		e, _ := big.NewInt(0).SetString(argE, 10)
		c, _ := big.NewInt(0).SetString(argC, 10)

		p, q, phi, d, decimal := ifa.Attack(n, e, c)
		ifa.Print(p, q, phi, d, decimal)

		message += string(decimal.Bytes())

		fmt.Println()
	}

	log.Info("", "final_str", message)

	return nil
}