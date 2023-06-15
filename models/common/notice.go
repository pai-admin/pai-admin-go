package common

// Notice 公告
type Notice struct {
	NoticeId   uint   `gorm:"primarykey;comment:'主键'"`
	Title      string `gorm:"not null;default:'';comment:'公告名称'"`
	Content    string `gorm:"not null;default:'';comment:'公告内容'"`
	Rank       uint   `gorm:"not null;default:0;comment:'权重'"`
	DelFlag    uint8  `gorm:"not null;default:0;comment:'删除标识：0.否 1.是'"`
	Status     uint8  `gorm:"not null;default:1;comment:'是否显示：0.否 1.是'"`
	CreateTime string `gorm:"comment:'创建时间'"`
	UpdateTime string `gorm:"comment:'更新时间'"`
}
