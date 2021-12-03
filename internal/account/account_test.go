package account_test

import (
	"testing"

	"github.com/cesarFuhr/bank-account/internal/account"
	"github.com/matryer/is"
)

func TestSaldo(t *testing.T) {
	b := account.NewBank()

	is := is.NewRelaxed(t)

	t.Run("retorna o saldo atual da conta", func(*testing.T) {
		id, err := b.AbrirConta("Teste")
		is.NoErr(err)

		want := 0.0

		got, err := b.Saldo(id)
		is.NoErr(err)

		is.Equal(got, want)
	})
}

func TestDeposito(t *testing.T) {
	b := account.NewBank()

	is := is.NewRelaxed(t)

	t.Run("incrementa a conta com o valor depositado", func(*testing.T) {
		id, err := b.AbrirConta("Teste")
		is.NoErr(err)

		want := 50.0
		err = b.Deposito(id, want)
		is.NoErr(err)

		got, err := b.Saldo(id)
		is.NoErr(err)

		is.Equal(got, want)
	})
	t.Run("incrementa em sequencia o saldo da conta", func(*testing.T) {
		id, err := b.AbrirConta("Teste")
		is.NoErr(err)

		deposito := 50.0
		want := deposito * 2

		err = b.Deposito(id, deposito)
		is.NoErr(err)
		err = b.Deposito(id, deposito)
		is.NoErr(err)

		got, err := b.Saldo(id)
		is.NoErr(err)

		is.Equal(got, want)
	})
	t.Run("incrementa valores diferentes no saldo da conta", func(*testing.T) {
		id, err := b.AbrirConta("Teste")
		is.NoErr(err)

		deposito1 := 50.0
		deposito2 := 30.0
		want := deposito1 + deposito2

		err = b.Deposito(id, deposito1)
		is.NoErr(err)
		err = b.Deposito(id, deposito2)
		is.NoErr(err)

		got, err := b.Saldo(id)
		is.NoErr(err)

		is.Equal(got, want)
	})
}

//
//func TestSaque(t *testing.T) {
//	t.Run("decrementa o valor do saldo", func(t *testing.T) {
//		conta := account.New(10, "Teste")
//		want := 0.0
//
//		saque := 50.0
//
//		conta.Deposito(saque)
//		conta.Saque(saque)
//
//		got := conta.Saldo()
//
//		if got != want {
//			t.Errorf("want %v,  got %v", want, got)
//		}
//	})
//	t.Run("decrementa em sequencia o valor do saldo", func(t *testing.T) {
//		conta := account.New(10, "Teste")
//		want := 0.0
//
//		saque := 50.0
//
//		conta.Deposito(saque * 2)
//		conta.Saque(saque)
//		conta.Saque(saque)
//
//		got := conta.Saldo()
//
//		if got != want {
//			t.Errorf("want %v,  got %v", want, got)
//		}
//	})
//	t.Run("decrementa diferentes valores do saldo", func(t *testing.T) {
//		conta := account.New(10, "Teste")
//		want := 20.0
//
//		saque := 50.0
//
//		conta.Deposito(saque * 2)
//		conta.Saque(saque)
//		conta.Saque(saque - want)
//
//		got := conta.Saldo()
//
//		if got != want {
//			t.Errorf("want %v,  got %v", want, got)
//		}
//	})
//}
//
//func TestExtrato(t *testing.T) {
//	t.Run("retorna as operacoes realizadas na conta", func(t *testing.T) {
//		conta := account.New(10, "Teste")
//		want := []account.Operation{
//			{
//				Code:  account.OperationDeposito,
//				Valor: 50.0,
//			},
//		}
//
//		conta.Deposito(50.0)
//
//		got := conta.Extrato()
//
//		if !reflect.DeepEqual(got, want) {
//			t.Errorf("want %v,  got %v", want, got)
//		}
//	})
//	t.Run("retorna diferentes operacoes", func(t *testing.T) {
//		conta := account.New(10, "Teste")
//		want := []account.Operation{
//			{
//				Code:  account.OperationDeposito,
//				Valor: 50.0,
//			},
//			{
//				Code:  account.OperationSaque,
//				Valor: 50.0,
//			},
//		}
//
//		conta.Deposito(50.0)
//		conta.Saque(50.0)
//
//		got := conta.Extrato()
//
//		if !reflect.DeepEqual(got, want) {
//			t.Errorf("want %v,  got %v", want, got)
//		}
//	})
//}
