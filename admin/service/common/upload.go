package common

import (
	"gocli/admin/schemas/resp"
	"gocli/plugins"
	"mime/multipart"
)

var UploadService = uploadService{}

// uploadService 上传服务实现类
type uploadService struct{}

// Upload 上传文件
func (upSrv uploadService) Upload(file *multipart.FileHeader) (res resp.CommonUploadFileResp) {
	upRes := plugins.StorageDriver.Upload(file)
	res.Name = upRes.Name
	res.Path = upRes.Path
	return res
}
