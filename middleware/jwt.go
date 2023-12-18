package middleware

import (
	"errors"
	"go_blog/common"
	"go_blog/global"
	"go_blog/utils"
	jwtauth "go_blog/utils/jwt_auth"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		pathList := strings.Split(ctx.Request.URL.Path, "/")
		lastPath := pathList[len(pathList)-1]
		if utils.Find[string](global.Config.InterceptApi.InterceptPath, lastPath) {
			token := ctx.Request.Header.Get("token")
			response := common.Response{}
			if token == "" {
				response.ResultWithError(ctx, common.RequestError, errors.New("请求未携带token"))
				ctx.Abort()
				return
			}

			j := jwtauth.Jwt{}
			claim, err := j.ParseToken(token)

			if err != nil {
				if err.Error() == "claim invalid" || err.Error() == "invalid claim type" {
					response.ResultWithError(ctx, common.ErrorStatus, errors.New("系统解析错误"))
					ctx.Abort()
					return
				}
				response.ResultWithError(ctx, common.RequestError, err)
				ctx.Abort()
				return
			}
			ctx.Set("claim", claim)
		}

	}
}
