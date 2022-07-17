package contas

type contaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *contaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque Realizado com Sucesso"
	} else {
		return "Saldo Isuficiente"
	}
}

func (c *contaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Deposito Realizado com Sucesso", c.Saldo
	} else {
		return "Valor Deposito menor que zero", c.Saldo
	}
}

func (c *contaCorrente) Transferir(valorDeTransferencia float64, contaDestino *contaCorrente) bool {
	if valorDeTransferencia < c.Saldo && valorDeTransferencia > 0 {
		c.Saldo -= valorDeTransferencia
		contaDestino.Depositar(valorDeTransferencia)
		return true

	} else {
		return false
	}

}