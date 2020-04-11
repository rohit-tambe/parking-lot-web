package api

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/rohit-tambe/parking-lot-web/model"
	"github.com/rohit-tambe/parking-lot-web/service"
)

// GetCarList get parked car list
func GetCarList(c echo.Context) error {
	responseDTO := model.ResponseDTO{Status: "Successful"}
	response, err := service.RagistrationMuster.GetParkingLotStatus()
	if err != nil {
		responseDTO.GetFailResponse(err)
		return c.JSON(http.StatusOK, responseDTO)
	}
	details := responseDTO.Details
	for _, muster := range response {
		details = append(details, model.Detail{muster.Slot, muster.RagistrationNo, muster.Colour})
	}
	responseDTO.Details = details
	return c.JSON(http.StatusOK, responseDTO)
}
