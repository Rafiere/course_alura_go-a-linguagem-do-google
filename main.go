package main

import (
	"fmt"
	"reflect"
)

func main() {

	var nome string = "Rafa"
	var idade int = 21

	/* Podemos declarar e atribuir valor com o operador ":=". */
	versao := 1.0

	fmt.Println("Hello world,", nome, "- Versão: ", versao, "- Idade: ", idade)
	fmt.Println("O tipo da variável \"versao\" é: ", reflect.TypeOf(versao))

	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("3 - Sair do Programa")

	var comando int

	/* Se a "comando" é do tipo "int", o "Scan" já identifica que queremos receber um "int". */
	_, err := fmt.Scan(&comando)
	fmt.Println(err)

	fmt.Println("O comando escolhido foi:", comando)
	fmt.Println("O endereço de memória da minha variável é:", &comando)
}
