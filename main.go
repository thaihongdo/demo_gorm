package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"togo_pre/config"
	"togo_pre/models"
	"togo_pre/router"
)

func main() {
	cf := config.InitFromFile("")
	closeFunc, errFile := models.InitFromSQLLite(cf.DbConnection)
	if errFile != nil {
		log.Fatalf("Read file Database error")
	}
	routersInit := router.InitRouter(cf.EnvironmentPrefix)
	endPoint := fmt.Sprintf(":%d", cf.ServerPort)
	maxHeaderBytes := 1 << 20
	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    time.Minute,
		WriteTimeout:   time.Minute,
		MaxHeaderBytes: maxHeaderBytes,
	}

	log.Printf("[info] start http server listening %s", endPoint)
	_ = server.ListenAndServe()
	defer closeFunc()
}
