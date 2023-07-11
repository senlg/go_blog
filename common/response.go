package common

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 响应通用返回
type Response struct {
	Code StatusType `json:"code"`
	Data any        `json:"data"`
	Msg  string     `json:"msg"`
}

// 列表通用结构
type ListResponse[T any] struct {
	List  []T   `json:"list"`
	Count int64 `json:"count"`
}
type StatusType int
type StatusMap map[StatusType]string

var MyStatusMap StatusMap

func init() {
	resJson, err := ioutil.ReadFile("./common/resError.json")
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
	err = json.Unmarshal(resJson, &MyStatusMap)
	if err != nil {
		log.Fatalf(err.Error())
		return
	}
}

const (
	SucceedStatus StatusType = 0    // 成功响应
	ErrorStatus   StatusType = 1001 // 系统错误
	ErrorAuth     StatusType = 1002 // 用户信息过期
	UploadError   StatusType = 1003 // 上传失败
	RequestError  StatusType = 1004 // 请求错误
)

// 返回json格式响应
func (r *Response) Result(ctx *gin.Context) {
	if r.Msg == "" {
		r.GetMsg()
	}
	ctx.JSON(http.StatusOK, r)
}

// 返回json格式的成功响应
func (r *Response) ResultOk(ctx *gin.Context, data any) {
	r.Code = SucceedStatus
	r.Data = data
	r.Result(ctx)
}

// 返回错误响应
func (r *Response) ResultWithError(ctx *gin.Context, code StatusType, err error) {
	r.Code = code
	r.Msg = err.Error()
	r.Data = ""
	ctx.JSON(http.StatusOK, r)
}

// 根据Code获取msg
func (r *Response) GetMsg() string {
	r.Msg = MyStatusMap[r.Code]
	return r.Msg
}
