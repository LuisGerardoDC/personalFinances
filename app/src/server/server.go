package server

import (
	"github.com/LuisGerardoDC/personalFinances/app/src/server/paths"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	paths.AddBudgetRoutes(router)
	return router
}
