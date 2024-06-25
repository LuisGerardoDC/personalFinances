package paths

import (
	"github.com/LuisGerardoDC/personalFinances/app/src/server/handlers"
	"github.com/gin-gonic/gin"
)

func AddBudgetRoutes(router *gin.Engine) {
	budgetRoutes := router.Group("/budget")
	budgetRoutes.POST("/new", handlers.ImplementedNewBudgetHandler.CreateNewBudget)
	budgetRoutes.POST("/add/record/", handlers.ImplementedAddRecordHandler.AddNewRecord)
	budgetRoutes.GET("/:id", handlers.ImplementedGetBudget.GetBudget)
}
