package account

import "errors"

// ErrNotFound erro de conta nao encontrada.
var ErrNotFound = errors.New("conta nao encontrada")

// InMemoryRepository e um repositorio que guarda em
// memoria as contas
type InMemoryRepository struct {
	store map[uint]Account
}

// NewInMemoryRepository cria um novo repositorio
// em memoria.
func NewInMemoryRepository() *InMemoryRepository {
	m := make(map[uint]Account)

	return &InMemoryRepository{
		store: m,
	}
}

// Ler retorna a conta com o ID solicitado
func (r InMemoryRepository) Ler(id uint) (Account, error) {
	acc, ok := r.store[id]
	if !ok {
		return Account{}, ErrNotFound
	}

	return acc, nil
}

// Escrever escreve os dados da conta no repositorio.
func (r InMemoryRepository) Escrever(acc Account) error {
	r.store[acc.ID] = acc

	return nil
}
