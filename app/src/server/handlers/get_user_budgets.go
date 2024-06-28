package handlers

import (
	"github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"
	"github.com/gin-gonic/gin"
)

type GetUserBudgetsHandler struct {
	useCase *usecases.GetUserBudgets
}

func (g *GetUserBudgetsHandler) GetBudgets(c *gin.Context) {

}
