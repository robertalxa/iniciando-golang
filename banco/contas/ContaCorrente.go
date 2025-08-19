package contas

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorSaque
		return "Saque realizado com sucesso"
	}

	return "Saldo insuficiente!"
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito < 0 {
		return "Valor de depósito inválido", c.Saldo
	}

	c.Saldo += valorDeposito
	return "Depósito realizado com sucesso", c.Saldo
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorTransferencia > 0 && valorTransferencia > c.Saldo {
		return false
	}

	c.Sacar(valorTransferencia)
	contaDestino.Depositar(valorTransferencia)
	return true
}
