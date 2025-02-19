package main

import "fmt"

// Imovel é a struct que armazena dados de um imovel
type Imovel struct {
	X     int
	Y     int
	Nome  string
	valor int
}

func main() {
	casa := new(Imovel)
	casa.valor = 0
	fmt.Printf("Casa é: %p - %+v\n", &casa, casa)
	chacara := Imovel{17, 28, "Chácara", 280000}
	fmt.Printf("Chacára é: %p - %+v\n", &chacara, chacara)
	mudaImovel(&chacara)
	fmt.Printf("Chacára é: %p - %+v\n", &chacara, chacara)
}

func mudaImovel(imovel *Imovel) {
	imovel.X = imovel.X + 10
	imovel.Y = imovel.Y - 5
}
