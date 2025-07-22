package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const MONITORAMENTOS = 3
const INTERVALO_MONITORAMENTO = 5 * time.Second

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
			imprimeLogs()
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
	// site := "https://httpbin.org/status/404" // Para requisição sem sucesso
	// sites := []string{"https://httpbin.org/status/200", "https://httpbin.org/status/404", "https://www.alura.com.br", "https://www.caelum.com.br"}
	sites := leSisteDoArquivo()

	for i := 0; i < MONITORAMENTOS; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("")
		time.Sleep(INTERVALO_MONITORAMENTO)
	}
}

func leSisteDoArquivo() []string {
	var sites []string
	arquivo, err := os.Open("sites.txt")
	// arquivo, err := ioutil.ReadFile("sites.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	leitor := bufio.NewReader(arquivo)

	for {
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		if err == io.EOF {
			break
		}
		sites = append(sites, linha)
	}

	arquivo.Close()

	return sites
}

func testaSite(site string) {
	fmt.Println("Testando site: '", site, "'")
	response, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro na requisição:", err)
	}

	switch response.StatusCode {
	case 200:
		fmt.Println("Site: ", site, " foi carregado com sucesso")
		registraLog(site, true)
	default:
		fmt.Println("Site: ", site, " está com problema, status code: ", response.StatusCode)
		registraLog(site, false)
	}
	// fmt.Println("")
}

func registraLog(site string, status bool) {
	const permissao = 0666
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, permissao)

	if err != nil {
		fmt.Println(err)
	}

	arquivo.WriteString(time.Now().Format("02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	arquivo.Close()
}

func imprimeLogs() {
	// Maneira com openfile
	// const permissao = 0666
	// arquivo, err := os.OpenFile("log.txt", os.O_RDONLY, permissao)

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// reader := bufio.NewReader(arquivo)

	// for {
	// 	linha, err := reader.ReadString('\n')
	// 	linha = strings.TrimSpace(linha)
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	fmt.Println(linha)
	// }

	// Maneira com ioutil
	arquivo, err := ioutil.ReadFile("log.txt")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(arquivo))

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
	nomes = append(nomes, "Nome extra")
	fmt.Println("O meu slice agora tem", len(nomes), "nomes e capacidade para ", cap(nomes), "nomes")
}

// Para buildar o executável é só rodar go build *nome do arquivo.go*
// Para não precisar ficar dando build toda hora é só rodar "go run arquivo.go"
