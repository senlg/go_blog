package utils

import (
	"crypto/md5"
	"encoding/hex"
	"go_blog/global"
	"time"
)

// 检验元素是否存在数组内
func Find[T int | string | bool](slice []T, val T) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

// 检验日期时间格式是否正确
func VerificationDateList(dateList []string) (ok bool, err error) {
	for _, v := range dateList {
		_, err := time.Parse("2006-01-02 15:04:05", v)
		if err != nil {
			global.Log.Warnln(err.Error())
			return false, err
		}
	}

	return true, nil
}

// Md5
// b 文件流
// salt 加的盐
func Md5(b []byte, salt string) string {
	h := md5.New()

	h.Write(b)
	return string(hex.EncodeToString(h.Sum([]byte(salt))))
}
