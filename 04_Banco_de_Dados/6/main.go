package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// many to many, precisamos usar anotacões
type Category struct {
	ID       int `gorm:"primaryKey"`
	Name     string
	Products []Product `gorm:"many2many:products_categories;"`
}

type Product struct {
	ID         int `gorm:"primaryKey"`
	Name       string
	Price      float32
	Categories []Category `gorm:"many2many:products_categories;"`
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
		{Name: "Cozinha"},
	}

	db.Create(&categoriesData)

	productsData := []Product{
		{Name: "Teclado", Price: 45, Categories: []Category{categoriesData[0], categoriesData[1]}},
		{Name: "Phones", Price: 89, Categories: []Category{categoriesData[0]}},
		{Name: "Geladeira", Price: 78, Categories: []Category{categoriesData[1]}},
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
