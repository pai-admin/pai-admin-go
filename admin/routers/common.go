package routers

import (
	"github.com/gin-gonic/gin"
	"gocli/admin/schemas/req"
	"gocli/admin/service/common"
	"gocli/core"
	"gocli/core/request"
	"gocli/core/response"
	"gocli/utils"
)

var CommonGroup = core.Group("/")

func init() {
	group := CommonGroup
	group.AddPOST("/upload", upload)
	group.AddGET("/config/get", configGet)
	group.AddPOST("/config/set", configSet)
	group.AddGET("/configs/get", configsGet)
	group.AddPOST("/configs/set", configsSet)
	group.AddGET("/notice/list", noticeList)
	group.AddGET("/notice/detail", noticeDetail)
	group.AddDELETE("/notice/del", noticeDel)
	group.AddPOST("/notice/add", noticeAdd)
	group.AddPUT("/notice/edit", noticeEdit)
	group.AddGET("/task/list", taskList)
	group.AddPOST("/task/add", taskAdd)
	group.AddPUT("/task/edit", taskEdit)
	group.AddDELETE("/task/del", taskDel)
}

// uploadImage 上传图片
func upload(c *gin.Context) {
	file := utils.VerifyUtil.VerifyFile(c, "file")
	response.OkWithData(c, common.UploadService.Upload(file))
}

func configSet(c *gin.Context) {
	var setReq req.ConfigSetReq
	utils.VerifyUtil.VerifyJSON(c, &setReq)
	common.ConfigService.ConfigSet(setReq)
	response.Ok(c)
}

func configGet(c *gin.Context) {
	var getReq req.ConfigGetReq
	utils.VerifyUtil.VerifyQuery(c, &getReq)
	response.OkWithData(c, common.ConfigService.ConfigGet(getReq))
}

func configsSet(c *gin.Context) {
	var setReq map[string]string
	utils.VerifyUtil.VerifyJSON(c, &setReq)
	common.ConfigService.ConfigsSet(setReq)
	response.Ok(c)
}

func configsGet(c *gin.Context) {
	var getReq req.ConfigGetReq
	utils.VerifyUtil.VerifyQuery(c, &getReq)
	response.OkWithData(c, common.ConfigService.ConfigsGet(getReq))
}

// noticeList 文章列表
func noticeList(c *gin.Context) {
	var page request.PageReq
	var listReq req.NoticeListReq
	utils.VerifyUtil.VerifyQuery(c, &page)
	utils.VerifyUtil.VerifyQuery(c, &listReq)
	response.OkWithData(c, common.ConfigService.NoticeList(page, listReq))
}

// noticeList 文章列表
func noticeDetail(c *gin.Context) {
	var detailReq req.NoticeDetailReq
	utils.VerifyUtil.VerifyQuery(c, &detailReq)
	response.OkWithData(c, common.ConfigService.NoticeDetail(detailReq))
}

// noticeDel 删除文章
func noticeDel(c *gin.Context) {
	var delReq req.NoticeDelReq
	utils.VerifyUtil.VerifyQuery(c, &delReq)
	common.ConfigService.NoticeDel(delReq)
	response.Ok(c)
}

func noticeAdd(c *gin.Context) {
	var addReq req.NoticeAddReq
	utils.VerifyUtil.VerifyJSON(c, &addReq)
	common.ConfigService.NoticeAdd(addReq)
	response.Ok(c)
}

func noticeEdit(c *gin.Context) {
	var editReq req.NoticeEditReq
	utils.VerifyUtil.VerifyJSON(c, &editReq)
	common.ConfigService.NoticeEdit(editReq)
	response.Ok(c)
}

func taskList(c *gin.Context) {
	var page request.PageReq
	var listReq req.TaskListReq
	utils.VerifyUtil.VerifyQuery(c, &page)
	utils.VerifyUtil.VerifyQuery(c, &listReq)
	response.OkWithData(c, common.ConfigService.TaskList(page, listReq))
}

func taskDel(c *gin.Context) {
	var delReq req.TaskDelReq
	utils.VerifyUtil.VerifyQuery(c, &delReq)
	common.ConfigService.TaskDel(delReq)
	response.Ok(c)
}

func taskAdd(c *gin.Context) {
	var addReq req.TaskAddReq
	utils.VerifyUtil.VerifyJSON(c, &addReq)
	common.ConfigService.TaskAdd(addReq)
	response.Ok(c)
}

func taskEdit(c *gin.Context) {
	var editReq req.TaskEditReq
	utils.VerifyUtil.VerifyJSON(c, &editReq)
	common.ConfigService.TaskEdit(editReq)
	response.Ok(c)
}
