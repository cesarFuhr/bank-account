package main

import (
	"fmt"
	"strconv"

	"github.com/cesarFuhr/bank-account/internal/account"
	"github.com/cesarFuhr/bank-account/internal/repository"
)

const (
	ComandoSair = iota
	ComandoCriar
	ComandoUsarConta
	ComandoSaldo
	ComandoDeposito
	ComandoSaque
	ComandoExtrato
)

func main() {
	// Implementar um fluxo de criacao,
	// saldo, deposito e saque de uma
	// conta bancaria interagindo
	// com o usuario.
	repo := repository.NewInMemoryRepository()
	bank := account.NewBank(repo)

	var sair bool
	var contaAtual uint

	fmt.Println("Bem vindo ao banco da confraria.")
	for !sair {
		fmt.Println(`
Selecione um dos comandos abaixo para interagir com nossos servicos:
` + strconv.Itoa(ComandoSair) + ` - Sair
` + strconv.Itoa(ComandoCriar) + ` - Criar conta
` + strconv.Itoa(ComandoUsarConta) + ` - Usar conta
` + strconv.Itoa(ComandoSaldo) + ` - Saldo
` + strconv.Itoa(ComandoDeposito) + ` - Deposito
` + strconv.Itoa(ComandoSaque) + ` - Saque
` + strconv.Itoa(ComandoExtrato) + ` - Extrato
`)
		fmt.Printf("Comando: ")
		var comando int
		fmt.Scanln(&comando)

		switch comando {
		case ComandoSair:
			sair = true
		case ComandoCriar:
			fmt.Printf("Nome da conta: ")
			var accName string
			fmt.Scanln(&accName)

			acc, err := bank.AbrirConta(accName)
			if err != nil {
				fmt.Println(err)
				continue
			}
			contaAtual = acc
		case ComandoUsarConta:
			fmt.Printf("Numero da conta: ")
			var acc uint
			fmt.Scanln(&acc)

			contaAtual = acc
		case ComandoSaldo:
			saldo, err := bank.Saldo(contaAtual)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Printf("Saldo: R$ %.2f\n", saldo)
		case ComandoDeposito:
			fmt.Printf("Valor para deposito: ")
			var valor float64
			fmt.Scanln(&valor)

			err := bank.Deposito(contaAtual, valor)
			if err != nil {
				fmt.Println(err)
			}
		case ComandoSaque:
			fmt.Printf("Valor para saque: ")
			var valor float64
			fmt.Scanln(&valor)

			err := bank.Saque(contaAtual, valor)
			if err != nil {
				fmt.Println(err)
			}
		case ComandoExtrato:
			extrato, err := bank.Extrato(contaAtual)
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println("Extrato:")
			for _, op := range extrato {
				fmt.Printf("R$ %2.2f \t%s\n", op.Valor, op.Code)
			}
		default:
			fmt.Println("Comando invalido, por favor escolha uma das opcoes listadas.")
		}
	}

	fmt.Println("Obrigado por utilizar o banco da confraria...")
}
