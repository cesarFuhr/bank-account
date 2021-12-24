package repository_test

import (
	"log"
	"testing"

	"github.com/cesarFuhr/bank-account/internal/account"
	"github.com/cesarFuhr/bank-account/internal/repository"
	"github.com/matryer/is"
)

var db *repository.PersistentRepository

func TestMain(m *testing.M) {
	var err error
	db, err = repository.NewPersistentRepository()
	if err != nil {
		log.Fatal(err)
	}

	m.Run()
}

func TestPersistenteLer(t *testing.T) {
	t.Run("Deve ler um registro no banco de dados e retornar uma account completa", func(t *testing.T) {
		is := is.NewRelaxed(t)

		acc := account.Account{
			ID:    10,
			Owner: "Teste",
		}

		err := db.Escrever(acc)
		is.NoErr(err)

		got, err := db.Ler(acc.ID)
		is.NoErr(err)

		is.Equal(got.ID, acc.ID)
		is.Equal(got.Owner, acc.Owner)
	})
}
