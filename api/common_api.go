package api

import (
	"errors"
	"go_blog/common"
	"go_blog/global"
	"go_blog/models"
	"go_blog/models/req"
	jwtauth "go_blog/utils/jwt_auth"

	"github.com/gin-gonic/gin"
)

type TokenApi struct {
}

// TODO
// func (t *TokenApi) GetVisitorToken(ctx *gin.Context) {

//		// j := jwtauth.Jwt{}
//		// j.CreateToken(jwtauth.MyCustomClaims{
//		// 	UserID: -1,
//		// })
//	}
func (t *TokenApi) RefreshToken(ctx *gin.Context) {
	var requestStruct req.RefreshToken
	var response common.Response
	ctx.BindJSON(&requestStruct)
	oldToken := requestStruct.OldToken
	if oldToken == "" {
		response.ResultWithError(ctx, common.RequestError, errors.New("缺少必要参数"))
		return
	}

	j := jwtauth.Jwt{}
	claims, err := j.ParseToken(oldToken)
	if err != nil {
		response.ResultWithError(ctx, common.ErrorStatus, err)
		return
	}
	userId := claims.UserID
	userModel := models.UserModel{}
	err = global.DB.Take(userModel, userId).Error
	if err != nil {
		response.ResultWithError(ctx, common.RequestError, errors.New("token错误,未找到相应用户"))
		return
	}
	token, err := j.CreateToken(jwtauth.MyCustomClaims{
		UserID:   userModel.ID,
		Username: userModel.UserName,
		Role:     userModel.Role,
	})
	if err != nil {
		response.ResultWithError(ctx, common.ErrorStatus, err)
		return
	}
	response.ResultOk(ctx, gin.H{
		"new_token": token,
	})
}
