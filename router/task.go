package router

import (
	"togo_pre/api"
	"togo_pre/utils"

	"github.com/gin-gonic/gin"
)

func initTaskRouter(Router *gin.RouterGroup) {
	taskRouter := Router.Group("task").Use(utils.JWTAuth())
	{
		taskRouter.POST("", api.AddTask) //login
	}

}
