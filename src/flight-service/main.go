package main

import (
	"database/sql"
	"fmt"
	"lab2/src/flight-service/dbhandler"
	"lab2/src/flight-service/handlers"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"postgres", 5432, "postgres", "flights", "postgres")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	flightHandler := &handlers.FlightHandler{
		DBHandler: dbhandler.InitDBHandler(db),
	}

	router := gin.Default()

	router.GET("/manage/health", flightHandler.GetHealth)

	router.GET("/api/v1/flights", flightHandler.GetAllFlightsHandler)
	router.GET("/api/v1/flight/:flightNumber", flightHandler.GetFlightHandler)
	router.GET("/api/v1/flight/airport/:airportId", flightHandler.GetAirportHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8060"
	}

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
