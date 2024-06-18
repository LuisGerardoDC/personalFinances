package handlers

import (
	models "github.com/LuisGerardoDC/personalFinances/app/src/models/mssql"
	"github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"
	"github.com/gin-gonic/gin"
)

type NewBudgetHandler struct {
	useCase *usecases.CreateNewBudget
}

func (h *NewBudgetHandler) CreateNewBudget(c *gin.Context) {
	h.useCase.CreateNewBudget(models.Budget{})
}
