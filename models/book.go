package models

type Book struct {
	BookId int
	Title  string
	Isbn10 string
	Isbn13 string
}

/**
 * 构造函数
 */
func NewBook() {
	return &Book{}
}

func (b *Book) FindById() (Book, error) {

}
