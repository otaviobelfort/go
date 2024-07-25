package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Pessoa struct {
	ID   string `gorm:"primary_key"`
	Name string
	Cpf  string
}


func conectDatabaseGorm() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/test"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}

func main() {
	//db, err := conectDatabaseGorm()
	//if err != nil {
	//	panic(err)
	//}
	// cria a tabela automaticamente
	//db.AutoMigrate(&Pessoa{})

	//db.Create(&Pessoa{ID: "3", Name: "Pedro", Cpf: "02345678901"})
	//println("Pessoa criada com sucesso!")

	// fazendo um select
	//var pessoa Pessoa
	//db.First(&pessoa, "Name = ?", "Pedro")
	//fmt.Println("Pessoa: ", pessoa)

	// fazendo um select all
	//var pessoas []Pessoa
	//db.Find(&pessoas)
	//fmt.Println("Pessoas: ", pessoas)

	// fazendo um update
	//db.Updates(&Pessoa{ID: "1", Name: "Maria", Cpf: "12345678901"})

	buscaPessoaComBaseEmUmaListaPrevia()
}

func buscaPessoaComBaseEmUmaListaPrevia() {
	db, err := conectDatabaseGorm()
	if err != nil {
		panic(err)
	}
	//defer db.Close()
	var pessoas []string = []string{"Pedro", "Maria", "João"}
	var pessoasEncontradas []Pessoa
	db.Where("Name IN ?", pessoas).Find(&pessoasEncontradas)
	fmt.Println("Pessoas encontradas: ", pessoasEncontradas)

}
