package article

// Article 文章
type Article struct {
	ArticleId  uint   `gorm:"primarykey;comment:'主键'"`
	Title      string `gorm:"not null;default:'';comment:'文章名称'"`
	ClassId    uint   `gorm:"not null;default:0;comment:'文章分类'"`
	Audio      string `gorm:"default:'';comment:'音频'"`
	Content    string `gorm:"not null;default:'';comment:'文章内容'"`
	ReadNum    uint   `gorm:"not null;default:0;comment:'阅读次数'"`
	Rank       uint8  `gorm:"not null;default:0;comment:'权重'"`
	DelFlag    uint8  `gorm:"not null;default:0;comment:'删除标识：0.否 1.是'"`
	Status     uint   `gorm:"not null;default:1;comment:'是否显示：0.否 1.是'"`
	CreateTime string `gorm:"comment:'创建时间'"`
	UpdateTime string `gorm:"comment:'更新时间'"`
}

// ArticleClass 文章分类
type ArticleClass struct {
	ClassId    uint   `gorm:"primarykey;comment:'主键'"`
	ClassName  string `gorm:"not null;default:'';comment:'分类名称'"`
	DelFlag    uint8  `gorm:"not null;default:0;comment:'删除标识：0.否 1.是'"`
	Status     uint   `gorm:"not null;default:1;comment:'是否显示：0.否 1.是'"`
	CreateTime string `gorm:"comment:'创建时间'"`
	UpdateTime string `gorm:"comment:'更新时间'"`
}

// ArticleRecord 文章分类
type ArticleRecord struct {
	RecordId   uint   `gorm:"primarykey;comment:'主键'"`
	UserId     uint   `gorm:"not null;comment:'用户ID'"`
	ArticleId  uint   `gorm:"not null;comment:'文章ID'"`
	CreateTime string `gorm:"comment:'创建时间'"`
}
