package databaseintegrationtest

import (
	"context"
	"database/sql"
)

type Book struct {
	Title  string
	Author string
}

var getBooksQuery = `SELECT * FROM book`
var insertBookQuery = `INSERT INTO book VALUES ($1, $2)`

func GetBooks(ctx context.Context, db *sql.DB) ([]Book, error) {
	rows, err := db.QueryContext(ctx, getBooksQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	books := []Book{}
	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.Title, &book.Author)
		if err != nil {
			return books, err
		}
		books = append(books, book)
	}

	return books, nil
}

func InsertBook(ctx context.Context, db *sql.DB, title, author string) error {
	_, err := db.ExecContext(ctx, insertBookQuery, title, author)

	return err
}
