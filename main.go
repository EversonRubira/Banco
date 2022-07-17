package main

import (
	"fmt"
	 "Banco\contas"
	
)

func main() {
	contaDoEverson := contas.contaCorrente{Titular: "Everson Rubira", Saldo: 500}
	contaDaClaudia := contas.contaCorrente{Titular: "Claudia Teixeira", Saldo: 400}

	status := contaDoEverson.Transferir(200, &contaDaClaudia)

	fmt.Println(status)
	fmt.Println(contaDoEverson)
	fmt.Println(contaDaClaudia)

}
