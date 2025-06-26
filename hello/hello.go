package main

import "fmt"

func main() {
	var nome string = "Douglas"
	var versao float32 = 1.1
	var idade int = 24
	fmt.Println("Olá, sr.", nome, "! Sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
}

// Para buildar o executável é só rodar go build *nome do arquivo.go*
// Para não precisar ficar dando build toda hora é só rodar "go run arquivo.go"
