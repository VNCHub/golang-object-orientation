package contas

import (
	"fmt"
	"orientacao-a-objetos/titulares"
)

type ContaPoupanca struct {
	Titular       titulares.PessoaFisica
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}

func (c *ContaPoupanca) Sacar(valor float64) bool {
	podeSacar := valor <= c.saldo && valor > 0
	fmt.Println("Realizando saque no valor de: R$", valor, "...")
	if podeSacar {
		c.saldo -= valor
	}
	return podeSacar
}

func (c *ContaPoupanca) Depositar(valor float64) bool {
	podeDepositar := valor > 0
	fmt.Println("Realizando deposito no valor de: R$", valor, "...")
	if podeDepositar {
		c.saldo += valor
	}
	return podeDepositar
}
