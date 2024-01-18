package service

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"lab2/src/gateway-service/models"
)

func CheckFlightHealth() (interface{}, error) {
	requestURL := "http://flight-service:8060/manage/health"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)

	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	resp, err := client.Do(req)
	return resp, err
}

func GetFlight(flightServiceAddress, flightNumber string) (*models.Flight, error) {

	_, herr := flightcb.Execute(CheckFlightHealth)
	if herr != nil {
		return &models.Flight{}, herr
	}

	requestURL := fmt.Sprintf("%s/api/v1/flight/%s", flightServiceAddress, flightNumber)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println("Failed to create an http request")
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed request to flight service: %w", err)
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Failed to close response body")
		}
	}(res.Body)

	flight := &models.Flight{}
	if err = json.NewDecoder(res.Body).Decode(flight); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return flight, nil
}

func GetAllFlightsInfo(flightServiceAddress string) (*[]models.FlightInfo, error) {

	_, herr := flightcb.Execute(CheckFlightHealth)
	if herr != nil {
		return &[]models.FlightInfo{}, herr
	}

	requestURL := fmt.Sprintf("%s/api/v1/flights", flightServiceAddress)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println("Failed to create an http request")
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed request to flight service: %w", err)
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Failed to close response body")
		}
	}(res.Body)

	flights := &[]models.Flight{}
	if err = json.NewDecoder(res.Body).Decode(flights); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	flightsInfo := make([]models.FlightInfo, 0)
	for _, flight := range *flights {
		airportFrom, err := GetAirport(flightServiceAddress, flight.FromAirportId)
		if err != nil {
			return nil, fmt.Errorf("failed to get airport: %s", err)
		}

		airportTo, err := GetAirport(flightServiceAddress, flight.ToAirportId)
		if err != nil {
			return nil, fmt.Errorf("failed to get airport: %s", err)
		}

		fInfo := models.FlightInfo{
			FlightNumber: flight.FlightNumber,
			FromAirport:  fmt.Sprintf("%s %s", airportFrom.City, airportFrom.Name),
			ToAirport:    fmt.Sprintf("%s %s", airportTo.City, airportTo.Name),
			Date:         flight.Date,
			Price:        flight.Price,
		}

		flightsInfo = append(flightsInfo, fInfo)
	}

	return &flightsInfo, nil
}

func GetAirport(flightServiceAddress string, airportID int) (*models.Airport, error) {

	_, herr := flightcb.Execute(CheckFlightHealth)
	if herr != nil {
		return &models.Airport{City: "Error", Name: "error"}, herr
	}

	requestURL := fmt.Sprintf("%s/api/v1/flight/airport/%d", flightServiceAddress, airportID)

	req, err := http.NewRequest(http.MethodGet, requestURL, nil)
	if err != nil {
		fmt.Println("Failed to create an http request")
		return nil, err
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed request to flight service: %w", err)
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Failed to close response body")
		}
	}(res.Body)

	airport := &models.Airport{}
	if err = json.NewDecoder(res.Body).Decode(airport); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return airport, nil
}
