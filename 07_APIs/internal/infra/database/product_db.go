package database

import (
	"github.com/CrisPassos/goexpert/7_APIs/internal/entity"
	"gorm.io/gorm"
)

type Product struct {
	DB *gorm.DB
}

func NewProduct(db *gorm.DB) *Product {
	return &Product{DB: db}
}

func (p *Product) Create(product *entity.Product) error {
	return p.DB.Create(product).Error
}

func (p *Product) FindAll(page, limit int, sort string) ([]entity.Product, error) {
	var product []entity.Product
	var err error

	//se não houver sort, vamos obrigar a ordernar por asc
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	// caso haja valor em page e limit, paginamos usando o offset e ordernamos usando o Order
	// offset sempre começa do 0
	// usamos o find para armazenar o resultado na varíavel(ponteiro) &product
	if page != 0 && limit != 0 {
		err = p.DB.Limit(limit).Offset((page - 1) * limit).Order("created_at " + sort).Find(&product).Error
	} else {
		err = p.DB.Order("created_at " + sort).Find(&product).Error
	}

	return product, err

}

func (p *Product) FindById(id string) (*entity.Product, error) {
	var product entity.Product
	err := p.DB.First(&product, "id = ?", id).Error
	return &product, err
}

func (p *Product) Update(product *entity.Product) error {
	_, err := p.FindById(product.ID.String())

	if err != nil {
		return err
	}

	return p.DB.Save(product).Error
}

func (p *Product) Delete(id string) error {
	product, err := p.FindById(id)

	if err != nil {
		return err
	}

	return p.DB.Delete(product).Error
}
