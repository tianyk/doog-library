package models

type Book struct {
	BookId string `form:"bookId" json:"book_id" xorm:"varchar(10) pk not null autoincr"`
	Title  string `form:"title" json:"title" xorm:"varchar(45) not null"`
	Isbn10 string `form:"isbn10" json:"isbn10" xorm:"varchar(10) not null"`
	Isbn13 string `form:"isbn13" json:"isbn13" xorm:"varchar(13) not null"`
}

func (b *Book) TableName() string {
	return "sys_book"
}

func (self *Book) GetById() error {
	_, err := orm.Id(self.BookId).Get(self)
	return err
}
