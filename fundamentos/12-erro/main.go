package main

import (
	"12-erro/model"
	"encoding/json"
	"fmt"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa da Alice"
	casa.X = 18
	casa.Y = 25
	if err := casa.SetValor(99999999999); err != nil {
		fmt.Println("[main] Houve um erro na atribuição de valor a casa: ", err.Error())
		if err == model.ErrValorDeImovelMuitoAlto {
			fmt.Println("Corretor, checar por favor sua avaliação.")
		}
		return
	}
	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("O valor da casa é: %d\r\n", casa.GetValor())
	objJSON, err := json.Marshal(casa)
	if err != nil {
		fmt.Println("[main] Houve um erro na geração do objeto JSON: ", err.Error())
		return
	}
	fmt.Println("A casa em JSON: ", string(objJSON))
}
