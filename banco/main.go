package main

import (
	"banco/contas"
	"fmt"
)

func main() {
	conta := contas.ContaCorrente{
		Titular:       "Bieber",
		NumeroAgencia: 0001,
		NumeroConta:   2000,
		Saldo:         100,
	}

	fmt.Println(conta)
	fmt.Println(conta.Sacar(90))
	fmt.Println(conta)
	status, novoSaldo := conta.Depositar(900)
	fmt.Println(status, novoSaldo)
	fmt.Println(conta)

	contaAriana := contas.ContaCorrente{
		Titular:       "Ariana Grande",
		NumeroAgencia: 0001,
		NumeroConta:   1000,
		Saldo:         500,
	}

	fmt.Println(contaAriana)
	statusTransferencia := conta.Transferir(1100, &contaAriana)
	if statusTransferencia {
		fmt.Println("Transferencia realizada com sucesso")
	} else {
		fmt.Println("Transferencia não foi realizada")
	}
	fmt.Println(conta)
	fmt.Println(contaAriana)

}
