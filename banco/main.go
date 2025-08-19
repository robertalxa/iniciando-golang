package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente!"
}

func (c *ContaCorrente) depositar(valorDeposito float64) (string, float64) {
	if valorDeposito < 0 {
		return "Valor de depósito inválido", c.saldo
	}

	c.saldo += valorDeposito
	return "Depósito realizado com sucesso", c.saldo
}

func main() {
	conta := ContaCorrente{
		titular:       "Robert",
		numeroAgencia: 0001,
		numeroConta:   2000,
		saldo:         100,
	}

	fmt.Println(conta)
	fmt.Println(conta.sacar(90))
	fmt.Println(conta)
	status, novoSaldo := conta.depositar(900)
	fmt.Println(status, novoSaldo)
	fmt.Println(conta)
}
