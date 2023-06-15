package user

import "github.com/shopspring/decimal"

// User 用户
type User struct {
	UserId     uint   `gorm:"primarykey;comment:'主键'"`
	Username   string `gorm:"default:'';comment:'用户名'"`
	Avatar     string `gorm:"default:'';comment:'头像'"`
	Sex        uint8  `gorm:"default:0;comment:'性别：0.保密 1.男 2.女'"`
	Birthday   string `gorm:"comment:'生日'"`
	Openid     string `gorm:"not null;comment:'openid'"`
	DelFlag    uint8  `gorm:"not null;default:0;comment:'删除标识：0.否 1.是'"`
	CreateTime string `gorm:"default:null;comment:'创建时间'"`
	UpdateTime string `gorm:"default:null;comment:'更新时间'"`
}

// UserFeedback 用户反馈
type UserFeedback struct {
	FeedbackId uint   `gorm:"primarykey;comment:'主键'"`
	UserId     uint   `gorm:"not null;default:0;comment:'用户ID'"`
	Title      string `gorm:"comment:'标题'"`
	Email      string `gorm:"comment:'邮箱'"`
	Wxcode     string `gorm:"comment:'微信号'"`
	Mobile     string `gorm:"comment:'手机号'"`
	Content    string `gorm:"not null;comment:'内容'"`
	CreateTime string `gorm:"default:null;comment:'创建时间'"`
}

// UserDonate 用户捐赠
type UserDonate struct {
	DonateId   uint            `gorm:"primarykey;comment:'主键'"`
	UserId     uint            `gorm:"not null;default:0;comment:'用户ID'"`
	Money      decimal.Decimal `gorm:"not null;comment:'支付金额'"`
	OrderNo    string          `gorm:"not null;comment:'订单编号'"`
	PayStatus  uint8           `gorm:"default:0;comment:'支付状态'"`
	PayTime    string          `gorm:"default:null;comment:'支付时间'"`
	CreateTime string          `gorm:"default:null;comment:'创建时间'"`
}

// UserFavorite 用户收藏
type UserFavorite struct {
	FavoriteId uint   `gorm:"primarykey;comment:'主键'"`
	UserId     uint   `gorm:"not null;default:0;comment:'用户ID'"`
	ArticleId  uint   `gorm:"not null;comment:'文章ID'"`
	DelFlag    uint8  `gorm:"not null;default:0;comment:'是否删除：0.否 1.是'"`
	CreateTime string `gorm:"default:null;comment:'创建时间'"`
	UpdateTime string `gorm:"default:null;comment:'更新时间'"`
}
