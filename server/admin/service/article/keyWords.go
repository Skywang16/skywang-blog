package article

import (
	"github.com/fatih/structs"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/article"
	"strings"
)

type IArticleNewsKeyWordsService interface {
	All() (res []resp.ArticleNewsKeyWordsSimpleResp, e error)
	List(page request.PageReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.ArticleNewsKeyWordsResp, e error)
	Add(addReq req.ArticleNewsKeyWordsAddReq) (e error)
	Edit(editReq req.ArticleNewsKeyWordsEditReq) (e error)
	Del(id uint) (e error)
}

// NewArticleNewsKeyWordsService 初始化
func NewArticleNewsKeyWordsService(db *gorm.DB) IArticleNewsKeyWordsService {
	return &articleNewsKeyWordsService{db: db}
}

// articleNewsKeyWordsService 系统文章关键词服务实现类
type articleNewsKeyWordsService struct {
	db *gorm.DB
}

// All 文章关键词所有
func (keyWordsSrv articleNewsKeyWordsService) All() (res []resp.ArticleNewsKeyWordsSimpleResp, e error) {
	var keyWordss []article.ArticleNewsKeyWords
	err := keyWordsSrv.db.Order("sort desc, id desc").Find(&keyWordss).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	response.Copy(&res, keyWordss)
	return
}

// List 根据文章关键词ID获取菜单ID
func (keyWordsSrv articleNewsKeyWordsService) List(page request.PageReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	keyWordsModel := keyWordsSrv.db.Model(&article.ArticleNewsKeyWords{})
	var count int64
	err := keyWordsModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var keyWordss []article.ArticleNewsKeyWords
	err = keyWordsModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&keyWordss).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var keyWordsResp []resp.ArticleNewsKeyWordsResp
	response.Copy(&keyWordsResp, keyWordss)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    keyWordsResp,
	}, nil
}

// Detail 文章关键词详情
func (keyWordsSrv articleNewsKeyWordsService) Detail(id uint) (res resp.ArticleNewsKeyWordsResp, e error) {
	var keyWords article.ArticleNewsKeyWords
	err := keyWordsSrv.db.Where("id = ?", id).Limit(1).First(&keyWords).Error
	if e = response.CheckErrDBNotRecord(err, "文章关键词已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, keyWords)
	return
}

// Add 新增文章关键词
func (keyWordsSrv articleNewsKeyWordsService) Add(addReq req.ArticleNewsKeyWordsAddReq) (e error) {
	var keyWords article.ArticleNewsKeyWords
	if r := keyWordsSrv.db.Where("name = ?", strings.Trim(addReq.Name, " ")).Limit(1).First(&keyWords); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("文章关键词名称已存在!")
	}
	response.Copy(&keyWords, addReq)
	keyWords.Name = strings.Trim(addReq.Name, " ")
	// 事务
	err := keyWordsSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Create(&keyWords).Error
		var te error
		if te = response.CheckErr(txErr, "Add Create in tx err"); te != nil {
			return te
		}
		return te
	})
	e = response.CheckErr(err, "Add Transaction err")
	return
}

// Edit 编辑文章关键词
func (keyWordsSrv articleNewsKeyWordsService) Edit(editReq req.ArticleNewsKeyWordsEditReq) (e error) {
	err := keyWordsSrv.db.Where("id = ?", editReq.ID).Limit(1).First(&article.ArticleNewsKeyWords{}).Error
	if e = response.CheckErrDBNotRecord(err, "文章关键词已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	var keyWords article.ArticleNewsKeyWords
	if r := keyWordsSrv.db.Where("id != ? AND name = ?", editReq.ID, strings.Trim(editReq.Name, " ")).Limit(1).First(&keyWords); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("文章关键词名称已存在!")
	}
	keyWords.ID = editReq.ID
	keyWordsMap := structs.Map(editReq)
	delete(keyWordsMap, "ID")
	keyWordsMap["Name"] = strings.Trim(editReq.Name, " ")
	// 事务
	err = keyWordsSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&keyWords).Updates(keyWordsMap).Error
		var te error
		if te = response.CheckErr(txErr, "Edit Updates in tx err"); te != nil {
			return te
		}
		return te
	})
	e = response.CheckErr(err, "Edit Transaction err")
	return
}

// Del 删除文章关键词
func (keyWordsSrv articleNewsKeyWordsService) Del(id uint) (e error) {
	err := keyWordsSrv.db.Where("id = ?", id).Limit(1).First(&article.ArticleNewsKeyWords{}).Error
	if e = response.CheckErrDBNotRecord(err, "文章关键词已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	// 事务
	err = keyWordsSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&article.ArticleNewsKeyWords{}, "id = ?", id).Error
		var te error
		if te = response.CheckErr(txErr, "Del Delete in tx err"); te != nil {
			return te
		}
		return nil
	})
	e = response.CheckErr(err, "Del Transaction err")
	return
}
