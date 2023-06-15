package middleware

import (
	"github.com/gin-gonic/gin"
	"gocli/config"
	"gocli/core"
	"gocli/core/response"
	"gocli/models/account"
	"gocli/utils"
	"strings"
)

// TokenAuth Token认证中间件
func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 根据接口前缀判断模块
		if strings.HasPrefix(c.Request.URL.Path, "/admin/") {
			// 管理端
			api := strings.ReplaceAll(strings.Replace(c.Request.URL.Path, "/admin/", "", 1), "/", ":")

			// 免登录接口
			if utils.ToolsUtil.Contains(config.AdminConfig.NotNeedLogin, api) {
				c.Next()
				return
			}

			// Token是否为空
			token := c.Request.Header.Get("Authorization")
			if token == "" {
				response.Failed.Make("登录已过期!")
				c.Abort()
				return
			}

			// Token是否过期
			token = config.AdminConfig.TokenKey + token
			existCnt := utils.RedisUtil.Exists(token)
			if existCnt < 0 {
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			} else if existCnt == 0 {
				response.Failed.Make("登录已过期!")
				c.Abort()
				return
			}

			// 用户信息缓存
			accountStr := utils.RedisUtil.Get(token)
			var mapping account.Account
			err := utils.ToolsUtil.JsonToObj(accountStr, &mapping)
			if err != nil {
				core.Logger.Errorf("读取缓存出错了: [%+v]", err)
				response.Fail(c, response.SystemError)
				c.Abort()
				return
			}

			// token自动续签
			if utils.RedisUtil.TTL(token) < 1800 {
				utils.RedisUtil.Expire(token, 7200)
			}

			c.Set(config.AdminConfig.ReqAdminIdKey, mapping.AccountId)
			c.Set(config.AdminConfig.ReqUsernameKey, mapping.Username)

			c.Next()
			return
		}
		core.Logger.Errorf("被禁止的页面")
		response.FailWithMsg(c, response.SystemError, "被禁止的页面")
		c.Abort()
		return
	}
}
