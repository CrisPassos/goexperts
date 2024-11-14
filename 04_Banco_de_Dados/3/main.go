package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Category struct {
	ID   int `gotm:primaryKey`
	Name string
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float32
	CategoryID int
	Category   Category
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Product{}, &Category{})

	// category := Category{Name: "Eletronicos"}
	// db.Create(&category)

	// db.Create(&Product{
	// 	Name:       "Notebook",
	// 	Price:      100.00,
	// 	CategoryID: category.ID,
	// })

	//adicionando novo produto
	db.Create(&Product{
		Name:       "Mouse",
		Price:      1000,
		CategoryID: 1,
	})

	var products []Product
	db.Preload("Category").Find(&products)
	for _, product := range products {
		fmt.Println(product.Name, product.Category.Name)

	}

}
