package routers

import (
	"github.com/gin-gonic/gin"
	"gocli/admin/schemas/req"
	"gocli/admin/service/account"
	"gocli/core"
	"gocli/core/response"
	"gocli/utils"
)

var SystemGroup = core.Group("/")

func init() {
	group := SystemGroup
	group.AddGET("/get-code", getCode)
	group.AddPOST("/login", login)

	group.AddPOST("/logout", logout)
	group.AddPOST("/changePwd", changePwd)
	group.AddPOST("/changeInfo", changeInfo)
	group.AddGET("/info", accountInfo)
}

// 获取验证码
func getCode(c *gin.Context) {
	resp := account.LoginService.GetCode(c)
	response.OkWithData(c, resp)
}

// login 登录系统
func login(c *gin.Context) {
	var loginReq req.AccountLoginReq
	utils.VerifyUtil.VerifyJSON(c, &loginReq)
	resp := account.LoginService.Login(c, &loginReq)
	response.OkWithData(c, resp)
}

// logout 退出登录
func logout(c *gin.Context) {
	account.LoginService.Logout(c)
	response.OkWithMsg(c, "登出成功")
}

// changePwd 修改密码
func changePwd(c *gin.Context) {
	var changePwdReq req.AccountChangePwd
	utils.VerifyUtil.VerifyJSON(c, &changePwdReq)
	account.LoginService.ChangePwd(c, &changePwdReq)
	response.OkWithMsg(c, "修改成功")
}

// changeInfo 修改管理员信息
func changeInfo(c *gin.Context) {
	var changePwdInfo req.AccountChangeInfo
	utils.VerifyUtil.VerifyJSON(c, &changePwdInfo)
	account.LoginService.ChangeInfo(c, &changePwdInfo)
	response.OkWithMsg(c, "修改成功")
}

func accountInfo(c *gin.Context) {
	response.OkWithData(c, account.LoginService.AccountInfo(c))
}
