package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"gustoliveira/speedtest-monitor/internal/database"
)

type Server struct {
	port int
	db   database.Service
}

func NewServer() (*http.Server, database.Service) {
	port, _ := strconv.Atoi(os.Getenv("PORT"))

	NewServer := &Server{
		port: port,
		db:   database.New(),
	}

	addr := fmt.Sprintf(":%d", NewServer.port)

	server := &http.Server{
		Addr:         addr,
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	fmt.Printf("Server started in address: %v\n", addr)

	return server, NewServer.db
}
