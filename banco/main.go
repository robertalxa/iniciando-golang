package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorSaque float64) string {
	podeSacar := valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente!"
}

func main() {
	conta := ContaCorrente{
		titular:       "Robert",
		numeroAgencia: 0001,
		numeroConta:   2000,
		saldo:         100,
	}

	fmt.Println(conta)
	fmt.Println(conta.sacar(900))
	fmt.Println(conta)
}
