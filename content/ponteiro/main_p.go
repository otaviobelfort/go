package ponteiro

import (
	"fmt"
)

func main_2() {
	// a variavel -> aponta para um ponteiro que tem um endereço na memoria -> valor
	nome := "Otávio"
	var ponteiro_nome *string = &nome

	// alterar o valor de um endereço
	*ponteiro_nome = "José"
	println(*ponteiro_nome)
	println(nome)

}

type Cliente struct {
	nome string
}

func (c *Cliente) dadosCliente() {
	c.nome = "José"
	fmt.Print("nome: ", c.nome)
}

func main() {
	cliente := Cliente{nome: "Otávio"}

	cliente.dadosCliente()
	fmt.Print("nome: ", cliente.nome)

}
