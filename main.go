package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Conta Comigo - primeiro programa em GO!")
}

func generateID() {
	id := uuid.New()
	fmt.Println("ID gerado:", id)
}

/*
	O GO não vai compilar caso você crie uma função e não a utilize, por isso, a função generateID() não é chamada no main()
	e o código compila normalmente. Se você tentar chamar a função generateID() dentro do main(), o código irá compilar e rodar,
	gerando um ID único cada vez que for executado.
*/
