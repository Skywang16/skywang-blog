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

type IInformationWeaponsSeriesService interface {
	All() (res []resp.InformationWeaponsSeriesSimpleResp, e error)
	List(page request.PageReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.InformationWeaponsSeriesResp, e error)
	Add(addReq req.InformationWeaponsSeriesAddReq) (e error)
	Edit(editReq req.InformationWeaponsSeriesEditReq) (e error)
	Del(id uint) (e error)
}

// NewInformationWeaponsSeriesService 初始化
func NewInformationWeaponsSeriesService(db *gorm.DB) IInformationWeaponsSeriesService {
	return &informationWeaponsSeriesService{db: db}
}

// informationWeaponsSeriesService 系统系列服务实现类
type informationWeaponsSeriesService struct {
	db *gorm.DB
}

// All 系列所有
func (seriesSrv informationWeaponsSeriesService) All() (res []resp.InformationWeaponsSeriesSimpleResp, e error) {
	var seriess []information.InformationWeaponsSeries
	err := seriesSrv.db.Order("sort desc, id desc").Find(&seriess).Error
	if e = response.CheckErr(err, "All Find err"); e != nil {
		return
	}
	response.Copy(&res, seriess)
	return
}

// List 根据系列ID获取菜单ID
func (seriesSrv informationWeaponsSeriesService) List(page request.PageReq) (res response.PageResp, e error) {
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	seriesModel := seriesSrv.db.Model(&information.InformationWeaponsSeries{})
	var count int64
	err := seriesModel.Count(&count).Error
	if e = response.CheckErr(err, "List Count err"); e != nil {
		return
	}
	var seriess []information.InformationWeaponsSeries
	err = seriesModel.Limit(limit).Offset(offset).Order("sort desc, id desc").Find(&seriess).Error
	if e = response.CheckErr(err, "List Find err"); e != nil {
		return
	}
	var seriesResp []resp.InformationWeaponsSeriesResp
	response.Copy(&seriesResp, seriess)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    seriesResp,
	}, nil
}

// Detail 系列详情
func (seriesSrv informationWeaponsSeriesService) Detail(id uint) (res resp.InformationWeaponsSeriesResp, e error) {
	var series information.InformationWeaponsSeries
	err := seriesSrv.db.Where("id = ?", id).Limit(1).First(&series).Error
	if e = response.CheckErrDBNotRecord(err, "系列已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, series)
	return
}

// Add 新增系列
func (seriesSrv informationWeaponsSeriesService) Add(addReq req.InformationWeaponsSeriesAddReq) (e error) {
	var series information.InformationWeaponsSeries
	if r := seriesSrv.db.Where("name = ?", strings.Trim(addReq.Name, " ")).Limit(1).First(&series); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("系列名称已存在!")
	}
	response.Copy(&series, addReq)
	series.Name = strings.Trim(addReq.Name, " ")
	// 事务
	err := seriesSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Create(&series).Error
		var te error
		if te = response.CheckErr(txErr, "Add Create in tx err"); te != nil {
			return te
		}
		return te
	})
	e = response.CheckErr(err, "Add Transaction err")
	return
}

// Edit 编辑系列
func (seriesSrv informationWeaponsSeriesService) Edit(editReq req.InformationWeaponsSeriesEditReq) (e error) {
	err := seriesSrv.db.Where("id = ?", editReq.ID).Limit(1).First(&information.InformationWeaponsSeries{}).Error
	if e = response.CheckErrDBNotRecord(err, "系列已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Edit First err"); e != nil {
		return
	}
	var series information.InformationWeaponsSeries
	if r := seriesSrv.db.Where("id != ? AND name = ?", editReq.ID, strings.Trim(editReq.Name, " ")).Limit(1).First(&series); r.RowsAffected > 0 {
		return response.AssertArgumentError.Make("系列名称已存在!")
	}
	series.ID = editReq.ID
	seriesMap := structs.Map(editReq)
	delete(seriesMap, "ID")
	seriesMap["Name"] = strings.Trim(editReq.Name, " ")
	// 事务
	err = seriesSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Model(&series).Updates(seriesMap).Error
		var te error
		if te = response.CheckErr(txErr, "Edit Updates in tx err"); te != nil {
			return te
		}
		return te
	})
	e = response.CheckErr(err, "Edit Transaction err")
	return
}

// Del 删除系列
func (seriesSrv informationWeaponsSeriesService) Del(id uint) (e error) {
	err := seriesSrv.db.Where("id = ?", id).Limit(1).First(&information.InformationWeaponsSeries{}).Error
	if e = response.CheckErrDBNotRecord(err, "系列已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Del First err"); e != nil {
		return
	}
	// 事务
	err = seriesSrv.db.Transaction(func(tx *gorm.DB) error {
		txErr := tx.Delete(&information.InformationWeaponsSeries{}, "id = ?", id).Error
		var te error
		if te = response.CheckErr(txErr, "Del Delete in tx err"); te != nil {
			return te
		}
		return nil
	})
	e = response.CheckErr(err, "Del Transaction err")
	return
}
