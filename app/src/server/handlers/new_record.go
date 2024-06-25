package handlers

import (
	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
	"github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"
	"github.com/gin-gonic/gin"
)

type AddRecordHandler struct {
	useCase *usecases.AddRecord2Budget
}

func (h *AddRecordHandler) AddNewRecord(c *gin.Context) {
	var reqRecord requestModel.Record

	if err := c.ShouldBindBodyWithJSON(&reqRecord); err != nil {
		c.JSON(400, responseModel.Response{
			Error:  err.Error(),
			Succes: false,
		})
	}

	response := h.useCase.CreateNewRecord(reqRecord)

	c.JSON(response.Code, response)

}
