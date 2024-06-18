package main

import (
	"github.com/LuisGerardoDC/personalFinances/app/src/server"
	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(gin.DebugMode)
	router := server.GetRouter()
	router.Run(":8080")
}
