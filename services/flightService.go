package services

import (
	//"net/http"
	//"flight/models"
	"flight/repositories"
	//"time"
)

type FlightService struct {
	flightRepository *repositories.FlightRepository
	//runnersRepository *repositories.RunnersRepository
}

func NewFlightService(flightRepository *repositories.FlightRepository) *FlightService {

	return &FlightService{
		flightRepository: flightRepository,
		//runnersRepository: runnersRepository,
	}
}