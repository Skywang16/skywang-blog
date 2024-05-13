package resp

import "likeadmin/core"

// InformationWeaponsResp 武器列表返回信息
type InformationWeaponsResp struct {
	ID     uint   `json:"id" structs:"id"`         // 武器ID
	Image  string `json:"image" structs:"image"`   // 武器图片
	Types  int8   `json:"types" structs:"types"`   // 武器类型
	Name   string `json:"name" structs:"name"`     // 武器名称
	Level  int8   `json:"level" structs:"level"`   // 武器等级
	Price  string `json:"price" structs:"price"`   // 武器价格
	Series string `json:"series" structs:"series"` // 武器系列
	Status int8   `json:"status" structs:"status"` // 武器状态: [-1=全部, 0=否, 1=是]
}

// InformationWeaponsSeriesSimpleResp 武器系列返回简单信息
type InformationWeaponsSeriesSimpleResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 角色名称
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

// InformationWeaponsSeriesResp 武器系列返回信息
type InformationWeaponsSeriesResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 角色名称
	Remark     string      `json:"remark" structs:"remark"`         // 角色备注
	Sort       uint16      `json:"sort" structs:"sort"`             // 角色排序
	IsDisable  uint8       `json:"isDisable" structs:"isDisable"`   // 是否禁用: [0=否, 1=是]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

// InformationHeroResp 角色列表返回信息
type InformationHeroResp struct {
	ID     uint   `json:"id" structs:"id"`         // 角色ID
	Image  string `json:"image" structs:"image"`   // 角色图片
	Types  int8   `json:"types" structs:"types"`   // 角色定位
	Name   string `json:"name" structs:"name"`     // 角色名称
	Status int8   `json:"status" structs:"status"` // 角色状态: [-1=全部, 0=否, 1=是]
}

// InformationHeroTypeSimpleResp 角色定位返回简单信息
type InformationHeroTypeSimpleResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 角色定位名称
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

// InformationHeroTypeResp 角色定位返回信息
type InformationHeroTypeResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 角色定位名称
	Remark     string      `json:"remark" structs:"remark"`         // 角色定位备注
	Sort       uint16      `json:"sort" structs:"sort"`             // 角色定位排序
	IsDisable  uint8       `json:"isDisable" structs:"isDisable"`   // 是否禁用: [0=否, 1=是]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

// InformationMapsResp 地图列表返回信息
type InformationMapsResp struct {
	ID            uint   `json:"id" structs:"id"`                       // 地图ID
	Name          string `json:"name" structs:"name"`                   // 地图名称
	Image         string `json:"image" structs:"image"`                 // 地图图片
	Locations     string `json:"locations" structs:"locations"`         // 地图地点
	PlantQuantity int8   `json:"plantQuantity" structs:"plantQuantity"` // 包点数量
	Mechanics     string `json:"mechanics" structs:"mechanics"`         // 地图机制
	Status        int8   `json:"status" structs:"status"`               // 地图状态: [-1=全部, 0=否, 1=是]
}
