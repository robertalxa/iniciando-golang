package main

import (
	"fmt"
	"os"
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
	// nome := "Robert"
	// fmt.Println("Olá, sr.", nome, "!")

	exibeIntroducao()
	exibeMenu()
	// fmt.Println("1 - Iniciar Monitoramento")
	// fmt.Println("2 - Exibir logs")
	// fmt.Println("0 - Sair do programa")

	// var comando int
	// fmt.Scanf("%d", &comando)
	// fmt.Scan(&comando) // Aqui eu não preciso passar o modificador da variável

	// fmt.Println("O comando escolhido foi:", comando)

	// if comando == 1 { // Tem que ser sempre uma condição que retorna um true ou false
	// 	fmt.Println("Monitorando")
	// } else if comando == 2 {
	// 	fmt.Println("Exibindo logs")
	// } else if comando == 0 {
	// 	fmt.Println("Saindo do programa")
	// } else {
	// 	fmt.Println("Não conheço este comando")
	// }

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando")
	case 2:
		fmt.Println("Exibindo logs")
	case 0:
		fmt.Println("Saindo do programa")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir logs")
	fmt.Println("0 - Sair do programa")
}

func exibeIntroducao() {
	nome := "Robert"
	fmt.Println("Olá, sr.", nome, "!")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) // Aqui eu não preciso passar o modificador da variável
	return comandoLido
}

// Para buildar o executável é só rodar go build *nome do arquivo.go*
// Para não precisar ficar dando build toda hora é só rodar "go run arquivo.go"
