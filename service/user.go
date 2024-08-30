package service

import (
	"context"
	"errors"
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

func (u UserSrv) Register(ctx context.Context, req *models.UserReq) (models.ResponseVo, error) {
	dao := db.NewUserDao(ctx)
	_, err := dao.FindUserByName(req.UserName)
	if err == nil {
		utils.Logger.Info("register failed user name:" + req.UserName + " has been occupied")
		return models.FailVo("Fail", "User name is already existed ", req.UserName), err
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		utils.Logger.Error("gorm find user by name failed in register:" + err.Error())
		return models.FailVo("Fail", "User register error:"+err.Error(), nil), err
	}

	err = dao.CreateUser(&models.User{
		UserName: req.UserName,
		PassWord: req.PassWord,
	})
	if err != nil {
		utils.Logger.Error("create user failed:" + err.Error())
		return models.FailVo("Fail", "User Register "+err.Error(), req.UserName), err
	}

	utils.Logger.Info("user " + req.UserName + " register OK")
	return models.SuccessVo("User Register Success", req.UserName), nil
}

func (u UserSrv) Login(ctx context.Context, req *models.UserReq) (models.ResponseVo, error) {
	dao := db.NewUserDao(ctx)
	user, err := dao.FindUserByName(req.UserName)
	if err != nil {
		utils.Logger.Error("gorm find user by name failed in login:" + err.Error())
		return models.FailVo("Fail in login "+err.Error(), "Login", nil), err
	}
	if user.PassWord != req.PassWord {
		utils.Logger.Info("login failed password do not match")
		return models.FailVo("wrong password", "Login", nil), nil
	}
	token, err := utils.CreateToken(user.ID, user.UserName)
	if err != nil {
		utils.Logger.Error("create token error in login:", err.Error())
		return models.FailVo("create token error "+err.Error(), "Login", nil), err
	}
	utils.Logger.Info("login ", user.UserName, " OK")
	return models.SuccessVo("login success", token), nil
}
