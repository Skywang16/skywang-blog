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

var HeroTypeGroup = core.Group("/information", newHeroTypeHandler, regHeroType, middleware.TokenAuth())

func newHeroTypeHandler(srv information.IInformationHeroTypeService) *heroTypeHandler {
	return &heroTypeHandler{srv: srv}
}

func regHeroType(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *heroTypeHandler) {
		rg.GET("/heroType/list", handle.list)
		rg.GET("/heroType/detail", handle.detail)
		rg.POST("/heroType/add", handle.add)
		rg.POST("/heroType/edit", handle.edit)
		rg.POST("/heroType/del", handle.del)
	})
}

type heroTypeHandler struct {
	srv information.IInformationHeroTypeService
}

// list 角色列表
func (dh heroTypeHandler) list(c *gin.Context) {
	var page request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := dh.srv.List(page)
	response.CheckAndRespWithData(c, res, err)
}

// detail 武器详情
func (dh heroTypeHandler) detail(c *gin.Context) {
	var detailReq req.InformationHeroTypeDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := dh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 武器新增
func (dh heroTypeHandler) add(c *gin.Context) {
	var addReq req.InformationHeroTypeAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Add(addReq))
}

// edit 武器编辑
func (dh heroTypeHandler) edit(c *gin.Context) {
	var editReq req.InformationHeroTypeEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Edit(editReq))
}

// del 武器删除
func (dh heroTypeHandler) del(c *gin.Context) {
	var delReq req.InformationHeroTypeDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Del(delReq.ID))
}
