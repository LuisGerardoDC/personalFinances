package handlers

import (
	requestModel "github.com/LuisGerardoDC/personalFinances/app/src/models/request"
	responseModel "github.com/LuisGerardoDC/personalFinances/app/src/models/response"
	"github.com/LuisGerardoDC/personalFinances/app/src/server/usecases"
	"github.com/gin-gonic/gin"
)

type DeleteRecordHandler struct {
	useCase *usecases.RemoveRecord
}

func (h *DeleteRecordHandler) DeleteRecord(c *gin.Context) {
	var reqRecord requestModel.DeleteRecord

	if err := c.ShouldBindBodyWithJSON(&reqRecord); err != nil {
		c.JSON(400, responseModel.Response{
			Error:  err.Error(),
			Succes: false,
		})
	}

	response := h.useCase.DeleteRecord(reqRecord)

	c.JSON(response.Code, response)
}
