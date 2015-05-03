package models

import (
	. "doog-library/utils"
)

type Book struct {
	BookId int
	Title  string
	Isbn10 string
	Isbn13 string
}

/**
 * 构造函数
 */
func NewBook() *Book {
	return new(Book)
}

func (b *Book) FindById(book_id string) *Book {
	db := GetConnection()
	// defer db.Close()

	stmtOut, err := db.Prepare("SELECT * FROM SYS_BOOK WHERE BOOK_ID = ?")
	CheckError(err)
	defer stmtOut.Close()

	var id int
	var title, isbn10, isbn13 string
	err = stmtOut.QueryRow(book_id).Scan(&id, &title, &isbn10, &isbn13)
	CheckError(err)

	return &Book{id, title, isbn10, isbn13}
}
