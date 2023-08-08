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

	"github.com/gin-gonic/gin"
)

type MenuApi struct {
}

func (m MenuApi) GetMenuList(ctx *gin.Context) {
	var responseStruct common.Response
	c, exit := ctx.Get("claim")
	if exit {
		global.Log.Infof("%+v\n", c)
	}
	var calim *jwtauth.MyCustomClaims = c.(*jwtauth.MyCustomClaims)
	if calim.Role != models.PermissionVisitor {
		var menu_list []res.MenuItem
		err := global.DB.Model(&models.MenuModel{}).Scan(&menu_list).Error
		if err != nil {
			responseStruct.ResultWithError(ctx, common.ErrorStatus, err)
			return
		}
		responseStruct.ResultOk(ctx, &menu_list)
		return
	}
	responseStruct.ResultWithError(ctx, common.RequestError, errors.New("暂无权限"))
}

func (m MenuApi) CreateMenu(ctx *gin.Context) {

	var response common.Response
	var requestStruct req.CreateMenuRequest
	defer func() {
		if p := recover(); p != nil {
			err := fmt.Errorf("internal error: %v", p)
			response.ResultWithError(ctx, common.RequestError, err)
			return
		}
	}()
	ctx.BindJSON(&requestStruct)
	c, exit := ctx.Get("claim")
	if exit {
		global.Log.Infof("%+v\n", c)
	}
	var calim *jwtauth.MyCustomClaims = c.(*jwtauth.MyCustomClaims)
	if calim.Role != models.PermissionAdmin {
		response.ResultWithError(ctx, common.RequestError, errors.New("暂无权限"))
		return
	}

	var parent_model models.MenuModel
	if requestStruct.ParentId != nil && *requestStruct.ParentId != 0 {
		parent_model.ID = *requestStruct.ParentId
		fmt.Printf("%+v", parent_model)
		err := global.DB.Model(&parent_model).Association("ChildrenMenu").Append(&models.MenuModel{
			MenuName: requestStruct.MenuName,
			MenuType: requestStruct.MenuType,
			ParentId: requestStruct.ParentId,
			UserId:   requestStruct.UserId,
			UserName: requestStruct.UserName,
		})
		if err != nil {
			response.ResultWithError(ctx, common.ErrorStatus, err)
			return
		}

		response.ResultOk(ctx, "")
		return

	}
	global.DB.Create(&models.MenuModel{
		MenuName: requestStruct.MenuName,
		MenuType: requestStruct.MenuType,
		ParentId: requestStruct.ParentId,
		UserId:   requestStruct.UserId,
		UserName: requestStruct.UserName,
	})
	response.ResultOk(ctx, "")
}
