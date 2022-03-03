package main

import (
	"fmt"
	"github.com/badoux/checkmail"
	"modulo/auxiliar"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("herminhoo1@gmail.com")
	fmt.Println(erro)

	erro = checkmail.ValidateFormat("herminhoo1")
	fmt.Println(erro)
}
