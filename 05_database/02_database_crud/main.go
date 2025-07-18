package main

// _blank identifier permite usar uma extensão sem que necessáriamente eu esteja usando
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
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

func main() {
	//Criando a conexão com o bando de dados
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// Inserindo um novo produto //
	product := NewProduct("Notebook", 2000.25)

	err = insertProduct(db, *product)

	if err != nil {
		panic(err)
	}
	// Inserindo um novo produto //

	// Atualizando o produto //
	product.Name = "Notebook Apple"

	err = updateProduct(db, product)

	if err != nil {
		panic(err)
	}

	// Atualizando o Produto //

	// Selecionando apenas 1 produto
	p, err := selectProduct(db, product.ID)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %v, possui o preço de %.2f", p.Name, p.Price)
	// Selecionando apenas 1 Produto

	// Selecionando todos os valores //
	products, err := selectAllProducts(db)

	if err != nil {
		panic(err)
	}

	for _, p := range products {
		fmt.Printf("Product: %v, possui o preço de %.2f\n", p.Name, p.Price)
	}
	// Selecionando todos os valores //

	// delete item //
	err = deleteProduct(db, product.ID)

	if err != nil {
		panic(err)
	}
	// delete item //

}

func insertProduct(db *sql.DB, product Product) error {
	//protegendo contra sql inject através do prepared statment, sanitização
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

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price =? where id = ?")

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

func selectProduct(db *sql.DB, id string) (*Product, error) {
	stmt, err := db.Prepare("select id, name, price from products where id = ?")

	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	// eu posso trabalhar com contextos tbm
	// stmt.QueryContext(ctx, id).Scan(&p.ID, &p.Name, &p.Price)

	// faz um de/para de database para memoria
	var p Product
	err = stmt.QueryRow(id).Scan(&p.ID, &p.Name, &p.Price)

	if err != nil {
		return nil, err
	}

	// eu vou retornar a referência do objeto em memória e não uma cópia
	return &p, nil
}

func selectAllProducts(db *sql.DB) ([]Product, error) {
	//chamada direta sem prepare, pq não vamos sofre ataque de inject
	rows, err := db.Query("select id, name, price from products")

	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var products []Product

	for rows.Next() {
		// vou adicionar o resultado na minha variavel P
		var p Product
		err = rows.Scan(&p.ID, &p.Name, &p.Price)

		if err != nil {
			return nil, err
		}

		products = append(products, p)
	}

	return products, nil
}

func deleteProduct(db *sql.DB, id string) error {
	stmt, err := db.Prepare("delete from products where id = ?")
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)

	if err != nil {
		return err
	}

	return nil
}
