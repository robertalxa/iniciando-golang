package main

import (
	"banco/clientes"
	"banco/contas"
	"fmt"
)

type verificarConta interface {
	Sacar(valor float64) string
}

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

func main() {
	titular := clientes.Titular{
		Nome:      "Bieber",
		CPF:       "123",
		Profissao: "Singer",
	}

	conta := contas.ContaCorrente{
		Titular:       titular,
		NumeroAgencia: 0001,
		NumeroConta:   2000,
	}

	conta.Depositar(100)
	conta.Sacar(90)
	conta.Depositar(900)

	titular2 := clientes.Titular{
		Nome:      "Ariana Grande",
		CPF:       "321",
		Profissao: "Singer",
	}

	contaAriana := contas.ContaCorrente{
		Titular:       titular2,
		NumeroAgencia: 0001,
		NumeroConta:   1000,
	}

	contaAriana.Depositar(500)

	statusTransferencia := conta.Transferir(1100, &contaAriana)
	if statusTransferencia {
		fmt.Println("Transferencia realizada com sucesso")
	} else {
		fmt.Println("Transferencia não foi realizada")
	}

	fmt.Println("Saldo da conta 2 =>", contaAriana.ObterSaldo())

	// ------------

	contaPoupa := contas.ContaPoupanca{}
	contaCorr := contas.ContaCorrente{}

	contaPoupa.Depositar(1000)
	contaCorr.Depositar(1000)

	PagarBoleto(&contaPoupa, 80) // Aqui o & passa o endereço de memória
	PagarBoleto(&contaCorr, 180) // Aqui o & passa o endereço de memória

	fmt.Println(contaPoupa.ObterSaldo())
	fmt.Println(contaCorr.ObterSaldo())
}
