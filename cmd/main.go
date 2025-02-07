package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/kevin19930919/CryptoAlert/controller"
	"github.com/kevin19930919/CryptoAlert/database"
)

var ctx = context.Background()

func main() {
	server := setupServer()
	server.Run(":8612")
}

func setupServer() *gin.Engine {
	var err error

	if err = database.StartPostgrel(); err != nil {
		panic("fail to initial db")
	}
	if err = database.StartRedis(ctx); err != nil {
		panic("fail to init redis")
	}

	router := gin.Default()
	AlertRouter := router.Group("/Alert")
	{
		AlertRouter.POST("", controller.AddAlert)
		AlertRouter.DELETE("", controller.RemoveAlert)
		// AlertRouter.GET("/get-alerts", ctl.GetAlerts)
		AlertRouter.PATCH("", controller.UpdateAlert)
	}

	return router
}
