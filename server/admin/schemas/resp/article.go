package resp

import "likeadmin/core"

// ArticleNewsResp 资讯列表返回信息
type ArticleNewsResp struct {
	ID          uint   `json:"id" structs:"id"`                   // 资讯ID
	Image       string `json:"image" structs:"image"`             // 封面
	Title       string `json:"title" structs:"title"`             // 标题
	Desc        string `json:"desc" structs:"desc"`               // 简介
	Cid         string `json:"cid" structs:"cid"`                 // 作者
	KeyWords    string `json:"keyWords" structs:"keyWords"`       // 关键词
	Content     string `json:"content" structs:"content"`         // 内容
	Sort        int    `json:"sort" binding:"gte=0,lte=9999"`     // 排序
	Status      int8   `json:"status" structs:"status"`           // 状态: [-1=全部, 0=否, 1=是]
	PublishTime string `json:"publishTime" structs:"publishTime"` // 发布时间
	Recommended int8   `json:"recommended" structs:"recommended"` // 是否推荐: [-1=全部, 0=否, 1=是]
}

// ArticleNewsKeyWordsSimpleResp 文章关键词返回简单信息
type ArticleNewsKeyWordsSimpleResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 文章名称
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}

// ArticleNewsKeyWordsResp 文章关键词返回信息
type ArticleNewsKeyWordsResp struct {
	ID         uint        `json:"id" structs:"id"`                 // 主键
	Name       string      `json:"name" structs:"name"`             // 文章名称
	Remark     string      `json:"remark" structs:"remark"`         // 文章备注
	Sort       uint16      `json:"sort" structs:"sort"`             // 文章排序
	IsDisable  uint8       `json:"isDisable" structs:"isDisable"`   // 是否禁用: [0=否, 1=是]
	CreateTime core.TsTime `json:"createTime" structs:"createTime"` // 创建时间
	UpdateTime core.TsTime `json:"updateTime" structs:"updateTime"` // 更新时间
}
