package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"reflect"
	"strings"
)

func main() {

	testeTipagem()

	/* Estamos criando um "while". */
	for {
		exibirMenu()

		var comando int

		lerComando(&comando)

		fmt.Println("O comando escolhido foi:", comando)
		fmt.Println("O endereço de memória da minha variável é:", &comando)

		fmt.Println("\n\n")

		if comando == 1 {
			iniciarMonitoramento()
		} else if comando == 2 {
			fmt.Println("Exibindo logs...")
		} else if comando == 3 {
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		} else {
			fmt.Println("Eu não conheço esse comando!")
			os.Exit(-1)
		}
	}

}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	sites := lerSitesDoArquivo()

	for _, site := range sites {
		testSite(site)
	}

}

func lerSitesDoArquivo() []string {

	var sites []string

	file, err := os.Open("sites.txt")
	if err != nil {
		fmt.Println("Ocorreu um erro: ", err)
	}

	leitor := bufio.NewReader(file)

	for {
		/* Retirando os espaços em branco do final do arquivo. */
		linha, err := leitor.ReadString('\n')
		linha = strings.TrimSpace(linha)
		fmt.Println(linha)
		if err == io.EOF {
			break
		}
	}

	file.Close()
	return sites
}

func testSite(site string) {
	resp, _ := http.Get(site)

	if resp.StatusCode == 200 {
		fmt.Println("O site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("O site", site, "está com problemas. Status code:", resp.StatusCode)
	}
}

func lerComando(comando *int) {
	/* Se a "comando" é do tipo "int", o "Scan" já identifica que queremos receber um "int". */
	_, err := fmt.Scan(comando)
	fmt.Println(err)
}

func testeTipagem() {
	var nome string = "Rafa"
	var idade int = 21

	/* Podemos declarar e atribuir valor com o operador ":=". */
	versao := 1.0

	fmt.Println("Hello world,", nome, "- Versão: ", versao, "- Idade: ", idade)
	fmt.Println("O tipo da variável \"versao\" é: ", reflect.TypeOf(versao))
}

func exibirMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")
}
