package router

import (
	"togo_pre/api"

	"github.com/gin-gonic/gin"
)

func initAuthenRouter(Router *gin.RouterGroup) {
	AuthenRouter := Router.Group("auth")
	{
		AuthenRouter.POST("login", api.Login) //login
	}

}
