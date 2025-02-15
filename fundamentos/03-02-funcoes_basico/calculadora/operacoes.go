package calculadora

import "fmt"

// Função que recebe uma operação e dois números
func Calcular(op func(float64, float64) float64, a, b float64) float64 {
	return op(a, b) // Chama a operação que foi passada
}

// Função para adição
func Somar(a, b float64) float64 {
	return a + b
}

// Função para subtração
func Subtrair(a, b float64) float64 {
	return a - b
}

// Função para multiplicação
func Multiplicar(a, b float64) float64 {
	return a * b
}

// Função para divisão
func Dividir(a, b float64) float64 {
	if b != 0 {
		return a / b
	}
	fmt.Println("Erro: Divisão por zero!")
	return 0
}
