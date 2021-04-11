package pointers_errors

import (
	"testing"
)

func confirmaSaldo(t *testing.T, carteira Carteira, esperado Bitcoin) {
	t.Helper()
	resultado := carteira.Saldo()

	if resultado != esperado {
		t.Errorf("resultado %s, esperaod %s", resultado, esperado)
	}
}

func confirmErro (t *testing.T, resultado error, esperado error) {
	t.Helper()
	if resultado == nil {
		t.Error("esperava um ruim aqui, mas fiquei no vácuo")
	}

	if resultado != esperado {
		t.Errorf("resultado %s, esperaod %s", resultado, esperado)
	}
}

func confirmNoError (t *testing.T, resultado error) {
	t.Helper()
	if resultado != nil {
		t.Errorf("recebemos um erro inesperado aqui hein?!")
	}
}

func TestCarteira(t *testing.T) {


	t.Run("Depositar", func(t *testing.T) {
		carteira := Carteira{}

		carteira.Depositar(Bitcoin(10))
		confirmaSaldo(t, carteira, Bitcoin(10))
	})

	t.Run("Retirar com saldo suficiente", func(t *testing.T) {
		carteira := Carteira{}
		carteira.Depositar(Bitcoin(20))

		err := carteira.Retirar(Bitcoin(10))

		confirmaSaldo(t, carteira, Bitcoin(10))
		confirmNoError(t, err)
	})

	t.Run("Retirar com saldo insuficiente", func(t *testing.T) {
		saldoInicial := Bitcoin(20)
		carteira := Carteira{}
		carteira.Depositar(saldoInicial)
		erro := carteira.Retirar(Bitcoin(100))

		confirmaSaldo(t, carteira, saldoInicial)

		confirmErro(t, erro, ErroSaldoInsuficiente)
	})

}
