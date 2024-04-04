package attacks

import (
	"fmt"
	"math/big"

	"github.com/GerlachSnezka/glsz/rsa/factorization"
)

func modInverse(a *big.Int, m *big.Int) *big.Int {
	m0 := big.NewInt(0).Set(m)
	y := big.NewInt(0)
	x := big.NewInt(1)

	if m.Cmp(big.NewInt(1)) == 0 {
		return big.NewInt(0)
	}

	for a.Cmp(big.NewInt(1)) == 1 {
		q := big.NewInt(0).Div(a, m)
		t := big.NewInt(0)
		t.Set(m)
		m.Set(big.NewInt(0).Mod(a, m))
		a.Set(t)
		t.Set(y)
		y.Sub(x, big.NewInt(0).Mul(q, y))
		x.Set(t)
	}

	if x.Cmp(big.NewInt(0)) == -1 {
		x.Add(x, m0)
	}

	return x
}

func pow(x *big.Int, y *big.Int, p *big.Int) *big.Int {
	res := big.NewInt(1)
	x = big.NewInt(0).Mod(x, p)

	for y.Cmp(big.NewInt(0)) == 1 {
		if big.NewInt(0).Mod(y, big.NewInt(2)).Cmp(big.NewInt(1)) == 0 {
			res.Mul(res, x)
			res.Mod(res, p)
		}

		y.Div(y, big.NewInt(2))
		x.Mul(x, x)
		x.Mod(x, p)
	}

	return res
}

func decimalToHex(decimal *big.Int) string {
	return fmt.Sprintf("%x", decimal)
}

func (ctx *Attacks) IfaAttack(n *big.Int, e *big.Int, c *big.Int) {
	p, q := factorization.FactorDBFactorize(n)

	phi := big.NewInt(0)
	phi.Mul(big.NewInt(0).Sub(p, big.NewInt(1)), big.NewInt(0).Sub(q, big.NewInt(1)))

	d := modInverse(big.NewInt(0).Set(e), big.NewInt(0).Set(phi))
	decimal := pow(big.NewInt(0).Set(c), big.NewInt(0).Set(d), big.NewInt(0).Set(n))

	ctx.logger.Debug("", "p", p)
	ctx.logger.Debug("", "q", q)
	ctx.logger.Debug("", "phi", phi)
	ctx.logger.Debug("", "d", d)
	ctx.logger.Info("", "decimal", decimal)
	ctx.logger.Info("", "hex", decimalToHex(decimal))
	ctx.logger.Info(string(decimal.Bytes()))
}