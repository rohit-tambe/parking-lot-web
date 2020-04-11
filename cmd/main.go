package main

import (
	"fmt"
	"log"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rohit-tambe/parking-lot-web/api"
)

func routes(serviceGroup *echo.Group) {
	serviceGroup.POST("/createParkingLot", api.CreateParkingLot)
	serviceGroup.POST("/park", api.ParkTheCar)
	serviceGroup.POST("/leave", api.LeaveTheCar)
	serviceGroup.GET("/status", api.GetCarList)

}
func main() {
	server := echo.New()

	server.File("/web", "web/index.html")
	server.Static("/js", "web/js")
	server.Static("/css", "web/css")
	// ragister routing group
	serviceGroup := server.Group("/parking-service", middleware.Logger())
	routes(serviceGroup)
	server.Validator = &CustomValidator{validator: validator.New()}
	serviceGroup.Use(middleware.BodyDump(BodyDump))
	// Start the Server
	if err := server.Start(":8081"); err != nil {
		log.Fatal(err)
	}
}

// BodyDump LastRequest
func BodyDump(c echo.Context, reqBody []byte, resBody []byte) {
	fmt.Println("Request", string(reqBody))
	fmt.Println("Response", string(resBody))
}

// CustomValidator validate api request
type CustomValidator struct {
	validator *validator.Validate
}

// Validate validate struct
func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
