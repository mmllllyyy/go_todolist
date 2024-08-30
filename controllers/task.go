package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/models"
	"todolist/service"
	"todolist/utils"
)

func CreateTaskHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var reqs models.TaskReq
		err := ctx.ShouldBind(&reqs)
		if err != nil {
			utils.Logger.Error("Create task bind request error:" + err.Error())
			ctx.JSON(http.StatusBadRequest, models.FailVo("Create task bind request error:"+err.Error(), "create task bind request", nil))
			return
		}
		l := service.GetTaskSrv()
		resp, err := l.CreateTask(ctx.Request.Context(), &reqs)
		if err != nil {
			utils.Logger.Error("Create task error:" + err.Error())
			ctx.JSON(http.StatusBadRequest, resp)
			return
		}
		utils.Logger.Info("create task done")
		ctx.JSON(http.StatusOK, resp)
	}
}
