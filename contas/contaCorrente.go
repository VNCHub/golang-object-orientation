package contas

import (
	"fmt"
	"orientacao-a-objetos/titulares"
)

type ContaCorrente struct {
	Titular       titulares.PessoaFisica
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

// sintaxe de um método = func + (objeto) + nome da função + (parametros) + retorno
func (c *ContaCorrente) Sacar(valor float64) bool {
	podeSacar := valor <= c.saldo && valor > 0
	fmt.Println("Realizando saque no valor de: R$", valor, "...")
	if podeSacar {
		c.saldo -= valor
	}
	return podeSacar
}

func (c *ContaCorrente) Depositar(valor float64) bool {
	podeDepositar := valor > 0
	fmt.Println("Realizando deposito no valor de: R$", valor, "...")
	if podeDepositar {
		c.saldo += valor
	}
	return podeDepositar
}

func (c *ContaCorrente) Transferir(valor float64, contaDestino *ContaCorrente) bool {
	podeTransferir := valor > 0 && valor < c.saldo
	fmt.Println("Realizando transferência no valor de: R$", valor, "...")
	if podeTransferir {
		c.saldo -= valor
		contaDestino.saldo += valor
	}
	return podeTransferir
}
