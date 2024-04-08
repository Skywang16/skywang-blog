package article

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/article"
)

type IArticleNewsService interface {
	All() (res []resp.ArticleNewsResp, e error)
	List(page request.PageReq, listReq req.ArticleNewsListReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.ArticleNewsResp, e error)
	Add(addReq req.ArticleNewsAddReq) (e error)
	Edit(editReq req.ArticleNewsEditReq) (e error)
	Del(id uint) (e error)
}

// NewArticleNewsService 初始化
func NewArticleNewsService(db *gorm.DB) IArticleNewsService {
	return &articleNewsService{db: db}
}

// articleNewsService 资讯服务实现类
type articleNewsService struct {
	db *gorm.DB
}

// All 资讯所有
func (newSrv articleNewsService) All() (res []resp.ArticleNewsResp, e error) {
	var news []article.ArticleNews
	err := newSrv.db.Where("pid > ? AND is_delete = ?", 0, 0).Order("sort desc, id desc").Find(&news).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	res = []resp.ArticleNewsResp{}
	response.Copy(&res, news)
	return
}

// List 资讯列表
func (newSrv articleNewsService) List(page request.PageReq, listReq req.ArticleNewsListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	newModel := newSrv.db.Model(&article.ArticleNews{}).Where("is_delete = ?", 0)
	if listReq.Title != "" {
		newModel = newModel.Where("title like ?", "%"+listReq.Title+"%")
	}
	if listReq.Cid != "" {
		newModel = newModel.Where("cid = ?", listReq.Cid)
	}
	if listReq.Status >= 0 {
		newModel = newModel.Where("status = ?", listReq.Status)
	} else {
		// 当Status为空时，不添加status条件，查询所有状态
	}
	if listReq.Recommended >= 0 {
		newModel = newModel.Where("Recommended = ?", listReq.Recommended)
	} else {
		// 当Recommended为空时，不添加Recommended条件，查询所有Recommended值
	}
	// 总数
	var count int64
	err := newModel.Count(&count).Error
	if e = response.CheckErr(err, "ArticleNewsList Count err"); e != nil {
		return
	}
	// 数据
	var news []article.ArticleNews
	err = newModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&news).Error
	if e = response.CheckErr(err, "ArticleNewsList Find err"); e != nil {
		return
	}
	newsResps := []resp.ArticleNewsResp{}
	response.Copy(&newsResps, news)
	// TODO: engine默认local
	for i := 0; i < len(newsResps); i++ {
		newsResps[i].Image = news[i].Image
		newsResps[i].Title = news[i].Title
		newsResps[i].Desc = news[i].Desc
		newsResps[i].Cid = news[i].Cid
		newsResps[i].Content = news[i].Content
		newsResps[i].PublishTime = news[i].PublishTime
		newsResps[i].Status = news[i].Status
		newsResps[i].Recommended = news[i].Recommended
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    newsResps,
	}, nil
}

// Detail 资讯详情
func (newSrv articleNewsService) Detail(id uint) (res resp.ArticleNewsResp, e error) {
	var new article.ArticleNews
	err := newSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&new).Error
	if e = response.CheckErrDBNotRecord(err, "资讯已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, new)
	return
}

// Add 资讯新增
func (newSrv articleNewsService) Add(addReq req.ArticleNewsAddReq) (e error) {
	var new article.ArticleNews
	response.Copy(&new, addReq)
	err := newSrv.db.Create(&new).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 资讯编辑
func (newSrv articleNewsService) Edit(editReq req.ArticleNewsEditReq) (e error) {
	var new article.ArticleNews
	err := newSrv.db.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&new).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "资讯不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	// 更新
	response.Copy(&new, editReq)
	err = newSrv.db.Save(&new).Error
	e = response.CheckErr(err, "Edit Save err")
	return
}

// Del 资讯删除
func (newSrv articleNewsService) Del(id uint) (e error) {
	var new article.ArticleNews
	// 根据ID和is_delete状态查询资讯
	err := newSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&new).Error
	if err != nil {
		// 资讯不存在
		return response.CheckErrDBNotRecord(err, "资讯不存在!")
	}

	// 设置删除标记
	new.IsDelete = 1
	// 更新数据库中的记录
	err = newSrv.db.Save(&new).Error
	if err != nil {
		// 更新出错
		return response.CheckErr(err, "Del Save err")
	}

	return nil
}
