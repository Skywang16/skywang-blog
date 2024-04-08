package article

import (
	"github.com/gin-gonic/gin"
	"likeadmin/admin/schemas/req"
	"likeadmin/admin/service/article"
	"likeadmin/core"
	"likeadmin/core/request"
	"likeadmin/core/response"
	"likeadmin/middleware"
	"likeadmin/util"
)

var NewsGroup = core.Group("/article", newNewsHandler, regNews, middleware.TokenAuth())

func newNewsHandler(srv article.IArticleNewsService) *newsHandler {
	return &newsHandler{srv: srv}
}

func regNews(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *newsHandler) {
		rg.GET("/news/all", handle.all)
		rg.GET("/news/list", handle.list)
		rg.GET("/news/detail", handle.detail)
		rg.POST("/news/add", handle.add)
		rg.POST("/news/edit", handle.edit)
		rg.POST("/news/del", handle.del)
	})
}

type newsHandler struct {
	srv article.IArticleNewsService
}

// all 资讯所有
func (dh newsHandler) all(c *gin.Context) {
	res, err := dh.srv.All()
	response.CheckAndRespWithData(c, res, err)
}

// list 资讯列表
func (dh newsHandler) list(c *gin.Context) {
	var page request.PageReq
	var listReq req.ArticleNewsListReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &listReq)) {
		return
	}
	res, err := dh.srv.List(page, listReq)
	response.CheckAndRespWithData(c, res, err)
}

// detail 资讯详情
func (dh newsHandler) detail(c *gin.Context) {
	var detailReq req.ArticleNewsDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := dh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 资讯新增
func (dh newsHandler) add(c *gin.Context) {
	var addReq req.ArticleNewsAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Add(addReq))
}

// edit 资讯编辑
func (dh newsHandler) edit(c *gin.Context) {
	var editReq req.ArticleNewsEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Edit(editReq))
}

// del 资讯删除
func (dh newsHandler) del(c *gin.Context) {
	var delReq req.ArticleNewsDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Del(delReq.ID))
}
