package main

import (
	"11-mapas/model"
	"fmt"
)

var cache map[string]model.Imovel

func main() {
	cache = make(map[string]model.Imovel, 0)
	casa := model.Imovel{}
	casa.Nome = "Casa Amarela"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(600000)

	cache["Casa Amarela"] = casa
	fmt.Println("Há uma casa amarela no cache?")
	imovel, achou := cache["Casa Amarela"]
	if achou {
		fmt.Printf("Sim, e o que achei foi: %+v\r\n", imovel)
	}

	apt := model.Imovel{}
	apt.Nome = "Apartamento Azul"
	apt.X = 38
	apt.Y = 66
	apt.SetValor(180000)

	cache[apt.Nome] = apt

	fmt.Println("Quantos itens há no cache?", len(cache))

	for chave, imovel := range cache {
		fmt.Printf("Chave[%s] = %+v\n\r", chave, imovel)
	}

	imovel, achou = cache["Casa Amarela"]
	if achou {
		delete(cache, imovel.Nome)
	}
	fmt.Println("Quantos itens há no cache?", len(cache))

	_, achou = cache["Casa Azul"]
	if achou {
		delete(cache, "Casa Azul")
	}
	fmt.Println("Quantos itens há no cache?", len(cache))
}
