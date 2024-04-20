package req

// SortParams 定义排序参数结构体
type SortParams struct {
	Type  string `json:"type"`      // 排序类型: Level 或 Price
	Field string `json:"sortField"` // 排序字段: desc 或 asc
}

// InformationWeaponsListReq 武器列表参数
type InformationWeaponsListReq struct {
	Name   string     `form:"name"`                                  // 武器名称
	Level  int8       `form:"level" binding:"oneof= -1 0 1 2 3 4"`   // 武器等级
	Status int8       `form:"status" binding:"oneof= -1 0 1"`        // 状态: [-1=全部, 0=否, 1=是]
	Types  int8       `form:"types" binding:"oneof= -1 0 1 2 3 4 5"` // 武器类型
	Sort   SortParams `form:"sort"`                                  // 排序对象
}

// InformationWeaponsDetailReq 武器详情参数
type InformationWeaponsDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// InformationWeaponsAddReq 武器新增参数
type InformationWeaponsAddReq struct {
	Image  string `form:"image" binding:"required"`              // 武器图片
	Name   string `form:"name" binding:"required"`               // 武器名称
	Types  int8   `form:"types" binding:"oneof= -1 0 1 2 3 4 5"` // 武器类型
	Level  int8   `form:"level" binding:"oneof= -1 0 1 2 3 4"`   // 武器等级
	Price  string `form:"price" binding:"required"`              // 武器价格
	Series string `form:"series" binding:"required"`             // 武器系列
	Status int    `form:"status" binding:"oneof=-1 0 1"`         // 武器状态: [-1=全部, 0=否, 1=是]
}

// InformationWeaponsEditReq 武器编辑参数
type InformationWeaponsEditReq struct {
	ID     uint   `form:"id" binding:"required,gt=0"`            // 武器主键
	Image  string `form:"image" binding:"required"`              // 武器图片
	Types  int8   `form:"types" binding:"oneof= -1 0 1 2 3 4 5"` // 武器类型
	Name   string `form:"name" binding:"required"`               // 武器名称
	Level  int8   `form:"level" binding:"oneof= -1 0 1 2 3 4"`   // 武器等级
	Price  string `form:"price" binding:"required"`              // 武器价格
	Series string `form:"series" binding:"required"`             // 武器系列
	Status int    `form:"status" binding:"oneof=-1 0 1"`         // 武器状态: [-1=全部, 0=否, 1=是]
}

// InformationWeaponsDelReq 武器删除参数
type InformationWeaponsDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// InformationWeaponsSeriesDetailReq 武器系列详情参数
type InformationWeaponsSeriesDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// InformationWeaponsSeriesAddReq 新增武器系列参数
type InformationWeaponsSeriesAddReq struct {
	Name      string `form:"name" binding:"required,min=1,max=30"` // 武器系列名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 武器系列排序
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 武器系列备注
}

// InformationWeaponsSeriesEditReq 编辑武器系列参数
type InformationWeaponsSeriesEditReq struct {
	ID        uint   `form:"id" binding:"required,gt=0"`           // 主键
	Name      string `form:"name" binding:"required,min=1,max=30"` // 武器系列名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 武器系列排序
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 武器系列备注
}

// InformationWeaponsSeriesDelReq 删除角色定位参数
type InformationWeaponsSeriesDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// InformationHeroListReq 角色列表参数
type InformationHeroListReq struct {
	Name   string `form:"name"`                                  // 角色名称
	Status int8   `form:"status" binding:"oneof= -1 0 1"`        // 状态: [-1=全部, 0=否, 1=是]
	Types  int8   `form:"types" binding:"oneof= -1 0 1 2 3 4 5"` // 角色定位
}

// InformationHeroDetailReq 角色详情参数
type InformationHeroDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// InformationHeroAddReq 角色新增参数
type InformationHeroAddReq struct {
	Image  string `form:"image" binding:"required"`              // 角色图片
	Name   string `form:"name" binding:"required"`               // 角色名称
	Types  int8   `form:"types" binding:"oneof= -1 0 1 2 3 4 5"` // 角色定位
	Status int    `form:"status" binding:"oneof=-1 0 1"`         // 角色状态: [-1=全部, 0=否, 1=是]
}

// InformationHeroEditReq 角色编辑参数
type InformationHeroEditReq struct {
	ID     uint   `form:"id" binding:"required,gt=0"`            // 角色主键
	Image  string `form:"image" binding:"required"`              // 角色图片
	Types  int8   `form:"types" binding:"oneof= -1 0 1 2 3 4 5"` // 角色定位
	Name   string `form:"name" binding:"required"`               // 角色名称
	Status int    `form:"status" binding:"oneof=-1 0 1"`         // 角色状态: [-1=全部, 0=否, 1=是]
}

// InformationHeroDelReq 角色删除参数
type InformationHeroDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// InformationHeroTypeDetailReq 角色定位详情参数
type InformationHeroTypeDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// InformationHeroTypeAddReq 新增角色定位参数
type InformationHeroTypeAddReq struct {
	Name      string `form:"name" binding:"required,min=1,max=30"` // 角色定位名称
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 角色定位备注
}

// InformationHeroTypeEditReq 编辑角色定位参数
type InformationHeroTypeEditReq struct {
	ID        uint   `form:"id" binding:"required,gt=0"`           // 主键
	Name      string `form:"name" binding:"required,min=1,max=30"` // 角色定位名称
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 角色定位备注
}

// InformationHeroTypeDelReq 删除角色定位参数
type InformationHeroTypeDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// InformationMapsListReq 地图列表参数
type InformationMapsListReq struct {
	Name   string `form:"name"`                           // 地图名称
	Status int8   `form:"status" binding:"oneof= -1 0 1"` // 状态: [-1=全部, 0=否, 1=是]
}

// InformationMapsDetailReq 地图详情参数
type InformationMapsDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// InformationMapsAddReq 地图新增参数
type InformationMapsAddReq struct {
	Name          string `form:"name" binding:"required"`                   // 地图名称
	Image         string `form:"image" binding:"required"`                  // 地图图片
	Locations     string `form:"locations"`                                 // 地图地点
	PlantQuantity int8   `form:"plantQuantity" binding:"oneof= -1 1 2 3 4"` // 包点数量
	Mechanics     string `form:"mechanics"`                                 // 地图机制
	Thumbnail     string `form:"thumbnail"`                                 // 缩略图
	Status        int    `form:"status" binding:"oneof=-1 0 1"`             // 地图状态: [-1=全部, 0=否, 1=是]
}

// InformationMapsEditReq 地图编辑参数
type InformationMapsEditReq struct {
	ID            uint   `form:"id" binding:"required,gt=0"`                // 地图主键
	Name          string `form:"name" binding:"required"`                   // 地图名称
	Image         string `form:"image" binding:"required"`                  // 地图图片
	Locations     string `form:"locations"`                                 // 地图地点
	PlantQuantity int8   `form:"plantQuantity" binding:"oneof= -1 1 2 3 4"` // 包点数量
	Mechanics     string `form:"mechanics"`                                 // 地图机制
	Thumbnail     string `form:"thumbnail"`                                 // 缩略图
	Status        int    `form:"status" binding:"oneof=-1 0 1"`             // 地图状态: [-1=全部, 0=否, 1=是]
}

// InformationMapsDelReq 地图删除参数
type InformationMapsDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}
