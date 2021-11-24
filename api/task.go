package api

import (
	"net/http"
	"togo_pre/msg"
	"togo_pre/services"
	"togo_pre/utils"

	"github.com/gin-gonic/gin"
)

func AddTask(c *gin.Context) {
	appG := Gin{C: c}

	claims, _ := c.Get("claims")
	waitUse := claims.(*utils.CustomClaims)

	var service services.TaskReq
	isValid := appG.BindAndValidate(&service)
	if isValid {
		service.UserID = waitUse.Id
		objRes, err := service.Add()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}
