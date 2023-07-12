package contas

import (
	"orientacao-a-objetos/titulares"
)

type contaPoupanca struct {
	Titular       titulares.PessoaFisica
	NumeroAgencia int
	NumeroConta   int
	Operacao      int
	saldo         float64
}
