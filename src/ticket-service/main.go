package main

import (
	"database/sql"
	"fmt"
	"lab2/src/ticket-service/dbhandler"
	"lab2/src/ticket-service/handlers"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"postgres", 5432, "postgres", "tickets", "postgres")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	ticketHandler := &handlers.TicketHandler{
		DBHandler: dbhandler.InitDBHandler(db),
	}

	router := gin.Default()

	router.GET("/manage/health", ticketHandler.GetHealth)

	router.POST("/api/v1/tickets", ticketHandler.BuyTicketHandler)
	router.GET("/api/v1/tickets/:username", ticketHandler.GetTicketsByUsernameHandler)
	router.DELETE("/api/v1/tickets/:uid", ticketHandler.DeleteTicket)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8070"
	}

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
