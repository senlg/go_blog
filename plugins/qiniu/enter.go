package qiniu

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"go_blog/config"
	"go_blog/global"
	"time"

	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
)

// 获取七牛云上传token
func getToken(c config.Qiniu) string {
	accessKey := c.AccessKey
	secretKey := c.SecretKey
	bucket := c.Bucket
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}

func getConfig(q config.Qiniu) storage.Config {
	cfg := storage.Config{}

	zone, _ := storage.GetRegionByID(storage.RegionID(q.Zone))
	cfg.Zone = &zone
	// 是否使用https 域名
	cfg.UseHTTPS = false
	// 是否使用cdn 加速上传
	cfg.UseCdnDomains = false
	return cfg
}
func UploadQiNiu(data []byte, fileName string, prefix string) (filePath string, err error) {
	q := global.Config.Qiniu
	if q.AccessKey == "" || q.SecretKey == "" {
		return "", errors.New("请配置 AccessKey 或 SecretKey")
	}

	if float64(len(data))/1024/1024 > float64(q.Size) {
		return "", errors.New("上传文件大于设定大小")
	}
	upToken := getToken(q)
	cfg := getConfig(q)

	formUploader := storage.NewFormUploader(&cfg)

	ret := storage.PutRet{}
	putExtra := storage.PutExtra{
		Params: map[string]string{},
	}
	dataLen := int64(len(data))
	// 获取当前时间
	now := time.Now().Format("20060102150405")
	key := fmt.Sprintf("%s/%s__%s", prefix, now, fileName)

	err = formUploader.Put(context.Background(), &ret, upToken, key, bytes.NewReader(data), dataLen, &putExtra)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/%s", q.Cdn, ret.Key), nil
}
