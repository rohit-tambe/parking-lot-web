package api

import (
	"net/http"

	"github.com/rohit-tambe/parking-lot-web/model"

	"github.com/labstack/echo/v4"
	"github.com/rohit-tambe/parking-lot-web/service"
)

// ParkTheCar park the car with details
func ParkTheCar(c echo.Context) error {
	requestDTO := model.CarDetails{}
	if err := c.Bind(&requestDTO); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	// validate
	if err := c.Validate(&requestDTO); err != nil {
		return c.JSON(http.StatusOK, err.Error())
	}
	responseDTO := model.ResponseDTO{Status: "Successful"}
	response, err := service.RagistrationMuster.ParkCar(requestDTO.RagistrationNo, requestDTO.Colour)
	if err != nil {
		responseDTO.GetFailResponse(err)
		return c.JSON(http.StatusOK, responseDTO)
	}
	responseDTO.Result = response
	return c.JSON(http.StatusOK, responseDTO)
}
