package middleware

import (
	"errors"
	"go_blog/common"
	jwtauth "go_blog/utils/jwt_auth"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		response := common.Response{}
		if token == "" {
			response.ResultWithError(ctx, common.RequestError, errors.New("请求未携带token"))
			ctx.Abort()
			return
		}

		j := jwtauth.Jwt{}
		claims, err := j.ParseToken(token)

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
		ctx.Set("claims", claims)
	}
}
