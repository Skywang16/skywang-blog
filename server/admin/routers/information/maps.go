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

var MapsGroup = core.Group("/information", newMapsHandler, regMaps, middleware.TokenAuth())

func newMapsHandler(srv information.IInformationMapsService) *mapsHandler {
	return &mapsHandler{srv: srv}
}

func regMaps(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *mapsHandler) {
		rg.GET("/maps/list", handle.list)
		rg.GET("/maps/detail", handle.detail)
		rg.POST("/maps/add", handle.add)
		rg.POST("/maps/edit", handle.edit)
		rg.POST("/maps/del", handle.del)
	})
}

type mapsHandler struct {
	srv information.IInformationMapsService
}

// list 地图列表
func (dh mapsHandler) list(c *gin.Context) {
	var page request.PageReq
	var listReq req.InformationMapsListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := dh.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 地图详情
func (dh mapsHandler) detail(c *gin.Context) {
	var detailReq req.InformationMapsDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := dh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 地图新增
func (dh mapsHandler) add(c *gin.Context) {
	var addReq req.InformationMapsAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Add(addReq))
}

// edit 地图编辑
func (dh mapsHandler) edit(c *gin.Context) {
	var editReq req.InformationMapsEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Edit(editReq))
}

// del 地图删除
func (dh mapsHandler) del(c *gin.Context) {
	var delReq req.InformationMapsDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Del(delReq.ID))
}
