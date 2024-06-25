package handlers

import (
	"strconv"

	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
	"github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"
	"github.com/gin-gonic/gin"
)

type GetBudgetHandler struct {
	useCase *usecases.GetBudget
}

func (h *GetBudgetHandler) GetBudget(c *gin.Context) {
	var budgetID int64
	stingID := c.Param("id")
	budgetID, err := strconv.ParseInt(stingID, 10, 64)

	if err != nil {
		c.JSON(400, responseModel.Response{
			Error:  err.Error(),
			Succes: false,
		})
	}

	response := h.useCase.GetBudget(budgetID)

	c.JSON(response.Code, response)

}
