package api

import (
	"net/http"
	"strconv"
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

func UpdateTask(c *gin.Context) {
	appG := Gin{C: c}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		appG.Response(http.StatusBadRequest, false, msg.GetMsg(msg.INVALID_PARAMS), nil, nil)
		return
	}
	var service services.TaskReq
	isValid := appG.BindAndValidate(&service)
	if isValid {
		service.ID = uint(id)
		objRes, err := service.Update()
		if err != nil {
			appG.Response(http.StatusBadRequest, false, err.Error(), nil, nil)
			return
		}
		appG.Response(http.StatusOK, true, msg.GetMsg(msg.SUCCESS), objRes, nil)
	}
}
