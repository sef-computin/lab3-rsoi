package dbhandler

import (
	"database/sql"
	"fmt"

	"lab2/src/ticket-service/models"
)

type TicketDB interface {
	GetTicketsByUsername(username string) ([]*models.Ticket, error)
	CreateTicket(ticket *models.Ticket) error
	DeleteTicketByUID(uid string) error
}

type DBHandler struct {
	db *sql.DB
}

func InitDBHandler(db *sql.DB) *DBHandler {
	return &DBHandler{
		db: db,
	}
}

func (r *DBHandler) GetTicketsByUsername(username string) ([]*models.Ticket, error) {

	var tickets []*models.Ticket
	rows, err := r.db.Query(`SELECT * FROM ticket WHERE username = $1;`, username)
	if err != nil {
		return nil, fmt.Errorf("failed to execute the query: %w", err)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to execute the query: %s", err)
	}

	for rows.Next() {
		ticket := new(models.Ticket)
		rows.Scan(
			&ticket.ID,
			&ticket.TicketUID,
			&ticket.Username,
			&ticket.FlightNumber,
			&ticket.Price,
			&ticket.Status)

		if err != nil {
			return nil, fmt.Errorf("failed to execute the query: %s", err)
		}

		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (r *DBHandler) CreateTicket(ticket *models.Ticket) error {

	_, err := r.db.Query(
		`INSERT INTO ticket (ticket_uid, username, flight_number, price, status) VALUES ($1, $2, $3, $4, $5) RETURNING id;`,
		ticket.TicketUID,
		ticket.Username,
		ticket.FlightNumber,
		ticket.Price,
		ticket.Status,
	)

	return err
}

func (r *DBHandler) DeleteTicketByUID(uid string) error {

	_, err := r.db.Query(`DELETE FROM ticket WHERE ticket_uid=$1`, uid)

	return err
}
