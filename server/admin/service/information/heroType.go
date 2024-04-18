package information

import (
	"github.com/fatih/structs"
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/information"
	"strings"
)

type IInformationHeroTypeService interface {
	All() (res []resp.InformationHeroTypeSimpleResp, e error)
	List(page request.PageReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.InformationHeroTypeResp, e error)
	Add(addReq req.InformationHeroTypeAddReq) (e error)
	Edit(editReq req.InformationHeroTypeEditReq) (e error)
	Del(id uint) (e error)
}

// NewInformationHeroTypeService 初始化
func NewInformationHeroTypeService(db *gorm.DB) IInformationHeroTypeService {
	return &informationHeroTypeService{db: db}
}

// informationHeroTypeService 系统角色定位服务实现类
type informationHeroTypeService struct {
	db *gorm.DB
}

// All 角色定位所有
func (type1Srv informationHeroTypeService) All() (res []resp.InformationHeroTypeSimpleResp, e error) {
	var type1s []information.InformationHeroType
	err := type1Srv.db.Order("sort desc, id desc").Find(&type1s).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	response.Copy(&res, type1s)
	return
}

// List 根据角色定位ID获取菜单ID
func (type1Srv informationHeroTypeService) List(page request.PageReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	type1Model := type1Srv.db.Model(&information.InformationHeroType{})
	var count int64
	err := type1Model.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var type1s []information.InformationHeroType
	err = type1Model.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&type1s).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var type1Resp []resp.InformationHeroTypeResp
	response.Copy(&type1Resp, type1s)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    type1Resp,
	}, nil
}

// Detail 角色定位详情
func (type1Srv informationHeroTypeService) Detail(id uint) (res resp.InformationHeroTypeResp, e error) {
	var type1 information.InformationHeroType
	err := type1Srv.db.Where("id = ?", id).Limit(1).First(&type1).Error
	if e = response.CheckErrDBNotRecord(err, "角色定位已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, type1)
	return
}

// Add 新增角色定位
func (type1Srv informationHeroTypeService) Add(addReq req.InformationHeroTypeAddReq) (e error) {
	var type1 information.InformationHeroType
	if r := type1Srv.db.Where("name = ?", strings.Trim(addReq.Name, " ")).Limit(1).First(&type1); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("角色定位名称已存在!")
	}
	response.Copy(&type1, addReq)
	type1.Name = strings.Trim(addReq.Name, " ")
	// 事务
	err := type1Srv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Create(&type1).Error
		var te error
		if te = response.CheckErr(txErr, "Add Create in tx err"); te != nil {
			return te
		}
		return te
	})
	e = response.CheckErr(err, "Add Transaction err")
	return
}

// Edit 编辑角色定位
func (type1Srv informationHeroTypeService) Edit(editReq req.InformationHeroTypeEditReq) (e error) {
	err := type1Srv.db.Where("id = ?", editReq.ID).Limit(1).First(&information.InformationHeroType{}).Error
	if e = response.CheckErrDBNotRecord(err, "角色定位已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	var type1 information.InformationHeroType
	if r := type1Srv.db.Where("id != ? AND name = ?", editReq.ID, strings.Trim(editReq.Name, " ")).Limit(1).First(&type1); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("角色定位名称已存在!")
	}
	type1.ID = editReq.ID
	type1Map := structs.Map(editReq)
	delete(type1Map, "ID")
	type1Map["Name"] = strings.Trim(editReq.Name, " ")
	// 事务
	err = type1Srv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&type1).Updates(type1Map).Error
		var te error
		if te = response.CheckErr(txErr, "Edit Updates in tx err"); te != nil {
			return te
		}
		return te
	})
	e = response.CheckErr(err, "Edit Transaction err")
	return
}

// Del 删除角色定位
func (type1Srv informationHeroTypeService) Del(id uint) (e error) {
	err := type1Srv.db.Where("id = ?", id).Limit(1).First(&information.InformationHeroType{}).Error
	if e = response.CheckErrDBNotRecord(err, "角色定位已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	// 事务
	err = type1Srv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&information.InformationHeroType{}, "id = ?", id).Error
		var te error
		if te = response.CheckErr(txErr, "Del Delete in tx err"); te != nil {
			return te
		}
		return nil
	})
	e = response.CheckErr(err, "Del Transaction err")
	return
}
