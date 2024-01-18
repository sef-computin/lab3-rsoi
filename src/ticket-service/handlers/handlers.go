package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"lab2/src/ticket-service/dbhandler"
	"lab2/src/ticket-service/models"

	"github.com/gin-gonic/gin"
)

type TicketHandler struct {
	DBHandler dbhandler.TicketDB
}

func (h *TicketHandler) GetTicketsByUsernameHandler(c *gin.Context) {
	username := c.Param("username")
	tickets, err := h.DBHandler.GetTicketsByUsername(username)
	if err != nil {
		log.Printf("Failed to get ticket: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, tickets)
}

func (h *TicketHandler) BuyTicketHandler(c *gin.Context) {
	var ticket models.Ticket

	err := json.NewDecoder(c.Request.Body).Decode(&ticket)
	if err != nil {
		c.IndentedJSON(http.StatusBadGateway, nil)
		return
	}

	if err := h.DBHandler.CreateTicket(&ticket); err != nil {
		log.Printf("Failed to create ticket: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (h *TicketHandler) DeleteTicket(c *gin.Context) {
	uid := c.Param("uid")
	err := h.DBHandler.DeleteTicketByUID(uid)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
	}
	c.IndentedJSON(http.StatusOK, nil)
}

func (h *TicketHandler) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}
