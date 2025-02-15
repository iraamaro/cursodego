package main

import (
	"02-pacotes/operadora"
	"02-pacotes/prefixo"
	"fmt"
)

// NomeDoUsuario é o nome do usuário do sistema
var NomeDoUsuario = "Alice"

func main() {
	fmt.Printf("Nome do usuŕio: %s\r\n", NomeDoUsuario)
	fmt.Printf("Prefixo da Capital: %d\r\n", prefixo.Capital)
	fmt.Printf("Prefixo da operadora: %s\r\n", operadora.NomeOperadora)
	fmt.Printf("Valor de teste: %s\r\n", prefixo.TesteComPrefixo)
}
