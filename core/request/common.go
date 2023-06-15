package request

// PageReq 分页请求参数
type PageReq struct {
	Page  int `form:"page,default=1" validate:"omitempty,gte=1"`         // 页码
	Limit int `form:"limit,default=10" validate:"omitempty,gt=0,lte=60"` // 每页大小
}
