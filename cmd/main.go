package main

import (
	"fmt"

	"github.com/cesarFuhr/bank-account/internal/account"
)

func main() {
	fmt.Println("bank-account")

	// Implementar um fluxo de criacao,
	// saldo, deposito e saque de uma
	// conta bancaria interagindo
	// com o usuario.

	fmt.Printf("Nome da conta: ")
	var accName string
	fmt.Scanln(&accName)

	acc := account.New(0, accName)
	fmt.Printf("%+v", acc)
}
