package main

import (
	"03-02-funcoes_basico/calculadora"
	"fmt"
)

func main() {
	a := 10.0
	b := 5.0

	fmt.Println(
		"Soma =",
		calculadora.Calcular(calculadora.Somar, a, b)) // 10 + 5
	fmt.Println(
		"Subtração =",
		calculadora.Calcular(calculadora.Subtrair, a, b)) // 10 - 5
	fmt.Println(
		"Multiplicação =",
		calculadora.Calcular(calculadora.Multiplicar, a, b)) // 10 * 5
	fmt.Println(
		"Divisão =",
		calculadora.Calcular(calculadora.Dividir, a, b)) // 10 / 5
}
