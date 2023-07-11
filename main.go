package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	Saldo         float64
}

func main() {
	contaVinicius := ContaCorrente{"Vinicius", 123, 001, 100}

	fmt.Println(contaVinicius)
}
