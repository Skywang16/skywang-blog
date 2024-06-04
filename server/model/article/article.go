package article

// ArticleNews 资讯实体
type ArticleNews struct {
	ID          uint   `gorm:"primarykey;comment:'主键'"`
	Title       string `gorm:"not null;default:'';comment:'资讯名称''"`
	Desc        string `gorm:"not null;default:'';comment:'资讯简介'"`
	Image       string `gorm:"not null;default:'';comment:'资讯封面'"`
	Cid         string `gorm:"not null;default:'';comment:'资讯作者'"`
	Content     string `gorm:"not null;default:'';comment:'资讯内容'"`
	Status      int8   `gorm:"not null;default:0;comment:'状态: 0=否, 1=是'"`
	KeyWords    string `gorm:"not null;default:'';comment:'关键词'"`
	Sort        uint16 `gorm:"not null;default:0;comment:'排序编号'"`
	Recommended int8   `gorm:"not null;default:0;comment:'是否推荐: 0=否, 1=是'"`
	IsDelete    uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	PublishTime string `gorm:"not null;default:'';comment:'发布时间'"`
	CreateTime  int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime  int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime  int64  `gorm:"not null;default:0;comment:'删除时间'"`
}

// ArticleNewsKeyWords 文章关键词实体
type ArticleNewsKeyWords struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	Name       string `gorm:"not null;default:'';comment:'文章关键词名称''"`
	Remark     string `gorm:"not null;default:'';comment:'备注信息'"`
	IsDisable  uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	Sort       uint16 `gorm:"not null;default:0;comment:'文章关键词排序'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}
