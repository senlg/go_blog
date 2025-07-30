package api

import (
	"errors"
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
// GetUserInfo
// @Summary 获取用户信息
// @Produce json
// @Description 获取用户信息
// @Tags users
// @Accept json
// @Param user_id body req.UserInfo true "用户ID"
// @Success 200 {object} res.UserInfo "成功"
// @Router /api/get_user [Post]
func (u *User) GetUserInfo(ctx *gin.Context) {
	var requestStruct req.UserInfo
	var response common.Response
	ctx.BindJSON(&requestStruct)
	if requestStruct.UserId == 0 {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	var user models.UserModel
	// var userInfo res.UserResModel
	global.DB.Preload("ArticleModels").Preload("CollectsModels").Find(&user, "id = ?", requestStruct.UserId)
	if global.DB.Error != nil {
		global.Log.Warnln(global.DB.Error)
		response.ResultWithError(ctx, common.RequestError, global.DB.Error)
		return
	}
	var role = models.Role(user.Role)
	var userInfo = res.UserInfo{
		Id:            user.ID,
		UserName:      user.UserName,
		NickName:      user.NickName,
		AvatarUrl:     user.NickName,
		Addr:          user.Addr,
		Role:          user.Role,
		RoleName:      role.ToString(),
		Phone:         user.Phone,
		ReleaseCount:  len(user.ArticleModels),
		CollectsCount: len(user.CollectsModels),
	}
	response.ResultOk(ctx, userInfo)
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
		Role:      models.PermissionUser,
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

// 查询用户列表
func (u *User) GetUserInfoList(ctx *gin.Context) {
	var requestStruct req.UserInfoRequest
	var response common.Response
	ctx.BindJSON(&requestStruct)
	tx := global.DB.Model(&models.UserModel{})
	if requestStruct.Limit < 1 {
		requestStruct.Limit = 5

	}
	if requestStruct.Page < 1 {
		requestStruct.Page = 1
	}
	if requestStruct.UserName != "" {
		tx = tx.Where("user_name = ?", requestStruct.UserName)
	}
	var userList []res.UserItem
	var count int64
	offset := requestStruct.Limit * (requestStruct.Page - 1)
	err := tx.Preload("ReleaseCount").Limit(requestStruct.Limit).Offset(offset).Scan(&userList).Count(&count).Error
	if err != nil {
		response.ResultWithError(ctx, common.ErrorStatus, err)
		return
	}
	response.ResultOk(ctx, common.ListResponse[res.UserItem]{
		List:  userList,
		Count: count,
	})
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
	if loginJson.UserName == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}
	if loginJson.Password == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}

	Affected := global.DB.Preload("ArticleModels").Preload("CollectsModels").Where("user_name = ? and password = ?", loginJson.UserName, loginJson.Password).Find(&userModel).RowsAffected
	if Affected == 0 {
		global.Log.Infoln(global.DB.Error)
		response.ResultWithError(ctx, common.RequestError, errors.New("用户名或密码错误"))
		return
	}
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

// 删除用户
func (u *User) DelUser(ctx *gin.Context) {

	var requestStruct req.DeleteUserInfo
	var response common.Response
	c, exit := ctx.Get("claim")
	var role models.Role

	if exit {
		v, ok := c.(jwtauth.MyCustomClaims)
		if ok {
			role = v.Role
		}
	}
	if role != models.PermissionAdmin {
		response.ResultWithError(ctx, common.RequestError, errors.New("无权限"))
	}
	ctx.BindJSON(requestStruct)
	if len(requestStruct.Ids) < 1 {
		response.ResultWithError(ctx, common.RequestError, errors.New("ids不能为空"))
		return
	}

	count := global.DB.Delete(&models.UserModel{}, requestStruct.Ids).RowsAffected
	response.ResultOk(ctx, fmt.Sprintf("已删除%d个用户", count))

}
