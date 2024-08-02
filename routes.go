package main

import (
	"hellow/controller"
	"hellow/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Registe)
	r.POST("/api/auth/login", controller.Login)

	r.POST("/api/auth/vip", controller.Vip_Time)

	return r

}
