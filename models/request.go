package models

type UserReq struct {
	UserName string `form:"user_name"`
	PassWord string `form:"pass_word"`
}

type TaskReq struct {
	TaskName string `form:"task_name"`
	Status   bool   `form:"status"`
	Content  string `form:"content"`
	StartAt  string `form:"start_at"`
	EndAt    string `form:"end_at"`
}
