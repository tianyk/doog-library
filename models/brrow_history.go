package models

type BrrowHistory struct {
	BrrowHistoryId int64  `form:"BrrowHistoryId" json:"brrow_history_id" xorm:"BigInt pk not null autoincr"`
	BookId         int64  `valid:"Required;" form:"bookId" json:"book_id" xorm:"BigInt not null"`
	UserId         int64  `valid:"Required;" form:"userId" json:"user_id" xorm:"BigInt not null"`
	OwnerUserId    int64  `valid:"Required;" form:"ownerUserId" json:"owner_user_id" xorm:"BigInt not null"`
	LoanAt         int64  `valid:"Required;" form:"loanAt" json:"loan_at" xorm:"BigInt not null"`
	ReturnAt       int64  `valid:"Required;" form:"returnAt" json:"return_at" xorm:"BigInt not null"`
	State          string `form:"state" json:"state" xorm:"varchar(10) not null default 'reading'"`
}

func (self *BrrowHistory) TableName() string {
	return "sys_brrow_history"
}

func (self *BrrowHistory) Save() error {
	_, err := orm.InsertOne(self)
	return err
}
