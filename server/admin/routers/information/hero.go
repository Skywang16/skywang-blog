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

var HeroGroup = core.Group("/information", newHeroHandler, regHero, middleware.TokenAuth())

func newHeroHandler(srv information.IInformationHeroService) *heroHandler {
	return &heroHandler{srv: srv}
}

func regHero(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *heroHandler) {
		rg.GET("/hero/list", handle.list)
		rg.GET("/hero/detail", handle.detail)
		rg.POST("/hero/add", handle.add)
		rg.POST("/hero/edit", handle.edit)
		rg.POST("/hero/del", handle.del)
	})
}

type heroHandler struct {
	srv information.IInformationHeroService
}

// list 武器列表
func (dh heroHandler) list(c *gin.Context) {
	var page request.PageReq
	var listReq req.InformationHeroListReq
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
func (dh heroHandler) detail(c *gin.Context) {
	var detailReq req.InformationHeroDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := dh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 武器新增
func (dh heroHandler) add(c *gin.Context) {
	var addReq req.InformationHeroAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Add(addReq))
}

// edit 武器编辑
func (dh heroHandler) edit(c *gin.Context) {
	var editReq req.InformationHeroEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Edit(editReq))
}

// del 武器删除
func (dh heroHandler) del(c *gin.Context) {
	var delReq req.InformationHeroDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Del(delReq.ID))
}
