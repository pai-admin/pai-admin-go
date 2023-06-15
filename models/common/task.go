package common

// Task 任务
type Task struct {
	TaskId     uint   `gorm:"primarykey;comment:'主键'"`
	UserId     string `gorm:"not null;default:'';comment:'发送对象'"`
	Title      string `gorm:"not null;default:'';comment:'任务名称'"`
	Content    string `gorm:"not null;default:'';comment:'发送内容'"`
	Template   string `gorm:"not null;default:'';comment:'发送模板'"`
	DelFlag    uint8  `gorm:"not null;default:0;comment:'删除标识：0.否 1.是'"`
	Status     uint8  `gorm:"not null;default:0;comment:'0未发送 1已发送 2发送失败'"`
	SendTime   string `gorm:"default:null;comment:'发送时间'"`
	CreateTime string `gorm:"default:null;comment:'创建时间'"`
	UpdateTime string `gorm:"default:null;comment:'更新时间'"`
}
