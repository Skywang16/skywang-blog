package resp

// ArticleNewsResp 资讯列表返回信息
type ArticleNewsResp struct {
	ID          uint   `json:"id" structs:"id"`                   // 资讯ID
	Image       string `json:"image" structs:"image"`             // 封面
	Title       string `json:"title" structs:"title"`             // 标题
	Desc        string `json:"desc" structs:"desc"`               // 简介
	Cid         string `json:"cid" structs:"cid"`                 // 作者
	Content     string `json:"content" structs:"content"`         // 内容
	Sort        int    `json:"sort" binding:"gte=0,lte=9999"`     // 排序
	Status      int8   `json:"status" structs:"status"`           // 状态: [-1=全部, 0=否, 1=是]
	PublishTime string `json:"publishTime" structs:"publishTime"` // 发布时间
	Recommended int8   `json:"recommended" structs:"recommended"` // 是否推荐: [-1=全部, 0=否, 1=是]
}
