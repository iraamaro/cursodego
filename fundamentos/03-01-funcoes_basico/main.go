package main

import (
	"03-funcoes_basico/matematica"
	"fmt"
)

func main() {
	resultado := matematica.Calculo(matematica.Produto, 2, 2)
	fmt.Printf("Produto = %d\r\n", resultado)
	resultado = matematica.Soma(3, 3)
	fmt.Printf("Soma = %d\r\n", resultado)
	resultado = matematica.Divisor(12, 3)
	fmt.Printf("Divisão = %d\r\n", resultado)
	resultado, resto := matematica.DivisorComresto(12, 5)
	fmt.Printf("Divisão = %d - Resto = %d\r\n", resultado, resto)
}
