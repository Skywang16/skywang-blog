package req

// ArticleNewsListReq 资讯列表参数
type ArticleNewsListReq struct {
	Title       string `form:"title"`                               // 标题
	Cid         string `form:"cid"`                                 // 作者
	KeyWords    string `form:"keyWords"`                            // 关键词
	Status      int8   `form:"status" binding:"oneof= -1 0 1"`      // 状态: [-1=全部, 0=否, 1=是]
	Recommended int8   `form:"recommended" binding:"oneof= -1 0 1"` // 是否推荐: [-1=全部, 0=否, 1=是]
}

// ArticleNewsDetailReq 资讯详情参数
type ArticleNewsDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// ArticleNewsAddReq 资讯新增参数
type ArticleNewsAddReq struct {
	Image       string `form:"image" binding:"required"`        // 封面
	Title       string `form:"title" binding:"required"`        // 标题
	Desc        string `form:"desc" binding:"required"`         // 简介
	Cid         string `form:"cid" binding:"required"`          // 作者
	Content     string `form:"content" binding:"required"`      // 内容
	KeyWords    string `form:"keywords" binding:"required"`     // 关键词
	Status      int    `form:"status" binding:"oneof= 0 1"`     // 状态: [-1=全部, 0=否, 1=是]
	Sort        int8   `form:"sort" binding:"gte=0,lte=9999"`   // 排序
	PublishTime string `form:"publishTime" binding:"required"`  // 发布时间
	Recommended int    `form:"recommended" binding:"oneof=0 1"` // 是否推荐: [-1=全部, 0=否, 1=是]
}

// ArticleNewsEditReq 资讯编辑参数
type ArticleNewsEditReq struct {
	ID          uint   `form:"id" binding:"required,gt=0"`          // 主键
	Image       string `form:"image" binding:"required"`            // 封面
	Title       string `form:"title" binding:"required"`            // 标题
	Desc        string `form:"desc" binding:"required"`             // 简介
	Cid         string `form:"cid" binding:"required"`              // 作者
	KeyWords    string `form:"keyWords" binding:"required"`         // 关键词
	Content     string `form:"content" binding:"required"`          // 内容
	Status      int    `form:"status" binding:"oneof= -1 0 1"`      // 状态: [-1=全部, 0=否, 1=是]
	Sort        int8   `form:"sort" binding:"gte=0,lte=9999"`       // 排序
	PublishTime string `form:"publishTime" binding:"required"`      // 发布时间
	Recommended int    `form:"recommended" binding:"oneof= -1 0 1"` // 是否推荐: [-1=全部, 0=否, 1=是]
}

// ArticleNewsDelReq 资讯删除参数
type ArticleNewsDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// ArticleNewsKeyWordsDetailReq 资讯关键词详情参数
type ArticleNewsKeyWordsDetailReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}

// ArticleNewsKeyWordsAddReq 新增资讯关键词参数
type ArticleNewsKeyWordsAddReq struct {
	Name      string `form:"name" binding:"required,min=1,max=30"` // 资讯关键词名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 资讯关键词排序
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 资讯关键词备注
}

// ArticleNewsKeyWordsEditReq 编辑资讯关键词参数
type ArticleNewsKeyWordsEditReq struct {
	ID        uint   `form:"id" binding:"required,gt=0"`           // 主键
	Name      string `form:"name" binding:"required,min=1,max=30"` // 资讯关键词名称
	Sort      int    `form:"sort" binding:"gte=0"`                 // 资讯关键词排序
	IsDisable uint8  `form:"isDisable" binding:"oneof=0 1"`        // 是否禁用: [0=否, 1=是]
	Remark    string `form:"remark" binding:"max=200"`             // 资讯关键词备注
}

// ArticleNewsKeyWordsDelReq 删除资讯关键词参数
type ArticleNewsKeyWordsDelReq struct {
	ID uint `form:"id" binding:"required,gt=0"` // 主键
}
