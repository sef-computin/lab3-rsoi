package main

import (
	"lab2/src/gateway-service/handlers"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	// r := handlers.Router()
	servicesConfig := handlers.ServicesStruct{
		TicketServiceAddress: "http://ticket-service:8070",
		FlightServiceAddress: "http://flight-service:8060",
		BonusServiceAddress:  "http://bonus-service:8050",
	}

	router := gin.Default()
	gs := handlers.NewGatewayService(&servicesConfig)

	router.GET("/manage/health", gs.GetHealth)

	router.GET("/api/v1/flights", gs.GetAllFlights)
	router.GET("/api/v1/me", gs.GetUserInfo)
	router.GET("/api/v1/tickets", gs.GetUserTickets)
	router.GET("/api/v1/tickets/:ticketUid", gs.GetUserTicket)
	router.POST("/api/v1/tickets", gs.BuyTicket)
	router.DELETE("/api/v1/tickets/:ticketUid", gs.CancelTicket)
	router.GET("/api/v1/privilege", gs.GetPrivilege)

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
