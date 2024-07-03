package server

import (
	"github.com/LuisGerardoDC/personalFinances/app/src/server/paths"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	router := gin.Default()
	config := cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	}
	router.Use(cors.New(config))
	paths.AddBudgetRoutes(router)
	return router
}
