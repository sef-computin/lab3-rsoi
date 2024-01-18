package handlers

import (
	"encoding/json"
	"lab2/src/gateway-service/models"
	"lab2/src/gateway-service/service"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ServicesStruct struct {
	TicketServiceAddress string
	FlightServiceAddress string
	BonusServiceAddress  string
}

type GatewayService struct {
	Config ServicesStruct
}

func NewGatewayService(config *ServicesStruct) *GatewayService {
	return &GatewayService{Config: *config}
}

func (gs *GatewayService) GetAllFlights(c *gin.Context) {
	params := c.Request.URL.Query()

	flights, err := service.GetAllFlightsInfo(gs.Config.FlightServiceAddress)
	if err != nil {
		log.Printf("failed to get response from flights service: %v\n", err)
		c.IndentedJSON(http.StatusNoContent, flights)
		return
	}

	pageParam := params.Get("page")
	if pageParam == "" {
		log.Println("invalid query parameter: (page)")
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	page, err := strconv.Atoi(pageParam)
	if err != nil {
		log.Printf("unable to convert the string into int: %s\n", err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	sizeParam := params.Get("size")
	if sizeParam == "" {
		log.Println("invalid query parameter (size)")
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}
	size, err := strconv.Atoi(sizeParam)
	if err != nil {
		log.Printf("unable to convert the string into int:  %v\n", err)
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	right := page * size
	if len(*flights) < right {
		right = len(*flights)
	}

	flightsStripped := (*flights)[(page-1)*size : right]
	cars := models.FlightsLimited{
		Page:          page,
		PageSize:      size,
		TotalElements: len(flightsStripped),
		Items:         &flightsStripped,
	}

	c.IndentedJSON(http.StatusOK, cars)
}

func (gs *GatewayService) GetUserInfo(c *gin.Context) {
	username := c.Request.Header.Get("X-User-Name")
	if username == "" {
		log.Printf("Username header is empty\n")
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	userInfo, err := service.UserInfoController(
		gs.Config.TicketServiceAddress,
		gs.Config.FlightServiceAddress,
		gs.Config.BonusServiceAddress,
		username,
	)

	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, userInfo)
}

func (gs *GatewayService) GetUserTickets(c *gin.Context) {
	username := c.Request.Header.Get("X-User-Name")

	if username == "" {
		log.Printf("Username header is empty\n")
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	ticketsInfo, err := service.UserTicketsController(
		gs.Config.TicketServiceAddress,
		gs.Config.FlightServiceAddress,
		username,
	)

	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, ticketsInfo)
}

func (gs *GatewayService) GetUserTicket(c *gin.Context) {
	username := c.Request.Header.Get("X-User-Name")

	if username == "" {
		log.Printf("Username header is empty\n")
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	ticketUID := c.Param("ticketUid")

	ticketsInfo, err := service.UserTicketsController(
		gs.Config.TicketServiceAddress,
		gs.Config.FlightServiceAddress,
		username,
	)

	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	var ticketInfo *models.TicketInfo
	for _, ticket := range *ticketsInfo {
		if ticket.TicketUID == ticketUID {
			ticketInfo = &ticket
		}
	}

	if ticketInfo == nil {
		log.Printf("Ticket not found")
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, ticketInfo)
}

func (gs *GatewayService) BuyTicket(c *gin.Context) {
	username := c.Request.Header.Get("X-User-Name")

	if username == "" {
		log.Printf("Username header is empty\n")
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	var ticketInfo models.BuyTicketInfo

	err := json.NewDecoder(c.Request.Body).Decode(&ticketInfo)
	if err != nil {
		log.Printf("failed to decode post request")
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	ticket, err := service.BuyTicketController(
		gs.Config.TicketServiceAddress,
		gs.Config.FlightServiceAddress,
		gs.Config.BonusServiceAddress,
		username,
		&ticketInfo,
	)

	if err != nil {
		log.Printf("failed to get response: %v\n", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, ticket)
}

func (gs *GatewayService) CancelTicket(c *gin.Context) {
	username := c.Request.Header.Get("X-User-Name")

	if username == "" {
		log.Printf("Username header is empty\n")
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	ticketUID := c.Param("ticketUid")

	err := service.CancelTicketController(
		gs.Config.TicketServiceAddress,
		gs.Config.BonusServiceAddress,
		username,
		ticketUID,
	)

	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusNoContent, nil)
}

func (gs *GatewayService) GetPrivilege(c *gin.Context) {
	username := c.Request.Header.Get("X-User-Name")
	if username == "" {
		log.Printf("Username header is empty\n")
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	privilegeInfo, err := service.UserPrivilegeController(
		gs.Config.BonusServiceAddress,
		username,
	)

	if err != nil {
		log.Printf("Failed to get response: %s\n", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, privilegeInfo)
}

func (gs *GatewayService) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}
