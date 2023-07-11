package contas

import "fmt"

type ContaCorrente struct {
	Titular       string
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

// sintaxe de um método = func + (objeto) + nome da função + (parametros) + retorno
func (c *ContaCorrente) Sacar(valor float64) bool {
	podeSacar := valor <= c.Saldo && valor > 0
	fmt.Println("Realizando saque no valor de: R$", valor, "...")
	if podeSacar {
		c.Saldo -= valor
	}
	return podeSacar
}

func (c *ContaCorrente) Depositar(valor float64) bool {
	podeDepositar := valor > 0
	fmt.Println("Realizando deposito no valor de: R$", valor, "...")
	if podeDepositar {
		c.Saldo += valor
	}
	return podeDepositar
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	podeTransferir := valor > 0 && valor < c.Saldo
	fmt.Println("Realizando transferência no valor de: R$", valor, "...")
	if podeTransferir {
		c.Saldo -= valor
		contaDestino.Saldo += valor
	}
	return podeTransferir
}