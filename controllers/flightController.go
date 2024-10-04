package controllers

import (
	//"encoding/json"
	//"io"
	//"log"
	//"net/http"
	//"flight/models"
	"flight/services"

	//"github.com/gin-gonic/gin"
)

type FlightController struct {
	flightService *services.FlightService
}

func NewFlightController(flightService *services.FlightService) *FlightController {//initialising new instance
	return &FlightController{
		flightService: flightService,
	}
}