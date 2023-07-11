package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	Saldo         float64
}

func main() {
	//Como declarar valores
	contaVinicius := ContaCorrente{"Vinicius", 123, 001, 100}

	//Como declarar ponteiros
	contaDenise := new(ContaCorrente)
	contaDenise.titular = "Denise"
	contaDenise.numeroAgencia = 123
	contaDenise.numeroConta = 002
	contaDenise.Saldo = 200

	fmt.Println(contaVinicius, *contaDenise)
}
