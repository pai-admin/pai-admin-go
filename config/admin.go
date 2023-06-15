package config

import "github.com/gin-gonic/gin"

// AdminConfig 后台公共配置
var AdminConfig = adminConfig{
	// 令牌缓存键
	TokenKey: "admin:token:",
	TokenTTL: 3600,

	ReqAdminIdKey:  "admin_id",
	ReqUsernameKey: "username",
	// 无需登录接口配置
	NotNeedLogin: []string{
		"login",    // 登录接口
		"get-code", // 获取验证码
	},
}

type adminConfig struct {
	NotNeedLogin   []string
	TokenKey       string
	ReqAdminIdKey  string
	ReqUsernameKey string
	TokenTTL       int
}

func (cnf adminConfig) GetAdminId(c *gin.Context) uint {
	adminId, _ := c.Get(cnf.ReqAdminIdKey)
	return adminId.(uint)
}

func (cnf adminConfig) GetUsername(c *gin.Context) string {
	username, _ := c.Get(cnf.ReqUsernameKey)
	return username.(string)
}
