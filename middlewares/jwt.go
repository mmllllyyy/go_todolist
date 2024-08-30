package middlewares

import (
	"github.com/gin-gonic/gin"
	"time"
	"todolist/models"
	"todolist/utils"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("token")
		if token == "" {
			utils.Logger.Error("token missing")
			models.FailResponse(ctx, "token missing", "jwt", nil)
			ctx.Abort()
			return
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			utils.Logger.Error("parse token error")
			models.FailResponse(ctx, "parse token error", "jwt", nil)
			ctx.Abort()
			return
		} else if time.Now().Unix() > claims.ExpiresAt {
			utils.Logger.Error("token has been expired")
			models.FailResponse(ctx, "token has been expired", "jwt", nil)
			ctx.Abort()
			return
		}
		ctx.Request = ctx.Request.WithContext(utils.NewContext(ctx.Request.Context(), &utils.UserInfo{Id: claims.UserID}))

		utils.Logger.Info("token assigned for user id:", claims.UserID)
		ctx.Next()
	}
}
