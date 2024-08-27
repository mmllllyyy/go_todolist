package utils

import "time"

type CreateUserReq struct {
	UserName string `form:"user_name"`
	PassWord string `form:"pass_word"`
}

type UpdateUserReq struct {
	UserName    string `form:"user_name"`
	OriPassWord string `form:"ori_pass_word"`
	NewPassWord string `form:"new_pass_word"`
}

type DeleteUserReq struct {
	UserName string `form:"user_name"`
	PassWord string `form:"pass_word"`
}

type CreateTaskReq struct {
	UserID   int       `form:"user_id"`
	TaskName string    `form:"task_name"`
	Status   bool      `form:"status"`
	Content  string    `form:"content"`
	StartAt  time.Time `form:"start_at"`
	EndAt    time.Time `form:"end_at"`
}

type UpdateTaskReq struct {
	UserID   int       `form:"user_id"`
	TaskName string    `form:"task_name"`
	Status   bool      `form:"status"`
	Content  string    `form:"content"`
	StartAt  time.Time `form:"start_at"`
	EndAt    time.Time `form:"end_at"`
}

type DeleteTaskReq struct {
	UserID   int       `form:"user_id"`
	TaskName string    `form:"task_name"`
	Status   bool      `form:"status"`
	Content  string    `form:"content"`
	StartAt  time.Time `form:"start_at"`
	EndAt    time.Time `form:"end_at"`
}
