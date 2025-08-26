package database

import (
	"database/sql"

	"github.com/google/uuid"
)

type Category struct {
	db          *sql.DB
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func NewCategory(db *sql.DB) *Category {
	return &Category{
		db: db,
	}
}

func (c *Category) Create(name string, description string) (Category, error) {
	id := uuid.New().String()
	_, err := c.db.Exec("INSERT INTO categories (id, name, description) VALUES ($1, $2, $3)",
		id, name, description)
	if err != nil {
		return Category{}, err
	}
	return Category{ID: id, Name: name, Description: description}, nil
}

func (c *Category) GetAll() ([]Category, error) {
	rows, err := c.db.Query("SELECT id, name, description FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		if err := rows.Scan(&category.ID, &category.Name, &category.Description); err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return categories, nil
}
