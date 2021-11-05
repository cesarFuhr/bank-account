package account_test

import (
	"reflect"
	"testing"

	"github.com/cesarFuhr/bank-account/internal/account"
)

func TestSaldo(t *testing.T) {
	t.Run("retorna o saldo atual da conta", func(t *testing.T) {
		conta := account.New(10, "Teste")
		want := 0.0

		got := conta.Saldo()

		if got != want {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
}

func TestDeposito(t *testing.T) {
	t.Run("incrementa a conta com o valor depositado", func(t *testing.T) {
		conta := account.New(10, "Teste")
		want := 50.0

		conta.Deposito(want)

		got := conta.Saldo()

		if got != want {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
	t.Run("incrementa em sequencia o saldo da conta", func(t *testing.T) {
		conta := account.New(10, "Teste")
		deposito := 50.0
		want := deposito * 2

		conta.Deposito(deposito)
		conta.Deposito(deposito)

		got := conta.Saldo()

		if got != want {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
	t.Run("incrementa valores diferentes no saldo da conta", func(t *testing.T) {
		conta := account.New(10, "Teste")
		deposito1 := 50.0
		deposito2 := 30.0
		want := deposito1 + deposito2

		conta.Deposito(deposito1)
		conta.Deposito(deposito2)

		got := conta.Saldo()

		if got != want {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
}

func TestSaque(t *testing.T) {
	t.Run("decrementa o valor do saldo", func(t *testing.T) {
		conta := account.New(10, "Teste")
		want := 0.0

		saque := 50.0

		conta.Deposito(saque)
		conta.Saque(saque)

		got := conta.Saldo()

		if got != want {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
	t.Run("decrementa em sequencia o valor do saldo", func(t *testing.T) {
		conta := account.New(10, "Teste")
		want := 0.0

		saque := 50.0

		conta.Deposito(saque * 2)
		conta.Saque(saque)
		conta.Saque(saque)

		got := conta.Saldo()

		if got != want {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
	t.Run("decrementa diferentes valores do saldo", func(t *testing.T) {
		conta := account.New(10, "Teste")
		want := 20.0

		saque := 50.0

		conta.Deposito(saque * 2)
		conta.Saque(saque)
		conta.Saque(saque - want)

		got := conta.Saldo()

		if got != want {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
}

func TestExtrato(t *testing.T) {
	t.Run("retorna as operacoes realizadas na conta", func(t *testing.T) {
		conta := account.New(10, "Teste")
		want := []account.Operation{
			{
				Code:  account.OperationDeposito,
				Valor: 50.0,
			},
		}

		conta.Deposito(50.0)

		got := conta.Extrato()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
	t.Run("retorna diferentes operacoes", func(t *testing.T) {
		conta := account.New(10, "Teste")
		want := []account.Operation{
			{
				Code:  account.OperationDeposito,
				Valor: 50.0,
			},
			{
				Code:  account.OperationSaque,
				Valor: 50.0,
			},
		}

		conta.Deposito(50.0)
		conta.Saque(50.0)

		got := conta.Extrato()

		if !reflect.DeepEqual(got, want) {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
}
