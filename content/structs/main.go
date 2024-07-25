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

type PessoaComposicao struct {
	nome  string
	cpf   string
	idade int
	Endereco
}

func main() {

	struct_basic()
}

//composição

func struct_basic() {

	pessoa := Pessoa{
		nome:  "Otávio",
		cpf:   "123.456.789-10",
		idade: 30,
	}
	pessoa.setIdade(50)
	fmt.Printf("Nome, %s\n CPF, %d\n IDADE, %d\n", pessoa.nome, pessoa.idade, pessoa.idade)
}

func struct_composta() {

	endereco := Endereco{
		logradouro: "Rua 1",
		cep:        "12345678",
		numero:     123,
	}

	pessoa := PessoaCompleta{
		nome:     "Otávio",
		cpf:      "123.456.789-10",
		idade:    30,
		endereco: endereco,
	}
	fmt.Printf("Nome: %s\n CPF: %d\n IDADE: %d\n", pessoa.nome, pessoa.cpf, pessoa.idade)
	fmt.Printf("Logradouro: %s\n cep: %s\n ", pessoa.endereco.logradouro, pessoa.endereco.cep)
}

func (p Pessoa) setIdade(idade int) {
	p.idade = idade
}

func struct_composicao() {

	endereco := Endereco{
		logradouro: "Rua 1",
		cep:        "12345678",
		numero:     123,
	}

	pessoa := PessoaComposicao{
		nome:     "Otávio",
		cpf:      "123.456.789-10",
		idade:    30,
		Endereco: endereco,
	}
	pessoa.cep = "123.456.789-10"
	fmt.Printf("Nome: %s\n CPF: %d\n IDADE: %d\n", pessoa.nome, pessoa.idade, pessoa.idade)
	fmt.Printf("Logradouro: %s\n cep: %s\n ", pessoa.logradouro, pessoa.cep)
}
