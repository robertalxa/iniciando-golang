package contas

import "banco/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso"
	}

	return "saldo insuficiente!"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito < 0 {
		return "Valor de depósito inválido", c.saldo
	}

	c.saldo += valorDeposito
	return "Depósito realizado com sucesso", c.saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && valorTransferencia > c.saldo {
		return false
	}

	c.Sacar(valorTransferencia)
	contaDestino.Depositar(valorTransferencia)
	return true
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
