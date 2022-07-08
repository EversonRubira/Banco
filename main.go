package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func main() {
	contaDoEverson := ContaCorrente{titular: "Everson Rubira", numeroAgencia: 149, numeroConta: 123456, saldo: 500.0}

	contaDaClaudia := ContaCorrente{"Claudia Teixeira", 141, 234567, 600.0}

	fmt.Println(contaDoEverson)
	fmt.Println(contaDaClaudia)

}
