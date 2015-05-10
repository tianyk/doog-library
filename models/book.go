package models

type Book struct {
	BookId int64  `form:"bookId" json:"book_id" xorm:"BigInt pk not null autoincr"`
	Title  string `valid:"Required; MaxSize(100)" form:"title" json:"title" xorm:"varchar(45) not null"`
	// TODO 添加校验类型Isbn10、Isbn13
	Isbn10 string `valid:"Required; Length(10)" form:"isbn10" json:"isbn10" xorm:"varchar(10) not null"`
	Isbn13 string `valid:"Required; Length(13)" form:"isbn13" json:"isbn13" xorm:"varchar(13) not null"`
}

// 如果你的 struct 实现了接口 validation.ValidFormer
// 当 StructTag 中的测试都成功时，将会执行 Valid 函数进行自定义验证
// func (u *user) Valid(v *validation.Validation) {
// 	if strings.Index(u.Name, "admin") != -1 {
// 		// 通过 SetError 设置 Name 的错误信息，HasErrors 将会返回 true
// 		v.SetError("Name", "名称里不能含有 admin")
// 	}
// }

func (self *Book) TableName() string {
	return "sys_book"
}

func (self *Book) GetById() error {
	_, err := orm.Id(self.BookId).Get(self)
	return err
}

func (self *Book) Save() error {
	_, err := orm.InsertOne(self)
	return err
}

func (self *Book) Page( /*q, tag string, */ start, count int) (total int64, bookPage []*Book, err error) {
	total, err = orm. /*Where("title like ?", q)*/ Count(self)
	err = orm. /*Where("title like ?", q)*/ Limit(count, start).Find(&bookPage)
	return
}
