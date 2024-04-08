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

var WeaponsGroup = core.Group("/information", newWeaponsHandler, regWeapons, middleware.TokenAuth())

func newWeaponsHandler(srv information.IInformationWeaponsService) *weaponsHandler {
	return &weaponsHandler{srv: srv}
}

func regWeapons(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *weaponsHandler) {
		rg.GET("/weapons/list", handle.list)
		rg.GET("/weapons/detail", handle.detail)
		rg.POST("/weapons/add", handle.add)
		rg.POST("/weapons/edit", handle.edit)
		rg.POST("/weapons/del", handle.del)
	})
}

type weaponsHandler struct {
	srv information.IInformationWeaponsService
}

// list 武器列表
func (dh weaponsHandler) list(c *gin.Context) {
	var page request.PageReq
	var listReq req.InformationWeaponsListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := dh.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 武器详情
func (dh weaponsHandler) detail(c *gin.Context) {
	var detailReq req.InformationWeaponsDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := dh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 武器新增
func (dh weaponsHandler) add(c *gin.Context) {
	var addReq req.InformationWeaponsAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Add(addReq))
}

// edit 武器编辑
func (dh weaponsHandler) edit(c *gin.Context) {
	var editReq req.InformationWeaponsEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Edit(editReq))
}

// del 武器删除
func (dh weaponsHandler) del(c *gin.Context) {
	var delReq req.InformationWeaponsDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Del(delReq.ID))
}
