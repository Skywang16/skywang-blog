package information

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/information"
	"strconv"
	"strings"
)

type IInformationWeaponsService interface {
	List(page request.PageReq, listReq req.InformationWeaponsListReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.InformationWeaponsResp, e error)
	Add(addReq req.InformationWeaponsAddReq) (e error)
	Edit(editReq req.InformationWeaponsEditReq) (e error)
	Del(id uint) (e error)
}

// NewInformationWeaponsService 初始化
func NewInformationWeaponsService(db *gorm.DB) IInformationWeaponsService {
	return &informationWeaponsService{db: db}
}

// informationWeaponsService 武器服务实现类
type informationWeaponsService struct {
	db *gorm.DB
}

// List 武器列表
func (newSrv informationWeaponsService) List(page request.PageReq, listReq req.InformationWeaponsListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	newModel := newSrv.db.Model(&information.InformationWeapons{}).Where("is_delete = ?", 0)
	if listReq.Name != "" {
		newModel = newModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.Level >= 0 {
		newModel = newModel.Where("level = ?", listReq.Level)
	}
	if listReq.Types >= 0 {
		newModel = newModel.Where("types = ?", listReq.Types)
	}
	if listReq.Status >= 0 {
		newModel = newModel.Where("status = ?", listReq.Status)
	}
	// 构建排序条件
	orderField := "id desc" // 默认排序字段
	if listReq.Sort.Type == "Level" {
		orderField = "level"
	} else if listReq.Sort.Type == "Price" {
		orderField = "price"
	}

	if listReq.Sort.Field == "asc" {
		if !strings.Contains(orderField, "asc") && !strings.Contains(orderField, "desc") {
			orderField += " asc"
		}
	} else {
		if !strings.Contains(orderField, "asc") && !strings.Contains(orderField, "desc") {
			orderField += " desc"
		}
	}

	// 总数
	var count int64
	err := newModel.Count(&count).Error
	if e = response.CheckErr(err, "InformationWeaponsList Count err"); e != nil {
		return
	}
	// 数据
	var weapons []information.InformationWeapons
	err = newModel.Limit(limit).Offset(offset).Order(orderField).Find(&weapons).Error
	if e = response.CheckErr(err, "InformationWeaponsList Find err"); e != nil {
		return
	}
	weaponsResps := []resp.InformationWeaponsResp{}
	response.Copy(&weaponsResps, weapons)
	// TODO: engine默认local
	for i := 0; i < len(weaponsResps); i++ {
		weaponsResps[i].Image = weapons[i].Image
		weaponsResps[i].Name = weapons[i].Name
		levelInt, err := strconv.Atoi(weapons[i].Level)
		if err != nil {
			// 处理转换错误
		}
		typesInt, err := strconv.Atoi(weapons[i].Types)
		if err != nil {
			// 处理转换错误
		}
		weaponsResps[i].Types = int8(typesInt)
		weaponsResps[i].Level = int8(levelInt)
		weaponsResps[i].Status = weapons[i].Status
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    weaponsResps,
	}, nil
}

// Detail 武器详情
func (newSrv informationWeaponsService) Detail(id uint) (res resp.InformationWeaponsResp, e error) {
	var new information.InformationWeapons
	err := newSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&new).Error
	if e = response.CheckErrDBNotRecord(err, "武器已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, new)
	return
}

// Add 武器新增
func (newSrv informationWeaponsService) Add(addReq req.InformationWeaponsAddReq) (e error) {
	var new information.InformationWeapons
	response.Copy(&new, addReq)
	err := newSrv.db.Create(&new).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 武器编辑
func (newSrv informationWeaponsService) Edit(editReq req.InformationWeaponsEditReq) (e error) {
	var new information.InformationWeapons
	err := newSrv.db.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&new).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "武器不存在!"); e != nil {
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

// Del 武器删除
func (newSrv informationWeaponsService) Del(id uint) (e error) {
	var new information.InformationWeapons
	// 根据ID和is_delete状态查询武器
	err := newSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&new).Error
	if err != nil {
		// 武器不存在
		return response.CheckErrDBNotRecord(err, "武器不存在!")
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
