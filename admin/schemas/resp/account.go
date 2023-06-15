package resp

// GetCodeResp 获取验证码返回
type GetCodeResp struct {
	VerifyId      string `json:"verifyId" structs:"verifyId"`
	Base64Content string `json:"base64Content" structs:"base64Content"`
}

// AccountLoginResp 登录成功
type AccountLoginResp struct {
	Token string `json:"token"`
}

// AccountInfoResp 管理员信息
type AccountInfoResp struct {
	AccountId  uint   `json:"accountId" structs:"accountId"`
	Avatar     string `json:"avatar"`
	Username   string `json:"username"`
	CreateTime string `json:"createTime" structs:"createTime"`
}
