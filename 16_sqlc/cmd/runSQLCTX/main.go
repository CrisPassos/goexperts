package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/cristianapassos/16sqlc/internal/db"
	_ "github.com/go-sql-driver/mysql"
)

type CourseDB struct {
	dbConn *sql.DB
	*db.Queries
}

func NewCourseDB(dbConn *sql.DB) *CourseDB {
	return &CourseDB{
		dbConn:  dbConn,
		Queries: db.New(dbConn),
	}
}

type CouseParams struct {
	ID          string
	Name        string
	Description sql.NullString
	Price       float64
}

type CategoryParams struct {
	ID          string
	Name        string
	Description sql.NullString
}

func (c *CourseDB) CreateCourseAndCategory(ctx context.Context, argsCategory CategoryParams, argsCourse CouseParams) error {
	err := c.callTX(ctx, func(q *db.Queries) error {
		var err error
		err = q.CreateCategory(ctx, db.CreateCategoryParams{
			ID:          argsCategory.ID,
			Name:        argsCategory.Name,
			Description: argsCategory.Description,
		})

		if err != nil {
			return err
		}

		err = q.CreateCourse(ctx, db.CreateCourseParams{
			ID:          argsCourse.ID,
			Name:        argsCourse.Name,
			CategoryID:  argsCategory.ID,
			Description: argsCourse.Description,
			Price:       argsCourse.Price,
		})

		if err != nil {
			return err
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (c *CourseDB) callTX(ctx context.Context, fn func(*db.Queries) error) error {
	tx, err := c.dbConn.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	q := db.New(tx)
	if err := fn(q); err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}
		return err
	}

	return tx.Commit()
}

func main() {
	ctx := context.Background()
	dbConn, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/courses")
	if err != nil {
		panic(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// courseArgs := CouseParams{
	// 	ID:          "course-1",
	// 	Name:        "Course 1",
	// 	Description: sql.NullString{String: "Description Course 1", Valid: true},
	// 	Price:       10.0,
	// }

	// categoryArgs := CategoryParams{
	// 	ID:          "category-1",
	// 	Name:        "Category 1",
	// 	Description: sql.NullString{String: "Description Category 1", Valid: true},
	// }

	// courseDB := NewCourseDB(dbConn)
	// err = courseDB.CreateCourseAndCategory(ctx, categoryArgs, courseArgs)
	// if err != nil {
	// 	panic(err)
	// }

	coursers, err := queries.ListCourses(ctx)
	if err != nil {
		panic(err)
	}
	for _, course := range coursers {
		fmt.Println(course.Name, course.Description.String, course.Price)
	}

}
