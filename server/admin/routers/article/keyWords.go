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

var KeyWordsGroup = core.Group("/article", newNewsKeyWordsHandler, regNewsKeyWords, middleware.TokenAuth())

func newNewsKeyWordsHandler(srv article.IArticleNewsKeyWordsService) *keyWordsHandler {
	return &keyWordsHandler{srv: srv}
}

func regNewsKeyWords(rg *gin.RouterGroup, group *core.GroupBase) error {
	return group.Reg(func(handle *keyWordsHandler) {
		rg.GET("/keyWords/list", handle.list)
		rg.GET("/keyWords/detail", handle.detail)
		rg.POST("/keyWords/add", handle.add)
		rg.POST("/keyWords/edit", handle.edit)
		rg.POST("/keyWords/del", handle.del)
	})
}

type keyWordsHandler struct {
	srv article.IArticleNewsKeyWordsService
}

// list 关键词列表
func (dh keyWordsHandler) list(c *gin.Context) {
	var page request.PageReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &page)) {
		return
	}
	res, err := dh.srv.List(page)
	response.CheckAndRespWithData(c, res, err)
}

// detail 关键词详情
func (dh keyWordsHandler) detail(c *gin.Context) {
	var detailReq req.ArticleNewsKeyWordsDetailReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyQuery(c, &detailReq)) {
		return
	}
	res, err := dh.srv.Detail(detailReq.ID)
	response.CheckAndRespWithData(c, res, err)
}

// add 关键词新增
func (dh keyWordsHandler) add(c *gin.Context) {
	var addReq req.ArticleNewsKeyWordsAddReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &addReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Add(addReq))
}

// edit 关键词编辑
func (dh keyWordsHandler) edit(c *gin.Context) {
	var editReq req.ArticleNewsKeyWordsEditReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &editReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Edit(editReq))
}

// del 关键词删除
func (dh keyWordsHandler) del(c *gin.Context) {
	var delReq req.ArticleNewsKeyWordsDelReq
	if response.IsFailWithResp(c, util.VerifyUtil.VerifyBody(c, &delReq)) {
		return
	}
	response.CheckAndResp(c, dh.srv.Del(delReq.ID))
}
