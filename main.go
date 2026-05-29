package main

import "fmt"

type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {

	carro := Carro{
		Nome: "Classic",
	}

	carro.andar()

	resultado := func(x ...int) func() int {

		resultado := 0

		for _, value := range x {
			resultado += value
		}
		return func() int {
			return resultado * resultado
		}
	}

	fmt.Println(resultado(54, 54, 54, 54)())
}

func soma(primeiroNumero int, segundoNumero int) (result int) {
	result = primeiroNumero + segundoNumero
	return
}
