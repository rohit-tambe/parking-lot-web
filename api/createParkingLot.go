package api

import (
	"net/http"

	"github.com/rohit-tambe/parking-lot-web/model"

	"github.com/labstack/echo/v4"
	"github.com/rohit-tambe/parking-lot-web/service"
)

// CreateParkingLot crate parking lot muster
func CreateParkingLot(c echo.Context) error {
	requestDTO := model.SlotRequest{}
	responseDTO := model.ResponseDTO{Status: "Successful"}
	if err := c.Bind(&requestDTO); err != nil {
		responseDTO.GetFailResponse(err)
		return c.JSON(http.StatusOK, responseDTO)
	}
	// validate
	if err := c.Validate(&requestDTO); err != nil {
		responseDTO.GetFailResponse(err)
		return c.JSON(http.StatusOK, responseDTO)
	}
	response, err := service.RagistrationMuster.CreateParkingLot(requestDTO.Slot)
	if err != nil {
		responseDTO.GetFailResponse(err)
		return c.JSON(http.StatusOK, responseDTO)
	}
	responseDTO.Result = response
	return c.JSON(http.StatusOK, responseDTO)
}
