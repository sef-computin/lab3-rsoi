package handlers

import (
	"lab2/src/flight-service/dbhandler"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type FlightHandler struct {
	DBHandler dbhandler.FlightDB
}

func (hand *FlightHandler) GetAirportHandler(c *gin.Context) {

	airportID := c.Param("airportId")

	airport, err := hand.DBHandler.GetAirportByID(airportID)

	if err != nil {
		log.Printf("failed to get airport: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, airport)
}

func (hand *FlightHandler) GetAllFlightsHandler(c *gin.Context) {

	flights, err := hand.DBHandler.GetAllFlights()
	if err != nil {
		log.Printf("failed to get flghts: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, flights)
}

func (hand *FlightHandler) GetFlightHandler(c *gin.Context) {

	flightID := c.Param("flightNumber")

	flight, err := hand.DBHandler.GetFlightByNumber(flightID)

	if err != nil {
		log.Printf("failed to get flight: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, flight)
}

func (hand *FlightHandler) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}
