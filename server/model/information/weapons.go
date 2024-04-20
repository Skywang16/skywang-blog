package information

// InformationWeapons 武器实体
type InformationWeapons struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	Name       string `gorm:"not null;default:'';comment:'武器名称''"`
	Types      string `gorm:"not null;default:'';comment:'武器类型'"`
	Image      string `gorm:"not null;default:'';comment:'武器图片'"`
	Level      string `gorm:"not null;default:'';comment:'武器等级'"`
	Price      string `gorm:"not null;default:'';comment:'武器价格'"`
	Series     string `gorm:"not null;default:'';comment:'武器系列'"`
	Status     int8   `gorm:"not null;default:0;comment:'状态: 0=否, 1=是'"`
	IsDelete   uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime int64  `gorm:"not null;default:0;comment:'删除时间'"`
}

// InformationWeaponsSeries 武器系列实体
type InformationWeaponsSeries struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	Name       string `gorm:"not null;default:'';comment:'武器系列名称''"`
	Remark     string `gorm:"not null;default:'';comment:'备注信息'"`
	IsDisable  uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	Sort       uint16 `gorm:"not null;default:0;comment:'武器系列排序'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}

// InformationHero 武器实体
type InformationHero struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	Name       string `gorm:"not null;default:'';comment:'角色名称''"`
	Types      string `gorm:"not null;default:'';comment:'角色定位'"`
	Image      string `gorm:"not null;default:'';comment:'角色图片'"`
	Status     int8   `gorm:"not null;default:0;comment:'状态: 0=否, 1=是'"`
	IsDelete   uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime int64  `gorm:"not null;default:0;comment:'删除时间'"`
}

// InformationHeroType 武器系列实体
type InformationHeroType struct {
	ID         uint   `gorm:"primarykey;comment:'主键'"`
	Name       string `gorm:"not null;default:'';comment:'角色定位名称''"`
	Remark     string `gorm:"not null;default:'';comment:'备注信息'"`
	IsDisable  uint8  `gorm:"not null;default:0;comment:'是否禁用: 0=否, 1=是'"`
	Sort       uint16 `gorm:"not null;default:0;comment:'角色定位排序'"`
	CreateTime int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
}

// InformationMaps 地图实体
type InformationMaps struct {
	ID            uint   `gorm:"primarykey;comment:'主键'"`
	Name          string `gorm:"not null;default:'';comment:'地图名称''"`
	Image         string `gorm:"not null;default:'';comment:'地图图片'"`
	Locations     string `gorm:"not null;default:'';comment:'地图地点'"`
	PlantQuantity int8   `gorm:"not null;default:2;comment:'包点数量'"`
	Mechanics     string `gorm:"not null;default:'';comment:'地图机制'"`
	Thumbnail     string `gorm:"not null;default:'';comment:'缩略图'"`
	Status        int8   `gorm:"not null;default:0;comment:'状态: 0=否, 1=是'"`
	IsDelete      uint8  `gorm:"not null;default:0;comment:'是否删除: 0=否, 1=是'"`
	CreateTime    int64  `gorm:"autoCreateTime;not null;comment:'创建时间'"`
	UpdateTime    int64  `gorm:"autoUpdateTime;not null;comment:'更新时间'"`
	DeleteTime    int64  `gorm:"not null;default:0;comment:'删除时间'"`
}
