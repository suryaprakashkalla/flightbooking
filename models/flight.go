package models

import "time"

type Flight struct {
    FlightID        int       `json:"flight_id"`
    FlightNumber    string    `json:"flight_number"`
    DepartureCity   string    `json:"departure_city"`
    ArrivalCity     string    `json:"arrival_city"`
    DepartureTime   time.Time `json:"departure_time"`
    ArrivalTime     time.Time `json:"arrival_time"`
    Price           float64   `json:"price"`
}
