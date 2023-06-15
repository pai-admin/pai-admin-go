package common

import (
	"fmt"
	"github.com/fatih/structs"
	"gocli/admin/schemas/req"
	"gocli/admin/schemas/resp"
	"gocli/core"
	"gocli/core/request"
	"gocli/core/response"
	"gocli/models/common"
	"gocli/utils"
	"strings"
)

var ConfigService = configService{}

// configService 配置项
type configService struct{}

// ConfigGet 获取
func (configSrv configService) ConfigGet(getReq req.ConfigGetReq) (configResp resp.ConfigGet) {
	var configGet common.Config
	err := core.DB.Where("flag = ? AND del_flag = 0", getReq.Flag).Limit(1).First(&configGet).Error
	utils.CheckUtil.CheckErrDBNotRecord(err, "配置项不存在或已被删除!")
	response.Copy(&configResp, configGet)
	return
}

// ConfigSet 修改
func (configSrv configService) ConfigSet(setReq req.ConfigSetReq) {
	var configEdit common.Config
	err := core.DB.Where("flag = ? AND del_flag = 0", setReq.Flag).Limit(1).First(&configEdit).Error
	utils.CheckUtil.CheckErrDBNotRecord(err, "配置项不存在或已被删除!")
	err = core.DB.Model(&configEdit).Updates(common.Config{Content: setReq.Content, UpdateTime: utils.ToolsUtil.CurDataTime()}).Error
	utils.CheckUtil.CheckErr(err, "修改失败")
}

func (configSrv configService) NoticeList(page request.PageReq, listReq req.NoticeListReq) response.PageResp {
	// 分页信息
	limit := page.Limit
	offset := page.Limit * (page.Page - 1)
	// 查询
	nModel := core.DB.Table(core.DBTableName(&common.Notice{})).Where("del_flag = 0").Select("notice_id,title,rank,status,create_time")
	// 条件
	if listReq.Title != "" {
		nModel = nModel.Where("title like ?", "%"+listReq.Title+"%")
	}
	// 总数
	var count int64
	err := nModel.Count(&count).Error
	utils.CheckUtil.CheckErr(err, "List Count err")
	// 数据
	var listResp []resp.NoticeListResp
	err = nModel.Limit(limit).Offset(offset).Order("rank desc, notice_id desc").Find(&listResp).Error
	utils.CheckUtil.CheckErr(err, "List Find err")
	return response.PageResp{
		Count: count,
		Lists: listResp,
	}
}

func (configSrv configService) NoticeDetail(detailReq req.NoticeDetailReq) (noticeDetailResp resp.NoticeDetailResp) {
	var noticeDetail common.Notice
	err := core.DB.Where("del_flag = 0").
		Where("notice_id = ?", detailReq.NoticeId).
		Select("notice_id,title,content,rank,status").
		Limit(1).
		First(&noticeDetail).Error
	utils.CheckUtil.CheckErrDBNotRecord(err, "公告不存在")
	response.Copy(&noticeDetailResp, noticeDetail)
	return
}

func (configSrv configService) NoticeDel(delReq req.NoticeDelReq) {
	err := core.DB.Table(core.DBTableName(&common.Notice{})).
		Where("notice_id in (" + delReq.NoticeId + ") AND del_flag = 0").
		Updates(common.Notice{DelFlag: 1, UpdateTime: utils.ToolsUtil.CurDataTime()}).Error
	utils.CheckUtil.CheckErr(err, "删除失败")
}

func (configSrv configService) NoticeAdd(addReq req.NoticeAddReq) {
	var noticeAdd common.Notice
	response.Copy(&noticeAdd, addReq)
	noticeAdd.CreateTime = utils.ToolsUtil.CurDataTime()
	noticeAdd.DelFlag = 0
	noticeAdd.UpdateTime = utils.ToolsUtil.CurDataTime()
	err := core.DB.Create(&noticeAdd).Error
	utils.CheckUtil.CheckErr(err, "添加失败")
}

func (configSrv configService) NoticeEdit(editReq req.NoticeEditReq) {
	var noticeEdit common.Notice
	err := core.DB.Where("notice_id = ? AND del_flag = 0", editReq.NoticeId).Limit(1).First(&noticeEdit).Error
	utils.CheckUtil.CheckErrDBNotRecord(err, "文章不存在或已被删除!")
	utils.CheckUtil.CheckErr(err, "修改失败")
	editMap := structs.Map(editReq)
	delete(editMap, "NoticeId")
	editMap["update_time"] = utils.ToolsUtil.CurDataTime()
	err = core.DB.Model(&noticeEdit).Updates(editMap).Error
	utils.CheckUtil.CheckErr(err, "删除失败")
}

