package main

import "fmt"

const a = "Otávio!"

var b bool

// criando um type
type ID int64

// pode ter declaração e atribuição
type Pessoa struct {
	nome  string
	cpf   string
	idade int
}
type Endereco struct {
	logradouro string
	cep        string
	numero     int
}
type PessoaCompleta struct {
	nome     string
	cpf      string
	idade    int
	endereco Endereco
}
type PessoaFisica interface {
	Desativar()
}

func (pessoa Pessoa) Desativar() {
	pessoa.cpf = "100.100"
	fmt.Println("Desativado")
}
func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
	fmt.Println("Desativado")
}

func main() {
	pessoa := Pessoa{
		nome:  "Otávio",
		cpf:   "123.456.789-10",
		idade: 30,
	}

	struct_basic()
	Desativacao(pessoa)
}

//composição

func struct_basic() {

	pessoa := Pessoa{
		nome:  "Otávio",
		cpf:   "123.456.789-10",
		idade: 30,
	}
	fmt.Printf("Nome, %s\n CPF, %d\n IDADE, %d\n", pessoa.nome, pessoa.idade, pessoa.idade)
}
