package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"lab2/src/bonus-service/dbhandler"
	"lab2/src/bonus-service/models"

	"github.com/gin-gonic/gin"
)

type BonusHandler struct {
	DBHandler dbhandler.BonusDB
}

func (h *BonusHandler) CreatePrivilegeHistoryHandler(c *gin.Context) {
	var record models.PrivilegeHistory

	err := json.NewDecoder(c.Request.Body).Decode(&record)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	if err := h.DBHandler.CreateHistoryRecord(&record); err != nil {
		log.Printf("Failed to create ticket: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (h *BonusHandler) UpdatePrivilegeHandler(c *gin.Context) {
	var record models.Privilege

	err := json.NewDecoder(c.Request.Body).Decode(&record)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	if err := h.DBHandler.UpdatePrivilege(&record); err != nil {
		log.Printf("Failed to update ticket: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (h *BonusHandler) CreatePrivilegeHandler(c *gin.Context) {
	var record models.Privilege

	err := json.NewDecoder(c.Request.Body).Decode(&record)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, nil)
		return
	}

	if err := h.DBHandler.CreatePrivilege(&record); err != nil {
		log.Printf("Failed to create ticket: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, nil)
}

func (h *BonusHandler) GetPrivilegeByUsernameHandler(c *gin.Context) {

	username := c.Param("username")
	privilege, err := h.DBHandler.GetPrvilegeByUsername(username)
	if err != nil {
		log.Printf("failed to get flghts: %s", err)
		c.IndentedJSON(http.StatusNotFound, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, privilege)
}

func (h *BonusHandler) GetHistoryByIdHandler(c *gin.Context) {

	privilegeId := c.Param("privilegeId")
	history, err := h.DBHandler.GetHistoryById(privilegeId)
	if err != nil {
		log.Printf("failed to get flghts: %s", err)
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, history)
}

func (h *BonusHandler) GetHealth(c *gin.Context) {
	c.Status(http.StatusOK)
}
