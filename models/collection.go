package models

type Collection struct {
	BookId string `form:"bookId" json:"book_id" xorm:"string(10) pk not null autoincr"`
	UserId string `form:"userId" json:"user_id" xorm:"string(10) pk not null autoincr"`
	State  string `form:"state" json:"state" xorm:"varchar(10) not null"`
}

func (c *Collection) TableName() string {
	return "sys_collection"
}

func (self *Collection) FindCollectionsByBookId(bookId string) (collections []*Collection, err error) {
	err = orm.Where("book_id = ?", bookId).Find(&collections)
	return
}
