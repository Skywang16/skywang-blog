package information

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/information"
)

type IInformationHeroService interface {
	List(page request.PageReq, listReq req.InformationHeroListReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.InformationHeroResp, e error)
	Add(addReq req.InformationHeroAddReq) (e error)
	Edit(editReq req.InformationHeroEditReq) (e error)
	Del(id uint) (e error)
}

// NewInformationHeroService 初始化
func NewInformationHeroService(db *gorm.DB) IInformationHeroService {
	return &informationHeroService{db: db}
}

// informationHeroService 角色服务实现类
type informationHeroService struct {
	db *gorm.DB
}

// List 角色列表
func (newSrv informationHeroService) List(page request.PageReq, listReq req.InformationHeroListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	newModel := newSrv.db.Model(&information.InformationHero{}).Where("is_delete = ?", 0)
	if listReq.Name != "" {
		newModel = newModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.Types >= 0 {
		newModel = newModel.Where("types = ?", listReq.Types)
	}
	if listReq.Status >= 0 {
		newModel = newModel.Where("status = ?", listReq.Status)
	}

	// 总数
	var count int64
	err := newModel.Count(&count).Error
	if e = response.CheckErr(err, "InformationHeroList Count err"); e != nil {
		return
	}
	// 数据
	var hero []information.InformationHero
	err = newModel.Limit(limit).Offset(offset).Find(&hero).Error
	if e = response.CheckErr(err, "InformationHeroList Find err"); e != nil {
		return
	}
	heroResps := []resp.InformationHeroResp{}
	response.Copy(&heroResps, hero)
	// TODO: engine默认local
	for i := 0; i < len(heroResps); i++ {
		heroResps[i].Name = hero[i].Name
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    heroResps,
	}, nil
}

// Detail 角色详情
func (newSrv informationHeroService) Detail(id uint) (res resp.InformationHeroResp, e error) {
	var new information.InformationHero
	err := newSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&new).Error
	if e = response.CheckErrDBNotRecord(err, "角色已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, new)
	return
}

// Add 角色新增
func (newSrv informationHeroService) Add(addReq req.InformationHeroAddReq) (e error) {
	var new information.InformationHero
	response.Copy(&new, addReq)
	err := newSrv.db.Create(&new).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 角色编辑
func (newSrv informationHeroService) Edit(editReq req.InformationHeroEditReq) (e error) {
	var new information.InformationHero
	err := newSrv.db.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&new).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "角色不存在!"); e != nil {
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

// Del 角色删除
func (newSrv informationHeroService) Del(id uint) (e error) {
	var new information.InformationHero
	// 根据ID和is_delete状态查询角色
	err := newSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&new).Error
	if err != nil {
		// 角色不存在
		return response.CheckErrDBNotRecord(err, "角色不存在!")
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
