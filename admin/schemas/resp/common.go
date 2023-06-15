package resp

// CommonUploadFileResp 上传图片返回信息
type CommonUploadFileResp struct {
	Name string `json:"name" structs:"name"` // 文件名称
	Path string `json:"path" structs:"path"` // 文件路径
}

// ConfigGet 获取信息
type ConfigGet struct {
	Flag    string `json:"flag" structs:"flag"`       // 配置标识
	Content string `json:"content" structs:"content"` // 配置项内容
}

// NoticeListResp 公告列表
type NoticeListResp struct {
	NoticeId   uint   `json:"noticeId" structs:"noticeId"`     // 公告ID
	Title      string `json:"title" structs:"title"`           // 标题
	Rank       uint   `json:"rank" structs:"rank"`             // 权重
	Status     uint   `json:"status" structs:"status"`         // 状态
	CreateTime string `json:"createTime" structs:"createTime"` // 创建时间
}

// NoticeDetailResp 公告详情
type NoticeDetailResp struct {
	NoticeId uint   `json:"noticeId" structs:"noticeId"` // 公告ID
	Title    string `json:"title" structs:"title"`       // 标题
	Content  string `json:"content" structs:"content"`   // 内容
	Status   uint   `json:"status" structs:"status"`     // 状态
	Rank     uint   `json:"rank" structs:"rank"`         // 权重
}

// TaskListResp 任务列表
type TaskListResp struct {
	TaskId     string `json:"taskId" binding:"required"`       // 任务ID
	Title      string `json:"title" binding:"required"`        // 任务名称
	UserId     string `json:"userId" binding:"required"`       // 发送对象
	Template   string `json:"template" binding:"required"`     // 模板
	Content    string `json:"content" binding:"required"`      // 内容
	Status     uint   `json:"status" structs:"status"`         // 状态
	SendTime   string `json:"sendTime,default=0"`              // 发送时间
	CreateTime string `json:"createTime" structs:"createTime"` // 创建时间
}
