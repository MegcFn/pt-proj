package router

import "github.com/gin-gonic/gin"

func SetupRouters(group *gin.RouterGroup) {
	LoadApplication(group)
}
