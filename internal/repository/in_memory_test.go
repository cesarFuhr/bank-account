package repository_test

import (
	"testing"

	"github.com/cesarFuhr/bank-account/internal/account"
	"github.com/cesarFuhr/bank-account/internal/repository"
)

func TestLer(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	repo.Escrever(account.Account{})

	t.Run("retorna os dados da conta solicitada", func(t *testing.T) {
		want := uint(0)

		got, err := repo.Ler(0)

		if err != nil {
			t.Errorf("was not expecting an error")
		}

		if got.ID != want {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
}

func TestEscrever(t *testing.T) {
	repo := repository.NewInMemoryRepository()
	t.Run("escreve dados na conta solicitada", func(t *testing.T) {
		want := "Teste"

		err := repo.Escrever(account.Account{
			ID:    1,
			Owner: want,
		})

		if err != nil {
			t.Errorf("was not expecting an error")
			return
		}

		acc, err := repo.Ler(1)
		if err != nil {
			t.Errorf("was not expecting an error")
			return
		}
		got := acc.Owner

		if got != want {
			t.Errorf("want %v,  got %v", want, got)
		}
	})
}
