package main

import (
	"github.com/MegcFn/pt-proj/router"
	"github.com/MegcFn/pt-proj/util"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	util.Migrate()
	v1 := r.Group("v1/")
	router.SetupRouters(v1)
	r.Run()
}
