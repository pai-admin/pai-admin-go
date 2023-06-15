package plugins

import (
	"gocli/config"
	"gocli/core"
	"gocli/core/response"
	"gocli/models/common"
	"gocli/utils"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"
)

var StorageDriver = storageDriver{}

// UploadFile 文件对象
type UploadFile struct {
	Name string // 文件名称
	Path string // 文件路径
}

// storageDriver 存储引擎
type storageDriver struct{}

// Upload 根据引擎类型上传文件
func (sd storageDriver) Upload(file *multipart.FileHeader) *UploadFile {
	var configs []common.Config
	// 获取微信小程序配置信息
	err := core.DB.Table(core.DBTableName(&common.Config{})).
		Where("del_flag = 0 AND flag in ('uploadtype','domain')").
		Select("flag, content").Find(&configs).Error
	utils.CheckUtil.CheckErrDBNotRecord(err, "读取微信配置失败!")
	utils.CheckUtil.CheckErr(err, "读取微信配置失败")
	var conf = make(map[string]string)
	for i := 0; i < len(configs); i++ {
		conf[configs[i].Flag] = configs[i].Content
	}
	fileName := sd.buildSaveName(file)
	date := time.Now().Format("20060201")
	sd.localUpload(file, path.Join(date, fileName))

	var filePath = path.Join(config.Config.StaticPath, "uploads", date, fileName)

	return &UploadFile{
		Name: fileName,
		Path: conf["domain"] + filePath,
	}
}

// localUpload 本地上传
func (sd storageDriver) localUpload(file *multipart.FileHeader, key string) {
	// 映射目录
	directory := config.Config.UploadDirectory
	// 打开源文件
	src, err := file.Open()
	if err != nil {
		core.Logger.Errorf("storageDriver.localUpload Open err: err=[%+v]", err)
		panic(response.Failed.Make("打开文件失败!"))
	}
	defer src.Close()
	// 文件信息
	savePath := path.Join(directory, path.Dir(key))
	saveFilePath := path.Join(directory, key)
	// 创建目录
	err = os.MkdirAll(savePath, 0755)
	if err != nil && !os.IsExist(err) {
		core.Logger.Errorf(
			"storageDriver.localUpload MkdirAll err: path=[%s], err=[%+v]", savePath, err)
		panic(response.Failed.Make("创建上传目录失败!"))
	}
	// 创建目标文件
	out, err := os.Create(saveFilePath)
	if err != nil {
		core.Logger.Errorf(
			"storageDriver.localUpload Create err: file=[%s], err=[%+v]", saveFilePath, err)
		panic(response.Failed.Make("创建文件失败!"))
	}
	defer out.Close()
	// 写入目标文件
	_, err = io.Copy(out, src)
	if err != nil {
		core.Logger.Errorf(
			"storageDriver.localUpload Copy err: file=[%s], err=[%+v]", saveFilePath, err)
		panic(response.Failed.Make("上传文件失败: " + err.Error()))
	}
}

// checkFile 生成文件名称
func (sd storageDriver) buildSaveName(file *multipart.FileHeader) string {
	name := file.Filename
	ext := strings.ToLower(path.Ext(name))
	return utils.ToolsUtil.MakeUuid() + ext
}
