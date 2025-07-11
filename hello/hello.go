package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// _, idade := devolveNomeEIdade() Esse underline faz o go ignorar o primeiro retorno
	// fmt.Println("tenho ", idade, " anos")
	exibeIntroducao()
	exibeMenu()
	comando := leComando()

	switch comando {
	case 1:
		iniciarMonitoramento()
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando")
	site := "https://www.alura.com.br"
	response, _ := http.Get(site)
	fmt.Println("Retorno da requisição: ", response)

}

func devolveNomeEIdade() (string, int) {
	nome := "Robert"
	idade := 24
	return nome, idade
}

// Para buildar o executável é só rodar go build *nome do arquivo.go*
// Para não precisar ficar dando build toda hora é só rodar "go run arquivo.go"
