package req

// ConfigSetReq 修改配置
type ConfigSetReq struct {
	Flag    string `json:"flag" binding:"required"` // 配置标识
	Content string `json:"content"`                 // 配置项内容
}

// ConfigGetReq 获取配置
type ConfigGetReq struct {
	Flag string `form:"flag" binding:"required"` // 配置标识
}

// NoticeListReq 公告列表
type NoticeListReq struct {
	Title string `form:"title"` // 标题
}

// NoticeDetailReq 公告详情
type NoticeDetailReq struct {
	NoticeId uint `form:"noticeId" binding:"required"` // 文章ID
}

// NoticeDelReq 公告删除
type NoticeDelReq struct {
	NoticeId string `form:"noticeId" binding:"required"` // 文章IDs
}

// NoticeAddReq 添加公告
type NoticeAddReq struct {
	Title   string `post:"title" binding:"required"`   // 公告名称
	Content string `post:"content" binding:"required"` // 公告内容
	Rank    uint   `post:"rank,default=0"`             // 排序权重
	Status  uint8  `post:"status,default=1"`           // 是否显示
}

// NoticeEditReq 修改公告
type NoticeEditReq struct {
	NoticeId uint   `post:"noticeId" binding:"required"` // 公告ID
	Title    string `post:"title" binding:"required"`    // 公告名称
	Content  string `post:"content" binding:"required"`  // 公告内容
	Rank     uint   `post:"rank,default=0"`              // 排序权重
	Status   uint8  `post:"status,default=1"`            // 是否显示
}

// TaskListReq 列表
type TaskListReq struct {
	Title string `form:"title"` // 标题
}

// TaskDelReq 删除
type TaskDelReq struct {
	TaskId string `form:"taskId" binding:"required"` // 任务IDs
}

// TaskAddReq 添加
type TaskAddReq struct {
	Title    string `post:"title" binding:"required"`    // 任务名称
	UserId   string `post:"userId" binding:"required"`   // 发送对象
	Template string `post:"template" binding:"required"` // 模板
	Content  string `post:"content" binding:"required"`  // 内容
	SendTime string `post:"sendTime,default=0"`          // 发送时间
}

// TaskEditReq 修改
type TaskEditReq struct {
	TaskId   uint   `post:"taskId" binding:"required"`   // 任务ID
	UserId   string `post:"userId" binding:"required"`   // 发送对象
	Title    string `post:"title" binding:"required"`    // 任务名称
	Template string `post:"template" binding:"required"` // 模板
	Content  string `post:"content" binding:"required"`  // 内容
	SendTime string `post:"sendTime,default=0"`          // 发送时间
}
