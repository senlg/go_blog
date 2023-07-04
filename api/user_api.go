package api

import (
	"fmt"
	"go_blog/common"
	"go_blog/global"
	"go_blog/models"
	"go_blog/models/req"
	"go_blog/models/res"
	"time"

	"github.com/gin-gonic/gin"
)

type User struct{}

// 获取用户信息
func (u *User) GetUserInfo(ctx *gin.Context) {
	ctx.Next()
	var user models.UserModel
	// var userInfo res.UserResModel
	global.DB.Association("ArticleModels").Find(&user, "id = ?", 1)
	fmt.Printf("\n%+v\n", &user)
	res := common.Response{
		Code: common.SucceedStatus,
		Data: &user,
	}
	if global.DB.Error != nil {
		global.Log.Warnln(global.DB.Error)
	}
	res.Result(ctx)
}

// 创建用户
func (u *User) CreateUserInfo(ctx *gin.Context) {
	var userReg req.RegisterUser
	ctx.BindJSON(&userReg)
	fmt.Printf("ctx.Request: %+v\n", &userReg)

	user := models.UserModel{
		NickName:  userReg.NickName,
		UserName:  userReg.UserName,
		Password:  userReg.Password,
		Phone:     userReg.Phone,
		AvatarUrl: userReg.AvatarUrl,
		Addr:      userReg.Addr,
		Email:     userReg.Email,
		Role:      2,
	}
	global.DB.Create(&user)
	res := common.Response{
		Code: common.SucceedStatus,
		Data: &user,
	}
	if global.DB.Error != nil {
		global.Log.Warnln(global.DB.Error)
	}
	res.Result(ctx)
}

// 登录
func (u *User) Login(ctx *gin.Context) {
	var loginJson req.Login
	var userModel models.UserModel
	fmt.Printf("ctx.Request.Header: %+v\n", ctx.ClientIP())
	err := ctx.BindJSON(&loginJson)

	global.DB.Preload("ArticleModels").Preload("CollectsModels").Where("user_name = ?", loginJson.UserName).Find(&userModel)

	userInfo := res.UserInfo{
		UserName:  userModel.UserName,
		NickName:  userModel.NickName,
		AvatarUrl: userModel.AvatarUrl,
		Addr:      userModel.Addr,
		// Token:         userModel.Token,
		Role:          userModel.Role,
		Phone:         userModel.Phone,
		ReleaseCount:  len(userModel.ArticleModels),
		CollectsCount: len(userModel.CollectsModels),
	}
	res := common.Response{
		Code: common.SucceedStatus,
		Data: &userInfo,
	}
	if err != nil {
		global.Log.Errorln(err.Error())
		return
	}
	if global.DB.Error != nil {
		global.Log.Warnln(global.DB.Error)
		return
	}
	global.DB.Create(&models.LoginRecordModel{
		IP:          ctx.ClientIP(),
		UserId:      userModel.ID,
		LoginTime:   time.Now(),
		LoginAdress: loginJson.LoginAdress,
	})
	res.Result(ctx)
}
