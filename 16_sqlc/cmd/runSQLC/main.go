package main

import (
	"context"
	"database/sql"

	"github.com/cristianapassos/16sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// err = queries.CreateCategory(ctx, db.CreateCategoryParams{
	// 	ID:          uuid.New().String(),
	// 	Name:        "Category 1",
	// 	Description: sql.NullString{String: "Description 1", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// categories, err := queries.ListCategories(ctx)

	// if err != nil {
	// 	panic(err)
	// }

	// for _, category := range categories {
	// 	println(category.Name, category.Description.String, category.ID)
	// }

	// err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
	// 	ID:          "2e8a421d-d750-488a-9427-6abe6a9eba56",
	// 	Name:        "Category 1 Updated",
	// 	Description: sql.NullString{String: "Description 1 Updated", Valid: true},
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// err = queries.DeleteCategory(ctx, "2e8a421d-d750-488a-9427-6abe6a9eba56")

	// if err != nil {
	// 	panic(err)
	// }

	categories, err := queries.ListCategories(ctx)

	if err != nil {
		panic(err)
	}
	for _, category := range categories {
		println(category.Name, category.Description.String, category.ID)
	}

}
