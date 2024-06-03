package main

import (
	"fmt"

	"gustoliveira/speedtest-monitor/internal/server"
	"gustoliveira/speedtest-monitor/monitor"
)

func main() {
	server, dbService := server.NewServer()

	err := server.ListenAndServe()
	if err != nil {
		panic(fmt.Sprintf("Cannot start server: %s", err))
	}

	monitor.RunSpeedtestCronJob(dbService)
}
