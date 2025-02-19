package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	cliente := &http.Client{
		Timeout: time.Second * 30,
	}

	resposta, err := cliente.Get("https://fooapi.com/api/products")
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina fooapi.com. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := io.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da pagina fooapi.com. Erro: ", err.Error())
			return
		}
		fmt.Println(string(corpo))
	}

	request, err := http.NewRequest("GET", "https://fooapi.com/api/products", nil)
	if err != nil {
		fmt.Println("[main] Erro ao criar um request para o https://fooapi.com/api/products. Erro: ", err.Error())
		return
	}
	request.SetBasicAuth("teste", "teste")
	resposta, err = cliente.Do(request)
	if err != nil {
		fmt.Println("[main] Erro ao abrir a pagina do https://fooapi.com/api/products. Erro: ", err.Error())
		return
	}
	defer resposta.Body.Close()

	if resposta.StatusCode == 200 {
		corpo, err := io.ReadAll(resposta.Body)
		if err != nil {
			fmt.Println("[main] Erro ao ler o conteudo da pagina do https://fooapi.com/api/products. Erro: ", err.Error())
			return
		}
		fmt.Println(" ")
		fmt.Println(string(corpo))
	}
}
