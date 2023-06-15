package req

// AccountLoginReq 账户登录
type AccountLoginReq struct {
	VerifyId   string `json:"verifyId" binding:"required"`              // 验证码ID
	VerifyCode int    `json:"verifyCode" binding:"required"`            // 验证码
	Username   string `json:"username" binding:"required,min=2,max=20"` // 账号
	Password   string `json:"password" binding:"required,min=6,max=32"` // 密码
}

// AccountChangePwd 修改密码
type AccountChangePwd struct {
	OldPassword string `json:"old_password" binding:"required,min=6,max=32"` // 原密码
	NewPassword string `json:"new_password" binding:"required,min=6,max=32"` // 新密码
}

// AccountChangeInfo 修改管理员信息
type AccountChangeInfo struct {
	Avatar string `json:"avatar"` // 头像
}
