package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ProductGorm struct {
	ID    int `gorm:"primaryKey"`
	Name  string
	Price float64
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	//cria a base de dado de acorto com a struct
	db.AutoMigrate(&ProductGorm{})

	// create //

	// db.Create(&ProductGorm{
	// 	Name:  "Notebook",
	// 	Price: 1000.00,
	// })

	// create //

	// create many //

	// products := []ProductGorm{
	// 	{Name: "Notebook 2", Price: 1000.0},
	// 	{Name: "Mouse", Price: 50.0},
	// 	{Name: "Keyboard", Price: 100.0},
	// }
	// db.Create(&products)

	// create many //

	// select //
	// busca de um unico registro //

	// var product ProductGorm
	// db.First(&product, 4)
	// fmt.Println(product)
	// busca de um unico registro //

	// busca o primeiro registro com uma condicao //
	// db.First(&product, "name = ?", "Keyboard")
	// fmt.Println(product)

	// busca o primeiro registro com uma condicao //

	// Select ALL //

	// var allProducts []ProductGorm
	// db.Find(&allProducts)

	// for _, product := range allProducts {
	// 	fmt.Println(product)
	// }

	// Select ALL //

	// Select WITH Limits//

	// var allProducts []ProductGorm
	// db.Limit(5).Offset(2).Find(&allProducts)

	// for _, product := range allProducts {
	// 	fmt.Println(product)
	// }

	// Select WITH Limits //

	// Select with WHERE //

	// var products []ProductGorm
	// db.Where("price > ?", 100).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// Select with WHERE//

	// Select with LIKE //

	// var products []ProductGorm
	// db.Where("name LIKE ?", "%book%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	// Select with LIKE //

	// UPDATE E DELETE//

	// var p ProductGorm
	// db.First(&p, 1)
	// p.Name = "New Mouse"
	// db.Save(&p)

	// var p2 ProductGorm
	// db.First(&p2, 1)
	// fmt.Println(p2.Name)

	// db.Delete(&p2)

	// UPDATE E DELETE //

	// SOFT DELETE //

	var p2 ProductGorm
	db.First(&p2, 1)
	fmt.Println(p2.Name)

	db.Delete(&p2)

	// SOFT DELETE //

}
