package account

const (
	// OperationDeposito é o código da operacao deposito.
	OperationDeposito = "DEPOSITO"

	// OperationSaque é o código da operacao saque.
	OperationSaque = "SAQUE"
)

// Account é uma conta bancária.
type Account struct {
	ID         uint
	Owner      string
	balance    float64
	operations []Operation
}

// Bank e um conjunto de contas.
type Bank struct {
	repo             *InMemoryRepository
	numberOfAccounts uint
}

// NewBank cria um novo banco e retorna um ponteiro
// para ele.
func NewBank() *Bank {
	repo := NewInMemoryRepository()

	return &Bank{
		repo:             repo,
		numberOfAccounts: 0,
	}
}

// AbrirConta cria uma nova conta no banco.
func (b *Bank) AbrirConta(nome string) (uint, error) {
	acc := Account{
		ID:    b.numberOfAccounts,
		Owner: nome,
	}

	if err := b.repo.Escrever(acc); err != nil {
		return 0, err
	}

	b.numberOfAccounts++

	return acc.ID, nil
}

// Saldo retorna o saldo atual da conta.
func (b *Bank) Saldo(id uint) (float64, error) {
	acc, err := b.repo.Ler(id)
	if err != nil {
		return 0, err
	}

	return acc.balance, nil
}

// Deposito incrementa o saldo com o valor depositado.
func (b *Bank) Deposito(accID uint, valor float64) error {
	acc, err := b.repo.Ler(accID)
	if err != nil {
		return err
	}

	acc.balance += valor

	op := Operation{
		Code:  OperationDeposito,
		Valor: valor,
	}

	acc.operations = append(acc.operations, op)

	if err := b.repo.Escrever(acc); err != nil {
		return err
	}

	return nil
}

// Saque decrementa o saldo no valor sacado
func (b *Bank) Saque(id uint, valor float64) error {
	acc, err := b.repo.Ler(id)
	if err != nil {
		return err
	}

	acc.balance -= valor

	op := Operation{
		Code:  OperationSaque,
		Valor: valor,
	}

	acc.operations = append(acc.operations, op)

	if err := b.repo.Escrever(acc); err != nil {
		return err
	}

	return nil
}

// Extrato retorna as operacoes realizadas na conta.
func (b *Bank) Extrato(id uint) ([]Operation, error) {
	acc, err := b.repo.Ler(id)
	if err != nil {
		return nil, err
	}

	return acc.operations, nil
}

// Operation é um registro de uma operacao que altere saldo na conta.
type Operation struct {
	Code  string
	Valor float64
}
