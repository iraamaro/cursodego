package matematica

/*
Calculo executa qualquer tipo de calculo basta enviar a função desejada
*/
func Calculo(funcao func(int, int) int, numeroA int, numeroB int) int {
	return funcao(numeroA, numeroB)
}

// Produto faz a multiplicação
func Produto(x int, y int) int {
	return x * y
}

// Divisor faz a divisão
func Divisor(x int, y int) (resultado int) {
	if y != 0 {
		resultado = x / y
		return
	}
	return 0
}

// DivisorComResto faz a divisão e retorna resultado e resto
func DivisorComresto(x int, y int) (resultado int, resto int) {
	if y != 0 {
		resultado = x / y
		resto = x % y
		return
	}
	return 0, 0
}
