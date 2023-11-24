package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"lab2/src/ticket-service/dbhandler/mock_dbhandler"
	"lab2/src/ticket-service/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/golang/mock/gomock"

	_ "github.com/lib/pq"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetTicketsByUsername(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockTicketDB, input *models.Ticket)
	tests := []struct {
		name                 string
		input                *models.Ticket
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{

			name: "Ok",
			input: &models.Ticket{
				TicketUID:    "1",
				Username:     "TestMax",
				FlightNumber: "123",
				Price:        1500,
				Status:       "PAID",
			},
			username: "TestMax",
			mockBehaviour: func(r *mock_dbhandler.MockTicketDB, input *models.Ticket) {
				r.EXPECT().CreateTicket(input)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			dbhandler := mock_dbhandler.NewMockTicketDB(c)
			test.mockBehaviour(dbhandler, test.input)

			handler := TicketHandler{DBHandler: dbhandler}

			r := gin.New()
			r.POST("/api/v1/tickets", handler.BuyTicketHandler)

			jsonValue, _ := json.Marshal(test.input)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/api/v1/tickets", bytes.NewBuffer(jsonValue))
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}

func TestBuyTicket(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockTicketDB, username string, output []*models.Ticket)
	tests := []struct {
		name                 string
		output               []*models.Ticket
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:     "Ok",
			username: "TestMax",
			mockBehaviour: func(r *mock_dbhandler.MockTicketDB, username string, output []*models.Ticket) {
				r.EXPECT().GetTicketsByUsername(username).Return(output, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			var tickets []*models.Ticket

			dbhandler := mock_dbhandler.NewMockTicketDB(c)
			test.mockBehaviour(dbhandler, test.username, tickets)

			handler := TicketHandler{DBHandler: dbhandler}

			r := gin.New()
			r.GET("/api/v1/tickets/:username", handler.GetTicketsByUsernameHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", fmt.Sprintf("/api/v1/tickets/%v", test.username), nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}
