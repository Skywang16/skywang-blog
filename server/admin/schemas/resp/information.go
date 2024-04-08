package resp

import "likeadmin/core"

// InformationWeaponsResp 资讯列表返回信息
type InformationWeaponsResp struct {
	ID     uint   `json:"id" structs:"id"`         // 武器资讯ID
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
