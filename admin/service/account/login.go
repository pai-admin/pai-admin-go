package account

import (
	"bytes"
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"gocli/admin/schemas/req"
	"gocli/admin/schemas/resp"
	"gocli/config"
	"gocli/core"
	"gocli/core/response"
	"gocli/models/account"
	"gocli/plugins"
	"gocli/utils"
	"gorm.io/gorm"
	"image/png"
	"strconv"
	"strings"
)

var LoginService = accountLoginService{}

// systemLoginService 系统登录服务实现类
type accountLoginService struct{}

// GetCode 获取验证码
func (loginSrv accountLoginService) GetCode(c *gin.Context) (codeResp resp.GetCodeResp) {
	// 生成随机字符串
	uuid := utils.ToolsUtil.MakeUuid()
	cp := plugins.NewCaptcha(120, 40, 4)
	cp.SetFontPath("static/fonts") // 指定字体目录
	cp.SetFontName("msyhbd")       // 指定字体名字
	cp.SetMode(1)                  // 1：数学算术运算;2普通字符串
	cp.SetDots(6)
	cp.SetLines(0)
	code, img := cp.OutPut()
	buf := new(bytes.Buffer)
	err := png.Encode(buf, img)
	utils.CheckUtil.CheckErr(err, "生成验证码失败")
	codeResp.Base64Content = "data:image/png;base64," + base64.StdEncoding.EncodeToString(buf.Bytes())
	// 存入缓存
	utils.RedisUtil.Set(config.AdminConfig.TokenKey+uuid, utils.ToolsUtil.MakeMd5(code), 7200)
	codeResp.VerifyId = uuid
	return
}

// Login 登录
func (loginSrv accountLoginService) Login(c *gin.Context, req *req.AccountLoginReq) resp.AccountLoginResp {
	// 验证码验证
	code := utils.RedisUtil.Get(config.AdminConfig.TokenKey + req.VerifyId)
	if code == "" {
		panic(response.Failed.Make("验证码已过期!"))
	}
	if utils.ToolsUtil.MakeMd5(strconv.Itoa(req.VerifyCode)) != code {
		panic(response.Failed.Make("验证码不正确!"))
	}
	// 验证一次删除
	utils.RedisUtil.Del(config.AdminConfig.TokenKey + req.VerifyId)
	// 查询用户是否存在
	accountInfo, err := account.FindByUsername(req.Username)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		panic(response.Failed.Make("账号或密码错误!"))
	} else if err != nil {
		panic(response.Failed)
	}
	if accountInfo.DelFlag == 1 {
		panic(response.Failed.Make("账号或密码错误!"))
	}
	// 验证密码
	md5Pwd := utils.ToolsUtil.MakeMd5(req.Password + accountInfo.Salt)
	if accountInfo.Password != md5Pwd {
		panic(response.Failed.Make("账号或密码错误!"))
	}
	// 账号状态
	if accountInfo.Status != 1 {
		panic(response.Failed.Make("账号已停用!"))
	}
	token := utils.ToolsUtil.MakeToken()
	// token缓存登录数据
	str, err := utils.ToolsUtil.ObjToJson(&accountInfo)
	if err != nil {
		panic(response.Failed.Make("账号或密码错误!"))
	}
	utils.RedisUtil.Set(config.AdminConfig.TokenKey+token, str, config.AdminConfig.TokenTTL)
	// 返回登录信息
	return resp.AccountLoginResp{Token: token}
}

// Logout 退出登录
func (loginSrv accountLoginService) Logout(c *gin.Context) {
	// Token是否为空
	token := c.Request.Header.Get("Authorization")
	utils.RedisUtil.Del(config.AdminConfig.TokenKey + token)
}

// ChangePwd 修改密码
func (loginSrv accountLoginService) ChangePwd(c *gin.Context, req *req.AccountChangePwd) {
	// 检查id
	var admin account.Account
	err := core.DB.Where("account_id = ? AND del_flag = ?", config.AdminConfig.GetAdminId(c), 0).Limit(1).First(&admin).Error
	utils.CheckUtil.CheckErrDBNotRecord(err, "账号不存在或已被删除!")
	utils.CheckUtil.CheckErr(err, "修改失败")
	// 验证原密码是否正确
	currPass := utils.ToolsUtil.MakeMd5(req.OldPassword + admin.Salt)
	if currPass != admin.Password {
		panic(response.Failed.Make("原密码不正确!"))
	}
	// 去掉空格
	req.NewPassword = strings.Trim(req.NewPassword, " ")
	passwdLen := len(req.NewPassword)
	if !(passwdLen >= 6 && passwdLen <= 20) {
		panic(response.Failed.Make("密码必须在6~20位"))
	}
	// 更新密码
	var upAccount account.Account
	salt := utils.ToolsUtil.RandomString(6)
	upAccount.Salt = salt
	upAccount.Password = utils.ToolsUtil.MakeMd5(req.NewPassword + salt)
	upAccount.UpdateTime = utils.ToolsUtil.CurDataTime()
	err = core.DB.Model(&admin).Updates(upAccount).Error
	utils.CheckUtil.CheckErr(err, "修改失败")
	// 修改成功退出登录
	loginSrv.Logout(c)
}

func (loginSrv accountLoginService) ChangeInfo(c *gin.Context, req *req.AccountChangeInfo) {
	// 检查id
	var admin account.Account
	err := core.DB.Where("account_id = ? AND del_flag = ?", config.AdminConfig.GetAdminId(c), 0).Limit(1).First(&admin).Error
	utils.CheckUtil.CheckErrDBNotRecord(err, "账号不存在或已被删除!")
	utils.CheckUtil.CheckErr(err, "修改失败")
	// 更新的数据
	var upAccount account.Account
	if req.Avatar == "" {
		req.Avatar = "assets/images/avatar.png"
	}
	upAccount.Avatar = req.Avatar
	upAccount.UpdateTime = utils.ToolsUtil.CurDataTime()
	err = core.DB.Model(&admin).Updates(upAccount).Error
	utils.CheckUtil.CheckErr(err, "修改失败")
}

func (loginSrv accountLoginService) AccountInfo(c *gin.Context) (accountResp resp.AccountInfoResp) {
	var accountInfo account.Account
	err := core.DB.Where("account_id = ?", config.AdminConfig.GetAdminId(c)).
		Where("del_flag = 0").
		Select("account_id, username, create_time, avatar").
		Limit(1).
		First(&accountInfo).Error
	utils.CheckUtil.CheckErr(err, "获取失败")
	utils.CheckUtil.CheckErrDBNotRecord(err, "用户不存在")
	response.Copy(&accountResp, accountInfo)
	return
}
