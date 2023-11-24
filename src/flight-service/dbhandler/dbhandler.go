package dbhandler

import (
	"database/sql"
	"errors"
	"fmt"
	"lab2/src/flight-service/models"
)

type FlightDB interface {
	GetAllFlights() ([]*models.Flight, error)
	GetFlightByNumber(flightNumber string) (*models.Flight, error)
	GetAirportByID(airportID string) (*models.Airport, error)
}

type DBHandler struct {
	db *sql.DB
}

func InitDBHandler(db *sql.DB) *DBHandler {
	return &DBHandler{
		db: db,
	}
}

func (dbhand *DBHandler) GetAllFlights() ([]*models.Flight, error) {

	var flights []*models.Flight
	rows, err := dbhand.db.Query(`SELECT * FROM flight;`)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	for rows.Next() {
		f := new(models.Flight)
		if err := rows.Scan(&f.ID, &f.FlightNumber, &f.Date, &f.FromAirportId, &f.ToAirportId, &f.Price); err != nil {
			return nil, fmt.Errorf("failed to execute the query: %w", err)
		}
		flights = append(flights, f)
	}

	defer rows.Close()

	return flights, nil
}

func (dbhand *DBHandler) GetFlightByNumber(flightNumber string) (*models.Flight, error) {
	var flight models.Flight
	row := dbhand.db.QueryRow(`SELECT * FROM flight WHERE flight_number = $1;`, flightNumber)
	err := row.Scan(
		&flight.ID,
		&flight.FlightNumber,
		&flight.Date,
		&flight.FromAirportId,
		&flight.ToAirportId,
		&flight.Price,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &flight, err
		}
	}

	return &flight, nil
}

func (dbhand *DBHandler) GetAirportByID(airportID string) (*models.Airport, error) {

	var airport models.Airport
	row := dbhand.db.QueryRow(`SELECT * FROM airport WHERE id = $1;`, airportID)
	err := row.Scan(
		&airport.ID,
		&airport.Name,
		&airport.City,
		&airport.Country,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &airport, err
		}
	}

	return &airport, nil
}
