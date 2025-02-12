// Main é package
package main

import "fmt"

var (
	// Endereco é valor importante
	Endereco string
	// Telefone também é importante
	Telefone            string
	quantidade, estoque int     // quantidade = 0
	comprou             bool    // comprou = false
	valor               float64 // valor = 0.00
	palavras            rune
)

func main() {
	teste := "Valor de teste"
	fmt.Printf("endereço: %s\r\n", Endereco)
	fmt.Printf("teste: %s\r\n", teste)
	fmt.Printf("telefone: %s\r\n", Telefone)
	fmt.Printf("quantidade: %d\r\n", quantidade)
	fmt.Printf("estoque: %d\r\n", estoque)
	fmt.Printf("comprou: %v\r\n", comprou)
	fmt.Printf("valor: %v\r\n", valor)
	fmt.Printf("palavras: %v\r\n", palavras)
}
