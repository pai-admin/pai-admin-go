package common

// Config 文章分类
type Config struct {
	ConfigId   uint   `gorm:"primarykey;comment:'主键'"`
	Flag       string `gorm:"not null;default:'';comment:'配置项标识'"`
	Name       string `gorm:"not null;comment:'配置项名称'"`
	Content    string `gorm:"comment:'配置项内容'"`
	DelFlag    uint8  `gorm:"not null;default:1;comment:'是否删除：0.否 1.是'"`
	CreateTime string `gorm:"comment:'创建时间'"`
	UpdateTime string `gorm:"comment:'更新时间'"`
}
