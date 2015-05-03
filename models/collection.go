package models

import (
	. "doog-library/utils"
)

type Collection struct {
	BookId string
	UserId string
	State  string
}

func NewCollection() *Collection {
	return new(Collection)
}

func (c *Collection) FindBookCollectionsByBookId(book_id string) []Collection {
	db := GetConnection()
	rows, err := db.Query("SELECT * FROM SYS_COLLECTION WHERE BOOK_ID = ?", book_id)
	CheckError(err)
	defer rows.Close()

	collections := make([]Collection, 0, 10)
	for rows.Next() {
		var bookId, userId, state string
		rows.Scan(&bookId, &userId, &state)
		collections = append(collections, Collection{bookId, userId, state})
	}

	return collections
}
