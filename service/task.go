package service

import (
	"context"
	"sync"
	"todolist/models"
	"todolist/repositories/db"
	"todolist/utils"
)

type TaskSrv struct{}

var TaskSrvIns *TaskSrv
var TaskSrvOnce sync.Once

func GetTaskSrv() *TaskSrv {
	TaskSrvOnce.Do(func() {
		TaskSrvIns = &TaskSrv{}
	})
	return TaskSrvIns
}

func (t TaskSrv) CreateTask(ctx context.Context, req *models.TaskReq) (models.ResponseVo, error) {
	u, err := utils.GetUserInfo(ctx)
	if err != nil {
		utils.Logger.Error("Create task error when getting user info:", err.Error())
		return models.FailVo("Failed in getting userinfo", "create task", nil), err
	}
	task := models.Task{
		UserID:   u.Id,
		TaskName: req.TaskName,
		Status:   req.Status,
		Content:  req.Content,
		StartAt:  utils.ParseTime(req.StartAt),
		EndAt:    utils.ParseTime(req.EndAt),
	}
	dao := db.NewTaskDao(ctx)
	err = dao.CreateTask(&task)
	if err != nil {
		utils.Logger.Error("Failed in create task ", err.Error(), task)
		return models.FailVo("Failed in create task"+err.Error(), "create task", nil), err
	}
	utils.Logger.Info("Create task OK")
	return models.SuccessVo("Create task OK", nil), nil
}
