package server

import (
	"database/sql"
	"log"
	"flight/controllers"
	"flight/repositories"
	"flight/services"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config            *viper.Viper
	router            *gin.Engine
	flightController *controllers.FlightController
	
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	flightRepository := repositories.NewFlightRepository(dbHandler)
	//resultRepository := repositories.NewResultsRepository(dbHandler)
	flightService := services.NewFlightService(flightRepository)
	//resultsService := services.NewResultsService(resultRepository, runnersRepository)
	flightController := controllers.NewFlightController(flightService)
	//resultsController := controllers.NewResultsController(resultsService)

	router := gin.Default()

	
       router.GET("/flight",flightController.GetFlight)
	// router.POST("/runner", runnersController.CreateRunner)
	// router.PUT("/runner", runnersController.UpdateRunner)
	// router.DELETE("/runner/:id", runnersController.DeleteRunner)
	// router.GET("/runner/:id", runnersController.GetRunner)
	// router.GET("/runner", runnersController.GetRunnersBatch)

	// router.POST("/result", resultsController.CreateResult)
	// router.DELETE("/result/:id", resultsController.DeleteResult)

	return HttpServer{
		config:            config,
		router:            router,
		flightController: flightController,
		
	}
}

func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
