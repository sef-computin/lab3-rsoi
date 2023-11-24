package main

import (
	"database/sql"
	"fmt"
	"lab2/src/bonus-service/dbhandler"
	"lab2/src/bonus-service/handlers"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
)

func main() {
	dbURL := fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s sslmode=disable",
		"postgres", 5432, "postgres", "privileges", "postgres")

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	bonusHandler := &handlers.BonusHandler{
		DBHandler: dbhandler.InitDBHandler(db),
	}

	router := gin.Default()

	router.GET("/manage/health", bonusHandler.GetHealth)

	router.GET("/api/v1/bonus/:username", bonusHandler.GetPrivilegeByUsernameHandler)
	router.GET("/api/v1/bonus/history/:privilegeId", bonusHandler.GetHistoryByIdHandler)
	router.POST("/api/v1/bonus", bonusHandler.CreatePrivilegeHistoryHandler)
	router.POST("/api/v1/bonus/privilege", bonusHandler.CreatePrivilegeHandler)
	router.POST("/api/v1/bonus/:username", bonusHandler.UpdatePrivilegeHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8050"
	}

	log.Println("Server is listening on port: ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
