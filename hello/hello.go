package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// _, idade := devolveNomeEIdade() Esse underline faz o go ignorar o primeiro retorno
	// fmt.Println("tenho ", idade, " anos")
	// exibeNomes()

	exibeIntroducao()
	for {
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
		}
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
	// site := "https://httpbin.org/status/200" // Para requisição com sucesso
	site := "https://httpbin.org/status/404" // Para requisição sem sucesso
	var sites [4]string
	sites[0] = "https://httpbin.org/status/404"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"

	fmt.Println(sites)

	response, _ := http.Get(site)

	switch response.StatusCode {
	case 200:
		fmt.Println("Site: ", site, " foi carregado com sucesso")
	default:
		fmt.Println("Site: ", site, " está com problema, status code: ", response.StatusCode)
	}
}

func devolveNomeEIdade() (string, int) {
	nome := "Robert"
	idade := 24
	return nome, idade
}

func exibeNomes() {
	nomes := []string{"Robert", "Alexandre", "Almeida"}
	fmt.Println(nomes)
	fmt.Println("O meu slice tem", len(nomes))
}

// Para buildar o executável é só rodar go build *nome do arquivo.go*
// Para não precisar ficar dando build toda hora é só rodar "go run arquivo.go"
