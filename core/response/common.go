package response

// PageResp 分页返回值
type PageResp struct {
	Count int64       `json:"count"` // 总数
	Lists interface{} `json:"lists"` // 数据
}
