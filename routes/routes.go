package routes

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"net/http"
	"todolist/controllers"
	"todolist/middlewares"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	// 创建一个基于cookie的会话存储,这个存储使用客户端的cookie来保存数据
	store := cookie.NewStore([]byte("todolist_cookie"))
	// 为每个请求设置了一个会话中间件。"todolist_session" 是cookie的名称
	//store 是之前创建的基于cookie的会话存储。
	r.Use(sessions.Sessions("todolist_session", store))
	v1 := r.Group("todolist")
	{
		v1.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, "ping success")
		})
		v1.POST("user/register", controllers.UserRegisterHandler())
		v1.POST("user/login", controllers.UserLoginHandler())
		authed := v1.Group("/")
		authed.Use(middlewares.JWT())
		{
			authed.POST("task/create", controllers.CreateTaskHandler())
		}
	}
	return r
}
