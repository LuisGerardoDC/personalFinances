package paths

import (
	"github.com/LuisGerardoDC/personalFinances/app/src/server/handlers"
	"github.com/gin-gonic/gin"
)

func AddBudgetRoutes(router *gin.Engine) {
	budgetRoutes := router.Group("/budget")
	budgetRoutes.POST("/new", handlers.ImplementedNewBudgetHandler.CreateNewBudget)
	budgetRoutes.POST("/record", handlers.ImplementedAddRecordHandler.AddNewRecord)
	budgetRoutes.DELETE("/record", handlers.ImplementedDeleteRecordHandler.DeleteRecord)
	budgetRoutes.GET("/:id", handlers.ImplementedGetBudget.GetBudget)
	budgetRoutes.GET("/:userid", handlers.ImplementedGetBudgets.GetBudgets)

}
