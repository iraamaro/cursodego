package main

import (
	"07-structs_avancado/model"
	"encoding/json"
	"fmt"
)

func main() {
	casa := model.Imovel{}
	casa.Nome = "Casa Azul"
	casa.X = 18
	casa.Y = 25
	casa.SetValor(600000)
	fmt.Printf("Casa é: %+v\r\n", casa)
	fmt.Printf("O valor da casa é: %d\r\n", casa.GetValor())
	objJSON, _ := json.Marshal(casa)
	fmt.Println("A casa em JSON: ", string(objJSON))
}
