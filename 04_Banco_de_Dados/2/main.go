package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Produto struct {
	ID    int `gorm:"primaryKey"`
	Nome  string
	Preco float32
	gorm.Model
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Produto{})

	// db.Create(&Produto{
	// 	Nome:  "Notebook",
	// 	Preco: 400.00,
	// })

	// products := []Produto{
	// 	{Nome: "Teclado", Preco: 45},
	// 	{Nome: "Phones", Preco: 89},
	// 	{Nome: "Mouse", Preco: 78},
	// }

	// db.Create(&products)

	//select one
	// var product Produto
	// db.First(&product, 2)
	// fmt.Println(product)

	//select where
	// db.First(&product, "nome = ?", "Teclado")
	// fmt.Println(product)

	//select all
	// var products []Produto
	// db.Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//operadores limit and offset (paginação)
	// var products []Produto
	// db.Limit(2).Offset(2).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//Usando where
	// var products []Produto
	// db.Where("preco > ?", 50).Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//LIKE
	// var products []Produto
	// db.Where("nome LIKE ?", "%a%").Find(&products)
	// for _, product := range products {
	// 	fmt.Println(product)
	// }

	//update
	// var p Produto
	// db.First(&p, 4)
	// p.Nome = "Rato"
	// db.Save(&p)

	// var p2 Produto
	// db.First(&p2, 4)
	// fmt.Println(p2.Nome)

	//deletar
	// db.Delete(p2)

	//DB create with , SOFT DELETE
	// db.Create(&Produto{
	// 	Nome:  "Notebook",
	// 	Preco: 400.00,
	// })

	// var p Produto
	// db.First(&p, 6)
	// p.Nome = "Rato"
	// db.Save(&p)

	// db.Delete(&p)

}
