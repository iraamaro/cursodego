package main

import "fmt"

func main() {
	meses := 0
	situacao := true
	cidade := "Teste"

	if meses <= 6 {
		fmt.Println("Esse credor deve há pouco tempo.")
	}

	if situacao {
		fmt.Println("Ele está devendo.")
	}

	if !situacao {
		fmt.Println("Ele está adimplente.")
	}

	if cidade == "Teste" {
		fmt.Println("O cliente vive no estado Teste.")
	}

	if descricao, status := haQuantoTempoEstaDevendo(meses); status {
		fmt.Println("Qual a situação do cliente?", descricao)
		return
	}

	// A instrução abaixo não compila, explicação escopo de descricao e status exceção
	// fmt.Println("Qual o status?", descricao)

	fmt.Println("Obrigado por nos consultar.")
}

func haQuantoTempoEstaDevendo(meses int) (descricao string, status bool) {
	if meses > 0 {
		status = true
		descricao = "O cliente está devendo."
		return
	}
	descricao = "O cliente está em dia nos pagamentos."
	return
}
