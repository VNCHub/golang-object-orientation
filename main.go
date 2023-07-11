package main

import (
	"fmt"
	"orientacao-a-objetos/contas"
)

func main() {
	//Como declarar valores
	contaVinicius := contas.ContaCorrente{
		Titular: "Vinicius", NumeroAgencia: 123, NumeroConta: 001, Saldo: 100}

	//Como declarar ponteiros
	contaDenise := new(contas.ContaCorrente)
	contaDenise.Titular = "Denise"
	contaDenise.NumeroAgencia = 123
	contaDenise.NumeroConta = 002
	contaDenise.Saldo = 1000

	if contaDenise.Depositar(399.99) {
		fmt.Println("Deposito realizado com sucesso!")
	} else {
		fmt.Println("Não foi possível realizar o depósito, valor negativo.")
	}

	if contaDenise.Sacar(495.50) {
		fmt.Println("Saque realizado com sucesso!")
	} else {
		fmt.Println("Não foi possível realizar o saque, saldo insuficiente.")
	}

	if contaDenise.Transferir(200, &contaVinicius) {
		fmt.Println("Transferência realizada com sucesso!")
	} else {
		fmt.Println("Não foi possível realizar a transferência, saldo insuficiente.")
	}

	fmt.Println(contaVinicius, *contaDenise)
}
