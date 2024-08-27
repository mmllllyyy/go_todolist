package service

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"sync"
	"todolist/models"
	"todolist/repositories/db"
	"todolist/utils"
)

type UserSrv struct{}

var userSrvIns *UserSrv
var userSrvOnce sync.Once

// GetUsrSrv 单例模式实现
func GetUsrSrv() *UserSrv {
	userSrvOnce.Do(func() {
		userSrvIns = &UserSrv{}
	})
	return userSrvIns
}

func (u UserSrv) Register(ctx *gin.Context, req *utils.CreateUserReq) (utils.ResponseVo, error) {
	dao := db.NewUserDao(ctx)
	_, err := dao.FindUserByName(req.UserName)
	if err == nil {
		return utils.FailVo("Fail", "User name is already existed ", req.UserName), err
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return utils.FailVo("Fail", "User register error:"+err.Error(), nil), err
	}

	err = dao.CreateUser(&models.User{
		UserName: req.UserName,
		PassWord: req.PassWord,
	})
	if err != nil {
		return utils.FailVo("Fail", "User Register "+err.Error(), req.UserName), err
	}

	return utils.SuccessVo("User Register Success", req.UserName), nil
}
