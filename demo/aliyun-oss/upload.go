package aliyun_oss

import (
	"errors"
	"mime/multipart"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
)

// 正式环境建议用 viper 读取 yaml 替换
var (
	AliyunOSSBasePath        = "xxxxx"
	AliyunOSSBucketUrl       = "xxxxx"
	AliyunOSSEndpoint        = "xxxxx"
	AliyunOSSAccessKeyId     = ""
	AliyunOSSAccessKeySecret = ""
	AliyunOSSBucketName      = ""
)

type AliyunOSS struct{}

func (*AliyunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	bucket, err := NewBucket()
	if err != nil {
		zap.L().Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	// 读取本地文件。
	f, openError := file.Open()
	if openError != nil {
		zap.L().Error("function file.Open() Failed", zap.Any("err", openError.Error()))
		return "", "", errors.New("function file.Open() Failed, err:" + openError.Error())
	}
	defer f.Close() // 创建文件 defer 关闭
	// 上传阿里云路径 文件名格式 自己可以改 建议保证唯一性
	// yunFileTmpPath := filepath.Join("uploads", time.Now().Format("2006-01-02")) + "/" + file.Filename
	yunFileTmpPath := AliyunOSSBasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + file.Filename

	// 上传文件流。
	err = bucket.PutObject(yunFileTmpPath, f)
	if err != nil {
		zap.L().Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return AliyunOSSBucketUrl + "/" + yunFileTmpPath, yunFileTmpPath, nil
}

func NewBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(AliyunOSSEndpoint, AliyunOSSAccessKeyId, AliyunOSSAccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(AliyunOSSBucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
