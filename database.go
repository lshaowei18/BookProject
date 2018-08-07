package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

func (app App) InsertSingleBook(name, author string) {
	//Prepare statement
	statement := `
  INSERT INTO books(name, author)
  VALUES ($1, $2)`

	//Execute statement
	_, err := app.DB.Exec(statement, name, author)
	CheckErr(err)
}

func (app App) DeleteBookByID(id int) {
	statement := `DELETE FROM books WHERE id = $1`
	result, err := app.DB.Exec(statement, id)
	CheckErr(err)
	fmt.Println(result.RowsAffected())
}

func (app App) FindSingleBookByName(name string) Book {
	statement := `SELECT * FROM books WHERE name = $1`
	row := app.DB.QueryRow(statement, name)
	book := Book{}
	err := row.Scan(&book.ID, &book.Name, &book.Author)
	CheckErr(err)
	return book
}
