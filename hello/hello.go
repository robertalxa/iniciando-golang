package main

import (
	"fmt"
)

func main() {
	// var nome = "Douglas"
	// var versao float32 = 1.1 // Sempre declarar o tipo dos floats pois existem dois
	// var idade = 24
	// variavelDeclaradaSemVar := "Se declara com :="
	// fmt.Println("Olá, sr.", nome, "! Sua idade é", idade)
	// fmt.Println("Este programa está na versão", versao)
	// fmt.Println("Como declara variável de maneira curta?", variavelDeclaradaSemVar)

	// fmt.Println("O tipo da variável nome é ", reflect.TypeOf(nome))
	// fmt.Println("O tipo da variável idade é ", reflect.TypeOf(idade))
	// fmt.Println("O tipo da variável versão é ", reflect.TypeOf(versao))

	// Iniciando de fato o programa
	nome := "Robert"
	fmt.Println("Olá, sr.", nome, "!")

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")

	var comando int
	// fmt.Scanf("%d", &comando)
	fmt.Scan(&comando) // Aqui eu não preciso passar o modificador da variável

	fmt.Println("O comando escolhido foi:", comando)
}

// Para buildar o executável é só rodar go build *nome do arquivo.go*
// Para não precisar ficar dando build toda hora é só rodar "go run arquivo.go"
