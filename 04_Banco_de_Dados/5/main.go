package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// has many
type Category struct {
	ID       int `gotm:primaryKey`
	Name     string
	Products []Product
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

	categoriesData := []Category{
		{Name: "Eletronicos"},
		{Name: "Cozinhas"},
	}

	db.Create(&categoriesData)

	productsData := []Product{
		{Name: "Teclado", Price: 45, CategoryID: 1},
		{Name: "Phones", Price: 89, CategoryID: 1},
		{Name: "Geladeira", Price: 78, CategoryID: 2},
	}

	db.Create(&productsData)

	var categories []Category
	err = db.Model(&Category{}).Preload("Products").Find(&categories).Error

	if err != nil {
		panic(err)
	}

	for _, cat := range categories {
		fmt.Println(cat.Name, ":")

		for _, prod := range cat.Products {
			fmt.Println("- ", prod.Name)
		}
	}

}
