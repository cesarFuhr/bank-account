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

// New cria uma nova conta.
func New(id uint, owner string) *Account {
	a := Account{
		ID:      id,
		Owner:   owner,
		balance: 0,
	}

	return &a
}

// Saldo retorna o saldo atual da conta.
func (a *Account) Saldo() float64 {
	return a.balance
}

// Deposito incrementa o saldo com o valor depositado.
func (a *Account) Deposito(valor float64) {
	a.balance += valor

	op := Operation{
		Code:  OperationDeposito,
		Valor: valor,
	}

	a.operations = append(a.operations, op)
}

// Saque decrementa o saldo no valor sacado
func (a *Account) Saque(valor float64) {
	a.balance -= valor

	op := Operation{
		Code:  OperationSaque,
		Valor: valor,
	}

	a.operations = append(a.operations, op)
}

// Extrato retorna as operacoes realizadas na conta.
func (a *Account) Extrato() []Operation {
	return a.operations
}

// Operation é um registro de uma operacao que altere saldo na conta.
type Operation struct {
	Code  string
	Valor float64
}
