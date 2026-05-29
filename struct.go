package main

import "fmt"

type Funcionario struct {
	Nome  string
	Email string
	CPF   int
}

type FuncionarioInternacional struct {
	Funcionario
	Pais string
}

func main() {

	funcionario1 := Funcionario{
		Nome:  "Wallace Freire",
		Email: "wallacehf12@gmail.com",
		CPF:   35196987061,
	}
	fmt.Println(funcionario1)

	funcionario2 := Funcionario{"Patrick", "patrick@contabil.com", 35196987061}
	fmt.Printf("Name: %s. Email: %s. CPF: %d\n", funcionario2.Nome, funcionario2.Email, funcionario2.CPF)

	funcionario3 := FuncionarioInternacional{
		Funcionario: Funcionario{
			Nome:  "Adema",
			Email: "adema@contabil.com",
			CPF:   74064231058,
		},
		Pais: "Escócia",
	}
	fmt.Printf("Name: %s. Email: %s. CPF: %d. País: %s\n", funcionario3.Nome, funcionario3.Email,
		funcionario3.CPF, funcionario3.Pais)
}
