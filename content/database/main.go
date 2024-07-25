package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
	"log"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func conectDatabase() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
	if err != nil {
		panic(err)
	}
	return db, err
}
func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

// crie uma função para alterar um produto no banco de dados
func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

// crie uma função para buscar um produto no banco de dados
func getProduct(db *sql.DB, product *Product) (Product, error) {
	stmt, err := db.Prepare("select name, price from products where id = ?")
	if err != nil {
		return Product{}, err
	}
	row := stmt.QueryRow(product.ID)
	err = row.Scan(&product.Name, &product.Price)
	if err != nil {
		return Product{}, err
	}
	return *product, nil

}

func processandoAlteracaoProduto() {
	db, err := conectDatabase()
	if err != nil {
		panic(err)
	}
	product, err := getProduct(db, &Product{ID: "1fefea91-126b-4bb3-bb8e-05c935ef550a"})
	if err != nil {
		log.Println("Erro ao buscar produto")
		panic(err)
	}
	product.Name = "NOTEBOOK DELL I7 2024"
	product.Price = 3200.99
	err = updateProduct(db, &product)
	if err != nil {
		log.Println("Erro ao atualizar o produto", product.Name)
		panic(err)
	}
	defer db.Close()

}

// crie uma função para listar todos os produtos no banco de dados
func getProducts() ([]Product, error) {
	db, err := conectDatabase()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	stmt, err := db.Prepare("select id, name, price from products")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []Product
	for rows.Next() {
		var product Product
		rows.Scan(&product.ID, &product.Name, &product.Price)
		products = append(products, product)
	}
	return products, nil

}

// crie uma função para deletar um produto no banco de dados
func deleteProduct(product *Product) error {
	db, err := conectDatabase()
	if err != nil {
		return err
	}
	defer db.Close()

	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID)
	if err != nil {
		return err
	}
	return nil
}

func setProduct() {
	db, err := conectDatabase()
	product := NewProduct("Product 3", 23.99)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	defer db.Close()

}

func main() {
	processandoAlteracaoProduto()
	products, _ := getProducts()
	fmt.Println(products)
	//fmt.Println("Deletando produto...")
	//deleteProduct(&Product{ID: "1fefea91-126b-4bb3-bb8e-05c935ef550a"})
	//products, _ = getProducts()

}