func (configSrv configService) TaskList(page request.PageReq, listReq req.TaskListReq) response.PageResp {
	// 分页信息
	limit := page.Limit
	offset := page.Limit * (page.Page - 1)
	// 查询
	tModel := core.DB.Table(core.DBTableName(&common.Task{})).
		Where("del_flag = 0").
		Select("task_id,title,user_id,status,create_time,template,content,send_time")
	// 条件
	if listReq.Title != "" {
		tModel = tModel.Where("title like ?", "%"+listReq.Title+"%")
	}
	// 总数
	var count int64
	err := tModel.Count(&count).Error
	utils.CheckUtil.CheckErr(err, "List Count err")
	// 数据
	var listResp []resp.TaskListResp
	err = tModel.Limit(limit).Offset(offset).Order("task_id desc").Find(&listResp).Error
	utils.CheckUtil.CheckErr(err, "List Find err")
	return response.PageResp{
		Count: count,
		Lists: listResp,
	}
}

func (configSrv configService) TaskDel(delReq req.TaskDelReq) {
	err := core.DB.Table(core.DBTableName(&common.Task{})).
		Where("task_id in (" + delReq.TaskId + ") AND del_flag = 0").
		Updates(common.Task{DelFlag: 1, UpdateTime: utils.ToolsUtil.CurDataTime()}).Error
	utils.CheckUtil.CheckErr(err, "删除失败")
}

func (configSrv configService) TaskAdd(addReq req.TaskAddReq) {
	var taskAdd common.Task
	response.Copy(&taskAdd, addReq)
	taskAdd.CreateTime = utils.ToolsUtil.CurDataTime()
	taskAdd.DelFlag = 0
	err := core.DB.Create(&taskAdd).Error
	utils.CheckUtil.CheckErr(err, "添加失败")
}

func (configSrv configService) TaskEdit(editReq req.TaskEditReq) {
	var taskEdit common.Task
	err := core.DB.Where("task_id = ? AND del_flag = 0", editReq.TaskId).Limit(1).First(&taskEdit).Error
	utils.CheckUtil.CheckErrDBNotRecord(err, "任务不存在或已被删除!")
	utils.CheckUtil.CheckErr(err, "修改失败")

	if taskEdit.Status == 1 || taskEdit.Status == 2 {
		panic(response.Failed.Make("已执行的任务不能修改"))
	}

	editMap := structs.Map(editReq)
	delete(editMap, "TaskId")
	editMap["update_time"] = utils.ToolsUtil.CurDataTime()
	err = core.DB.Model(&taskEdit).Updates(editMap).Error
	utils.CheckUtil.CheckErr(err, "删除失败")
}

func (configSrv configService) UnTask() (tasks []common.Task) {
	tModel := core.DB.Table(core.DBTableName(&common.Task{})).
		Where("del_flag = 0").
		Where("status = 0").
		Where("send_time <= ? OR send_time = '1970-01-01 08:00:00'", utils.ToolsUtil.CurDataTime()).
		Select("task_id,title,user_id,status,template,content,send_time")
	// 数据
	err := tModel.Find(&tasks).Error
	if err != nil || len(tasks) == 0 {
		return nil
	}

	// 将所有任务状态更改为已执行
	var taskIds []uint
	for i := 0; i < len(tasks); i++ {
		taskIds = append(taskIds, tasks[i].TaskId)
	}
	taskIdStr := strings.Replace(strings.Trim(fmt.Sprint(taskIds), "[]"), " ", ",", -1)
	err = core.DB.Table(core.DBTableName(&common.Task{})).
		Where("task_id in (" + taskIdStr + ") AND del_flag = 0").
		Updates(common.Task{Status: 1, UpdateTime: utils.ToolsUtil.CurDataTime()}).Error
	if err != nil {
		return nil
	}
	return
}

// ConfigsGet 获取
func (configSrv configService) ConfigsGet(getReq req.ConfigGetReq) (conf map[string]string) {
	var configs []common.Config
	err := core.DB.Table(core.DBTableName(&common.Config{})).
		Where("flag in (?) AND del_flag = 0", strings.Split(getReq.Flag, ",")).
		Select("flag,content").
		Find(&configs).Error
	utils.CheckUtil.CheckErrDBNotRecord(err, "配置项不存在或已被删除!")
	conf = make(map[string]string)
	for i := 0; i < len(configs); i++ {
		conf[configs[i].Flag] = configs[i].Content
	}
	return
}

// ConfigsSet 修改
func (configSrv configService) ConfigsSet(setsReq map[string]string) {
	for k, v := range setsReq {
		err := core.DB.Table(core.DBTableName(&common.Config{})).
			Where("del_flag = 0").
			Where("flag = ?", k).
			Updates(common.Config{Content: v, UpdateTime: utils.ToolsUtil.CurDataTime()}).
			Error
		utils.CheckUtil.CheckErr(err, "修改失败")
	}
}
