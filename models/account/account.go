package account

import "gocli/core"

// Account 管理员表
type Account struct {
	AccountId  uint   `gorm:"primarykey;comment:'主键'"`
	DeptId     uint8  `gorm:"not null;default:0;comment:'所属部门'"`
	Username   string `gorm:"not null;default:'';comment:'用户名'"`
	Avatar     string `gorm:"default:'';comment:'头像''"`
	Password   string `gorm:"not null;default:'';comment:'密码'"`
	Salt       string `gorm:"not null;default:'';comment:'密码盐'"`
	Status     uint8  `gorm:"not null;default:1;comment:'状态：0.禁用 1.正常'"`
	DelFlag    uint8  `gorm:"not null;default:0;comment:'删除标识：0.否 1.是'"`
	CreateTime string `gorm:"comment:'创建时间'"`
	UpdateTime string `gorm:"comment:'更新时间'"`
}

func FindByUsername(username string) (account Account, err error) {
	err = core.DB.Where("username = ?", username).
		Where("del_flag = 0").
		Select("account_id, username, password, salt, status").
		Limit(1).
		First(&account).Error
	return
}
