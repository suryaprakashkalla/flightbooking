package repositories

import (
	"database/sql"
	//"net/http"
	//"flight/models"
	//"strconv"
)

type FlightRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewFlightRepository(dbHAndler *sql.DB) *FlightRepository {
	return &FlightRepository{
		dbHandler: dbHAndler,
	}
}