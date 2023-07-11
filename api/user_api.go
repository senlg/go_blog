package api

import (
	"fmt"
	"go_blog/common"
	"go_blog/global"
	"go_blog/models"
	"go_blog/models/req"
	"go_blog/models/res"
	jwtauth "go_blog/utils/jwt_auth"
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
	response := common.Response{}
	fmt.Printf("ctx.Request.Header: %+v\n", ctx.ClientIP())
	err := ctx.BindJSON(&loginJson)
	if err != nil {
		response.ResultWithError(ctx, common.ErrorStatus, err)
	}
	global.DB.Preload("ArticleModels").Preload("CollectsModels").Where("user_name = ?", loginJson.UserName).Find(&userModel)
	j := jwtauth.Jwt{}
	claim := &jwtauth.MyCustomClaims{
		UserID:   userModel.ID,
		Role:     userModel.Role,
		Username: userModel.UserName,
	}
	token, err := j.CreateToken(*claim)
	if err != nil {
		response.ResultWithError(ctx, common.ErrorStatus, err)
	}

	userInfo := res.UserInfo{
		UserName:      userModel.UserName,
		NickName:      userModel.NickName,
		AvatarUrl:     userModel.AvatarUrl,
		Addr:          userModel.Addr,
		Token:         token,
		Role:          userModel.Role,
		Phone:         userModel.Phone,
		ReleaseCount:  len(userModel.ArticleModels),
		CollectsCount: len(userModel.CollectsModels),
	}
	global.DB.Create(&models.LoginRecordModel{
		IP:          ctx.ClientIP(),
		UserId:      userModel.ID,
		LoginTime:   time.Now(),
		LoginAdress: loginJson.LoginAdress,
	})
	if global.DB.Error != nil {
		global.Log.Warnln(global.DB.Error)
		response.ResultWithError(ctx, common.ErrorStatus, err)
		return
	}
	response.ResultOk(ctx, userInfo)
}
