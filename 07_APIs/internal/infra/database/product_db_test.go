package database

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/CrisPassos/goexpert/7_APIs/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func createDatabaseConnection() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	return db, err
}

func TestCreateNewProduct(t *testing.T) {
	db, err := createDatabaseConnection()

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})
	product, err := entity.NewProduct("Product 1", 10.5)
	assert.NoError(t, err)
	productDB := NewProduct(db)
	err = productDB.Create(product)
	assert.NoError(t, err)
	assert.NotEmpty(t, product.ID)
}

func TestFindAllProducts(t *testing.T) {
	db, err := createDatabaseConnection()

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	for i := 1; i < 24; i++ {
		product, err := entity.NewProduct(fmt.Sprintf("Product %d", i), rand.Float64()*100)
		assert.NoError(t, err)
		db.Create(product)
	}

	productDB := NewProduct(db)
	products, err := productDB.FindAll(1, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 1", products[0].Name)
	assert.Equal(t, "Product 10", products[9].Name)

	products, err = productDB.FindAll(2, 10, "asc")
	assert.NoError(t, err)
	assert.Len(t, products, 10)
	assert.Equal(t, "Product 11", products[0].Name)
	assert.Equal(t, "Product 20", products[9].Name)
}

func TestFindbyI(t *testing.T) {
	db, err := createDatabaseConnection()

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.5)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	productFind, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFind.Name)
	assert.Equal(t, product.Price, productFind.Price)
}

func TestUpdate(t *testing.T) {
	db, err := createDatabaseConnection()

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.5)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	product.Name = "Product 2"
	product.Price = 20.5
	err = productDB.Update(product)
	assert.NoError(t, err)

	productFind, err := productDB.FindById(product.ID.String())
	assert.NoError(t, err)
	assert.Equal(t, product.Name, productFind.Name)
	assert.Equal(t, product.Price, productFind.Price)
}

func TestDelete(t *testing.T) {
	db, err := createDatabaseConnection()

	if err != nil {
		t.Error(err)
	}

	db.AutoMigrate(&entity.Product{})

	product, err := entity.NewProduct("Product 1", 10.5)
	assert.NoError(t, err)
	db.Create(product)

	productDB := NewProduct(db)
	err = productDB.Delete(product.ID.String())
	assert.NoError(t, err)

	_, err = productDB.FindById(product.ID.String())
	assert.Error(t, err)
}
