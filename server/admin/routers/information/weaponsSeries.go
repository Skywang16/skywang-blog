package information

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/information"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var WeaponsSeriesGroup = core.Group("/information", newWeaponsSeriesHandler, regWeaponsSeries, middleware.TokenAuth())

func newWeaponsSeriesHandler(srv information.IInformationWeaponsSeriesService) *weaponsSeriesHandler {
	return &weaponsSeriesHandler{srv: srv}
}

func regWeaponsSeries(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *weaponsSeriesHandler) {
		rg.GET("/weaponsSeries/list", handle.list)
		rg.GET("/weaponsSeries/detail", handle.detail)
		rg.POST("/weaponsSeries/add", handle.add)
		rg.POST("/weaponsSeries/edit", handle.edit)
		rg.POST("/weaponsSeries/del", handle.del)
	})
}

type weaponsSeriesHandler struct {
	srv information.IInformationWeaponsSeriesService
}

// list 角色列表
func (dh weaponsSeriesHandler) list(c *gin.Context) {
	var page request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := dh.srv.List(page)
	response.CheckAndRespWithData(c, res, err)
}

// detail 武器详情
func (dh weaponsSeriesHandler) detail(c *gin.Context) {
	var detailReq req.InformationWeaponsSeriesDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := dh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 武器新增
func (dh weaponsSeriesHandler) add(c *gin.Context) {
	var addReq req.InformationWeaponsSeriesAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Add(addReq))
}

// edit 武器编辑
func (dh weaponsSeriesHandler) edit(c *gin.Context) {
	var editReq req.InformationWeaponsSeriesEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Edit(editReq))
}

// del 武器删除
func (dh weaponsSeriesHandler) del(c *gin.Context) {
	var delReq req.InformationWeaponsSeriesDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Del(delReq.ID))
}
