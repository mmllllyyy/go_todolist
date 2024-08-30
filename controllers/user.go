package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/models"
	"todolist/service"
	"todolist/utils"
)

func UserRegisterHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.UserReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			utils.Logger.Error("Register user error when bind request:", err.Error())
			ctx.JSON(http.StatusBadRequest, models.FailVo("Register user error when bind request:"+err.Error(), "Register user bind request", nil))
			return
		}
		u := service.GetUsrSrv()
		resp, err := u.Register(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Error("Register user error when register:", models.FailVo("Register user error:"+err.Error(), "Register user", nil))
			ctx.JSON(http.StatusBadRequest, resp)
			return
		}
		utils.Logger.Info("Register OK")
		ctx.JSON(http.StatusOK, resp)

	}
}

func UserLoginHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req models.UserReq
		err := ctx.ShouldBind(&req)
		if err != nil {
			utils.Logger.Error("Login error when bind request:", err)
			ctx.JSON(http.StatusBadRequest, err.Error())
			return
		}
		u := service.GetUsrSrv()
		resp, err := u.Login(ctx.Request.Context(), &req)
		if err != nil {
			utils.Logger.Error("Login error when request:", err)
			ctx.JSON(http.StatusBadRequest, resp)
			return
		}
		utils.Logger.Info("Login OK")
		ctx.JSON(http.StatusOK, resp)
	}
}
