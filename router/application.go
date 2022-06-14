package router

import (
	"github.com/MegcFn/pt-proj/router/handler"
	"github.com/gin-gonic/gin"
)

func LoadApplication(group *gin.RouterGroup) {

	group.POST("/applications", handler.CreateApplication)
	group.GET("/applications", handler.FindApplications)
	group.GET("/applications/:sid", handler.FindApplicationsBySid)
	group.PUT("/applications/:id", handler.UpdateApplication)
	group.DELETE("applications/:id", handler.DeleteApplication)
}
