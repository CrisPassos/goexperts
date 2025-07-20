package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gorm:"primaryKey"`
	Name string
	//Has Many
	Products []Product
}

type Product struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	// um produto tem uma Categoria
	CategoryID int
	Category   Category
	//has one
	SerialNumber SerialNumber
	gorm.Model
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	//cria a base de dado de acorto com a struct
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

	// BELONGS TO //
	//create category

	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// create product
	// db.Create((&Product{
	// 	Name:       "Notebook",
	// 	Price:      1000.00,
	// 	CategoryID: category.ID,
	// }))
	// BELONGS TO //

	// HAS ONE //

	// create serial number
	// db.Create(&SerialNumber{
	// 	Number:    "1233453",
	// 	ProductID: 1,
	// })

	// //Buscar todos os produtos
	// var products []Product
	// db.Preload("Category").Preload("SerialNumber").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product.Name, product.Category.Name, product.SerialNumber.Number)
	// }

	// HAS ONE //

	// HAS MANY //

	// category2 := Category{Name: "Eletronicos Has Many"}
	// db.Create(&category2)

	// db.Create(&Product{
	// 	Name:       "Margarida",
	// 	Price:      2000.00,
	// 	CategoryID: 3,
	// })

	// category2 := Category{Name: "Cozinha"}
	// db.Create(&category2)

	// db.Create(&Product{
	// 	Name:       "Margarida",
	// 	Price:      2000.00,
	// 	CategoryID: 5,
	// })

	// var categories2 []Category
	// err = db.Model(&Category{}).Preload("Products").Find(&categories2).Error

	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories2 {
	// 	fmt.Print(category.Name, ":")
	// 	for _, product := range category.Products {
	// 		println("-", product.Name, category.Name)

	// 	}
	// }

	// HAS MANY //

	// PEGADINHA //

	// category := Category{Name: "Cozinha"}
	// db.Create(&category)

	// //create product
	// db.Create(&Product{
	// 	Name:       "Panela",
	// 	Price:      99.0,
	// 	CategoryID: 1,
	// })

	// //create serial number
	// db.Create(&SerialNumber{
	// 	Number:    "12312414",
	// 	ProductID: 1,
	// })

	// var categories []Category
	// //o Preload não é na tabela e sim no relacionamento
	// err = db.Model(&Category{}).Preload("Products").Preload("Products.SerialNumber").Find(&categories).Error

	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	fmt.Println(category.Name, ":")

	// 	for _, product := range category.Products {
	// 		println("- ", product.Name, "Serial Number: ", product.SerialNumber.Number)
	// 	}
	// }

	// PEGADINHA

}
