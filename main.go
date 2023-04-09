package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {

	testeTipagem()

	exibirMenu()

	var comando int

	lerComando(&comando)

	fmt.Println("O comando escolhido foi:", comando)
	fmt.Println("O endereço de memória da minha variável é:", &comando)

	fmt.Println("\n\n")

	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Exibindo logs...")
	} else if comando == 3 {
		fmt.Println("Saindo do programa...")
	} else {
		fmt.Println("Eu não conheço esse comando!")
		os.Exit(-1)
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
