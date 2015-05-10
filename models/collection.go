package models

type Collection struct {
	BookId int64  `valid:"Required;" form:"bookId" json:"book_id" xorm:"BigInt pk not null autoincr"`
	UserId int64  `valid:"Required;" form:"userId" json:"user_id" xorm:"BigInt pk not null autoincr"`
	State  string `form:"state" json:"state" xorm:"varchar(10) not null default 'await'"`
}

func (c *Collection) TableName() string {
	return "sys_collection"
}

// 查询所有藏书
func (self *Collection) FindCollectionsByBookId(bookId int64) (collections []*Collection, err error) {
	err = orm.Where("book_id = ?", bookId).Find(&collections)
	return
}

func (self *Collection) Save() error {
	_, err := orm.InsertOne(self)
	return err
}

// 查询用户的所有藏书
func (self *Collection) FindUserBooks(userId int64) (collections []*Collection, err error) {
	err = orm.Where("user_id = ?", userId).Find(&collections)
	return
}
