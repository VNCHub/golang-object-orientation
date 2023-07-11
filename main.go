package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

// sintaxe de um método = func + (objeto) + nome da função + (parametros) + retorno
func (c *ContaCorrente) sacar(valor float64) bool {
	podeSacar := valor <= c.saldo && valor > 0
	fmt.Println("Realizando saque no valor de: R$", valor, "...")
	if podeSacar {
		c.saldo -= valor
	}
	return podeSacar
}

func (c *ContaCorrente) depositar(valor float64) bool {
	podeDepositar := valor > 0
	fmt.Println("Realizando deposito no valor de: R$", valor, "...")
	if podeDepositar {
		c.saldo += valor
	}
	return podeDepositar
}

func main() {
	//Como declarar valores
	contaVinicius := ContaCorrente{"Vinicius", 123, 001, 100}

	//Como declarar ponteiros
	contaDenise := new(ContaCorrente)
	contaDenise.titular = "Denise"
	contaDenise.numeroAgencia = 123
	contaDenise.numeroConta = 002
	contaDenise.saldo = 1000

	if contaDenise.depositar(399.99) {
		fmt.Println("Deposito realizado com sucesso!")
	} else {
		fmt.Println("Não foi possível realizar o depósito, valor negativo.")
	}

	if contaDenise.sacar(495.50) {
		fmt.Println("Saque realizado com sucesso!")
	} else {
		fmt.Println("Não foi possível realizar o saque, saldo insuficiente.")
	}

	fmt.Println(contaVinicius, *contaDenise)
}
