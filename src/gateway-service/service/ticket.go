package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"lab2/src/gateway-service/models"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CheckTicketHealth() (interface{}, error) {
	requestURL := "http://ticket-service:8070/manage/health"
	req, err := http.NewRequest(http.MethodGet, requestURL, nil)

	if err != nil {
		return req, err
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	resp, err := client.Do(req)
	return resp, err
}

func GetUserTickets(ticketsServiceAddress, username string) (*[]models.Ticket, error) {

	_, herr := ticketcb.Execute(CheckTicketHealth)
	if herr != nil {
		return &[]models.Ticket{}, herr
	}

	requestURL := fmt.Sprintf("%s/api/v1/tickets/%s", ticketsServiceAddress, username)

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

	tickets := &[]models.Ticket{}
	if err = json.NewDecoder(res.Body).Decode(tickets); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return tickets, nil
}

func CreateTicket(ticketsServiceAddress, username, flightNumber string, price int) (string, error) {

	_, herr := ticketcb.Execute(CheckTicketHealth)
	if herr != nil {
		return "", herr
	}

	requestURL := fmt.Sprintf("%s/api/v1/tickets", ticketsServiceAddress)
	uid := uuid.New().String()

	ticket := &models.Ticket{
		TicketUID:    uid,
		FlightNumber: flightNumber,
		Status:       "PAID",
		Username:     username,
		Price:        price,
	}

	data, err := json.Marshal(ticket)
	if err != nil {
		return "", fmt.Errorf("encoding error: %w", err)
	}

	req, err := http.NewRequest(http.MethodPost, requestURL, bytes.NewReader(data))
	if err != nil {
		fmt.Println("Failed to create an http request")
		return "", err
	}

	client := &http.Client{
		Timeout: 10 * time.Minute,
	}

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed request to flight service: %w", err)
	}

	defer func(Body io.ReadCloser) {
		if err := Body.Close(); err != nil {
			fmt.Println("Failed to close response body")
		}
	}(res.Body)

	return uid, nil
}
