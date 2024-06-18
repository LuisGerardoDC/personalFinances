package handlers

import (
	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
	"github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"
	"github.com/gin-gonic/gin"
)

type NewBudgetHandler struct {
	useCase *usecases.CreateNewBudget
}

func (h *NewBudgetHandler) CreateNewBudget(c *gin.Context) {
	var reqBudget requestModel.Budget

	if err := c.ShouldBindBodyWithJSON(&reqBudget); err != nil {
		c.JSON(400, responseModel.Budget{
			Error: "Invalid JSON Format",
		})
		return
	}

	resBudget := h.useCase.CreateNewBudget(reqBudget)

	c.JSON(resBudget.Code, resBudget)
}
