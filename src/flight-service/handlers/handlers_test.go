package handlers

import (
	"lab2/src/flight-service/dbhandler/mock_dbhandler"
	"lab2/src/flight-service/models"
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

func TestGetAllFlights(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockFlightDB, record []*models.Flight)
	tests := []struct {
		name                 string
		output               []*models.Flight
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name: "Ok",
			mockBehaviour: func(r *mock_dbhandler.MockFlightDB, output []*models.Flight) {
				r.EXPECT().GetAllFlights().Return(output, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			var flights []*models.Flight
			dbhandler := mock_dbhandler.NewMockFlightDB(c)
			test.mockBehaviour(dbhandler, flights)

			handler := FlightHandler{DBHandler: dbhandler}

			r := gin.New()
			r.GET("/flights", handler.GetAllFlightsHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/flights", nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}
func TestGetFlight(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockFlightDB, input string, output *models.Flight)
	tests := []struct {
		name                 string
		output               *models.Flight
		flightnum            string
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			flightnum: "1",
			mockBehaviour: func(r *mock_dbhandler.MockFlightDB, input string, output *models.Flight) {
				r.EXPECT().GetFlightByNumber(input).Return(output, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			var flight *models.Flight
			dbhandler := mock_dbhandler.NewMockFlightDB(c)
			test.mockBehaviour(dbhandler, test.flightnum, flight)

			handler := FlightHandler{DBHandler: dbhandler}

			r := gin.New()
			r.GET("/flights/:flightNumber", handler.GetFlightHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/flights/1", nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}

func TestGetAirport(t *testing.T) {
	gin.SetMode(gin.ReleaseMode)

	type mockBehaviour func(r *mock_dbhandler.MockFlightDB, input string, output *models.Airport)
	tests := []struct {
		name                 string
		output               *models.Airport
		airportId            string
		username             string
		mockBehaviour        mockBehaviour
		expectedStatusCode   int
		expectedResponseBody string
	}{
		{
			name:      "Ok",
			airportId: "1",
			mockBehaviour: func(r *mock_dbhandler.MockFlightDB, input string, output *models.Airport) {
				r.EXPECT().GetAirportByID(input).Return(output, nil)
			},
			expectedStatusCode: http.StatusOK,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			c := gomock.NewController(t)
			defer c.Finish()

			var airport *models.Airport
			dbhandler := mock_dbhandler.NewMockFlightDB(c)
			test.mockBehaviour(dbhandler, test.airportId, airport)

			handler := FlightHandler{DBHandler: dbhandler}

			r := gin.New()
			r.GET("/airports/:airportId", handler.GetAirportHandler)

			w := httptest.NewRecorder()
			req := httptest.NewRequest("GET", "/airports/1", nil)
			req.Header.Set("X-User-Name", test.username)

			r.ServeHTTP(w, req)

			assert.Equal(t, w.Code, test.expectedStatusCode)
		})
	}
}
