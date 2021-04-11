package pointers_errors

import "errors"

var (
	ErroSaldoInsuficiente = errors.New("saldo insuficiente")
)

type Carteira struct {
	saldo Bitcoin
}

func (c *Carteira) Depositar(qtd Bitcoin) {
	c.saldo += qtd
}

func (c *Carteira) Saldo() Bitcoin {
	return c.saldo
}

func (c *Carteira) Retirar(qtd Bitcoin) error{
	if c.saldo < qtd {
		return ErroSaldoInsuficiente
	}

	c.saldo -= qtd
	return nil
}