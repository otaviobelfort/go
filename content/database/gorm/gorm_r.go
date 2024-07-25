package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// mysql -uroot -p test
// casos diferentes de many to many
//type Product struct {
//	ID           int `gorm:"primary_key"`
//	Name         string
//	Price        string
//	CategoryID   int
//	Category     Category
//	SerialNumber SerialNumber
//}

type Product struct {
	ID         int `gorm:"primary_key"`
	Name       string
	Price      string
	Categories []Category `gorm:"many2many:product_categories"`
}

type Category struct {
	ID       int `gorm:"primary_key"`
	Name     string
	Products []Product `gorm:"many2many:product_categories"`
}

//type SerialNumber struct {
//	ID        int `gorm:"primary_key"`
//	Number    string
//	ProductID int
//}

func conectDatabase() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:root@tcp(localhost:3306)/test"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}

func main() {

	db, err := conectDatabase()
	if err != nil {
		panic(err)
	}

	// cria a tabela automaticamente
	db.AutoMigrate(&Product{}, &Category{})

	//category := Category{Name: "Eletrônicos"}
	//db.Create(&category)
	//category2 := Category{Name: "Cozinha"}
	//db.Create(&category2)
	//
	//product := Product{Name: "Notebook", Price: "2000", Categories: []Category{category, category2}}
	//db.Create(&product)

	//serialNumber := SerialNumber{Number: "123456", ProductID: 1}
	//db.Create(&serialNumber)
	//buscaProductsFromCategory(db)
}

//func buscaProductComCategory(db *gorm.DB) {
//	// db, err := conectDatabase()
//	// if err != nil {
//	// 	panic(err)
//	//}
//	var products []Product
//	db.Preload("Category").Preload("SerialNumber").Find(&products)
//	for _, product := range products {
//		fmt.Printf("Product: %s, Category: %s, Number: %s ", product.Name, product.Category.Name, product.SerialNumber.Number)
//	}
//
//}

//func buscaProductsFromCategory(db *gorm.DB) {
//	var categories []Category
//	db.Model(&Category{}).Preload("Products.SerialNumber").Find(&categories)
//
//	for _, category := range categories {
//		for _, product := range category.Products {
//			fmt.Printf("Product: %s, Number: %s, Category: %s", product.Name, product.SerialNumber.Number, category.Name)
//		}
//		//fmt.Printf("Product: %s, Category: %s, Number: %s ", product.Name, product.Category.Name, product.SerialNumber.Number)
//	}
//
//}
