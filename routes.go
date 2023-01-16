package main

import (
	"hellow/controller"
	"hellow/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine{
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/test",controller.WwWw)
    r.POST("/api/auth/login", controller.Login)
	
	return r

}