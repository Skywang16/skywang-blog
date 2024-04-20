package information

import (
	"gorm.io/gorm"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/schemas/resp"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/model/information"
)

type IInformationMapsService interface {
	List(page request.PageReq, listReq req.InformationMapsListReq) (res response.PageResp, e error)
	Detail(id uint) (res resp.InformationMapsResp, e error)
	Add(addReq req.InformationMapsAddReq) (e error)
	Edit(editReq req.InformationMapsEditReq) (e error)
	Del(id uint) (e error)
}

// NewInformationMapsService 初始化
func NewInformationMapsService(db *gorm.DB) IInformationMapsService {
	return &informationMapsService{db: db}
}

// informationMapsService 地图服务实现类
type informationMapsService struct {
	db *gorm.DB
}

// List 地图列表
func (newSrv informationMapsService) List(page request.PageReq, listReq req.InformationMapsListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	newModel := newSrv.db.Model(&information.InformationMaps{}).Where("is_delete = ?", 0)
	if listReq.Name != "" {
		newModel = newModel.Where("name like ?", "%"+listReq.Name+"%")
	}
	if listReq.Status >= 0 {
		newModel = newModel.Where("status = ?", listReq.Status)
	}
	// 总数
	var count int64
	err := newModel.Count(&count).Error
	if e = response.CheckErr(err, "InformationMapsList Count err"); e != nil {
		return
	}
	// 数据
	var maps []information.InformationMaps
	err = newModel.Limit(limit).Offset(offset).Order("id desc").Find(&maps).Error
	if e = response.CheckErr(err, "InformationMapsList Find err"); e != nil {
		return
	}
	mapsResps := []resp.InformationMapsResp{}
	response.Copy(&mapsResps, maps)
	// TODO: engine默认local
	for i := 0; i < len(mapsResps); i++ {
		mapsResps[i].Image = maps[i].Image
		mapsResps[i].Name = maps[i].Name
		mapsResps[i].Locations = maps[i].Locations
		mapsResps[i].PlantQuantity = maps[i].PlantQuantity
		mapsResps[i].Mechanics = maps[i].Mechanics
		mapsResps[i].Status = maps[i].Status
	}
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Lists:    mapsResps,
	}, nil
}

// Detail 地图详情
func (newSrv informationMapsService) Detail(id uint) (res resp.InformationMapsResp, e error) {
	var new information.InformationMaps
	err := newSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&new).Error
	if e = response.CheckErrDBNotRecord(err, "地图已不存在!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "Detail First err"); e != nil {
		return
	}
	response.Copy(&res, new)
	return
}

// Add 地图新增
func (newSrv informationMapsService) Add(addReq req.InformationMapsAddReq) (e error) {
	var new information.InformationMaps
	response.Copy(&new, addReq)
	err := newSrv.db.Create(&new).Error
	e = response.CheckErr(err, "Add Create err")
	return
}

// Edit 地图编辑
func (newSrv informationMapsService) Edit(editReq req.InformationMapsEditReq) (e error) {
	var new information.InformationMaps
	err := newSrv.db.Where("id = ? AND is_delete = ?", editReq.ID, 0).Limit(1).First(&new).Error
	// 校验
	if e = response.CheckErrDBNotRecord(err, "地图不存在!"); e != nil {
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

// Del 地图删除
func (newSrv informationMapsService) Del(id uint) (e error) {
	var new information.InformationMaps
	// 根据ID和is_delete状态查询地图
	err := newSrv.db.Where("id = ? AND is_delete = ?", id, 0).Limit(1).First(&new).Error
	if err != nil {
		// 地图不存在
		return response.CheckErrDBNotRecord(err, "地图不存在!")
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
