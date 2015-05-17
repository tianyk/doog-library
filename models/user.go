package models

type User struct {
	UserId       int64  `form:"userId" json:"user_id" xorm:"BigInt pk not null autoincr"`
	NickName     string `form:"nickName" json:"nick_name" xorm:"string not null"`
	Email        string `form:"email" json:"email" xorm:"string not null"`
	WeiboOpenId  string `form:"weiboOpenId" json:"weibo_openid" xorm:"string"`
	WeixinOpenId string `form:"weixinOpenId" json:"weixin_openid" xorm:"string"`
	QOOpenId     string `form:"qqOpenId" json:"qq_openid" xorm:"string"`
	DoubanOpenId string `form:"doubanOpenId" json:"douban_openid" xorm:"string"`
}
