package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/service"
	"todolist/utils"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req utils.CreateUserReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			utils.Logger.Error("Register user error when bind request:", err)
			ctx.JSON(http.StatusBadRequest, err.Error())
		} else {
			u := service.GetUsrSrv()
			resp, err := u.Register(ctx, &req)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err.Error())
				return
			}
			ctx.JSON(http.StatusOK, resp)
		}
	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
