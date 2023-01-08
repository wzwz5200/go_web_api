package main

import (
	"hellow/controller"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine{

	r.POST("/test",controller.WwWw)

	return r

}