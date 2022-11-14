package main

//Pointer Example

import "fmt"

type Carro struct {
	Marca string
}

//with * changes the OBJ in memory
func (c *Carro) MudaMarca(marca string) {
	c.Marca = marca
	fmt.Println(c.Marca)
}

func main() {
	carro := Carro{Marca: "Fiat"}
	carro.MudaMarca("Ford")
	fmt.Println(carro.Marca)
}