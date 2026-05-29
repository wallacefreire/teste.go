package main

import "fmt"

type Carro struct {
	Name string
}

func (c *Carro) andou() {
	c.Name = "Maserati Granturismo"
	fmt.Println(c.Name, "andou!")
}

func main() {
	carro := Carro{
		Name: "Cruze LTZ",
	}

	carro.andou()
	fmt.Println(carro.Name)
}
