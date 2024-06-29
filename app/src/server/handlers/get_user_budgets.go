package handlers

import (
	"strconv"

	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
	"github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"
	"github.com/gin-gonic/gin"
)

type GetUserBudgetsHandler struct {
	useCase *usecases.GetUserBudgets
}

func (g *GetUserBudgetsHandler) GetBudgets(c *gin.Context) {
	stingID := c.Param("userid")
	userID, err := strconv.ParseInt(stingID, 10, 64)
	if err != nil {
		c.JSON(400, responseModel.Response{
			Error:  err.Error(),
			Succes: false,
		})
	}

	response := g.useCase.GetBudgets(userID)
	c.JSON(response.Code, response)

}
