package repository

import (
	"errors"

	"github.com/cesarFuhr/bank-account/internal/account"
)

// ErrNotFound erro de conta nao encontrada.
var ErrNotFound = errors.New("conta nao encontrada")

// InMemoryRepository e um repositorio que guarda em
// memoria as contas
type InMemoryRepository struct {
	store map[uint]account.Account
}

// NewInMemoryRepository cria um novo repositorio
// em memoria.
func NewInMemoryRepository() *InMemoryRepository {
	m := make(map[uint]account.Account)

	return &InMemoryRepository{
		store: m,
	}
}

// Ler retorna a conta com o ID solicitado
func (r InMemoryRepository) Ler(id uint) (account.Account, error) {
	acc, ok := r.store[id]
	if !ok {
		return account.Account{}, ErrNotFound
	}

	return acc, nil
}

// Escrever escreve os dados da conta no repositorio.
func (r InMemoryRepository) Escrever(acc account.Account) error {
	r.store[acc.ID] = acc

	return nil
}
