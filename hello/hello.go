package main

import (
	"fmt"
	"reflect"
)

func main() {
	var nome = "Douglas"
	var versao float32 = 1.1 // Sempre declarar o tipo dos floats pois existem dois
	var idade = 24
	variavelDeclaradaSemVar := "Se declara com :="
	fmt.Println("Olá, sr.", nome, "! Sua idade é", idade)
	fmt.Println("Este programa está na versão", versao)
	fmt.Println("Como declara variável de maneira curta?", variavelDeclaradaSemVar)

	fmt.Println("O tipo da variável nome é ", reflect.TypeOf(nome))
	fmt.Println("O tipo da variável idade é ", reflect.TypeOf(idade))
	fmt.Println("O tipo da variável versão é ", reflect.TypeOf(versao))
}

// Para buildar o executável é só rodar go build *nome do arquivo.go*
// Para não precisar ficar dando build toda hora é só rodar "go run arquivo.go"
