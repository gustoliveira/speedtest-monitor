package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron"
)

type SpeedtestResponse struct {
	ID           int64
	DownloadBits float64   `json:"download"`
	UploadBits   float64   `json:"upload"`
	Ping         float32   `json:"ping"`
	Timestamp    time.Time `json:"timestamp"`
	Client       struct {
		IP  string `json:"ip"`
		ISP string `json:"isp"`
	}
	Server struct {
		Sponsor string `json:"sponsor"`
		Country string `json:"country"`
		Name    string `json:"name"`
	}
}

func main() {
	db, err := sql.Open("sqlite3", "./database/speedtest.db")
	if err != nil {
		log.Fatalln(err)
	}
	defer db.Close()

	createTable(db)

	startCronJob(func() {
		body := getSpeedtest()

		var response SpeedtestResponse
		err := json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalln(err)
		}

		insertSpeedtest(db, response)
	})

	select {}
}

func startCronJob(callback func()) {
	timezone := os.Getenv("TIMEZONE")
	if timezone == "" {
		log.Fatalln("TIMEZONE is not set\n")
	}

	loc, err := time.LoadLocation(timezone)
	if err != nil {
		log.Fatalln("Could not get location from TIMEZONE", timezone, "=", err)
	}

	test_period_string := os.Getenv("TEST_PERIOD_MIN")
	if test_period_string == "" {
		log.Fatalln("TEST_PERIOD_MIN is not set")
	}

	test_period, err := strconv.Atoi(test_period_string)
	if err != nil {
		log.Fatalln("Error converting TEST_PERIOD_MIN =", err)
	}

	seconds := test_period * 60
	seconds = 30
	spec := fmt.Sprintf("@every %vs", seconds)

	cronJob := cron.NewWithLocation(loc)

	cronJob.AddFunc(spec, func() {
		now := time.Now()
		fmt.Printf("Running cron job %v\n", now)
		callback()
	})

	cronJob.Start()
}

func getSpeedtest() []byte {
	body, err := exec.Command("speedtest-cli", "--json").Output()
	if err != nil {
		log.Fatalln("Could not run speedtest-cli =", err)
	}
	return body
}

func createTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS speedtest (
			id INTEGER PRIMARY KEY,
			downloadBits FLOAT,
			uploadBits FLOAT,
			ping FLOAT,
			timestamp TIMESTAMP,
			ip_client TEXT,
			isp_clinet TEXT,
			sponsor_server TEXT,
			country_server TEXT,
			name_server TEXT
		)
	`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	statement.Exec()
}

func insertSpeedtest(db *sql.DB, entry SpeedtestResponse) {
	query := `
		INSERT INTO speedtest (
			downloadBits,
			uploadBits,
			ping,
			timestamp,
			ip_client,
			isp_clinet,
			sponsor_server,
			country_server,
			name_server
		) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ? )
	`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	statement.Exec(
		entry.DownloadBits,
		entry.UploadBits,
		entry.Ping,
		entry.Timestamp,
		entry.Client.IP,
		entry.Client.ISP,
		entry.Server.Sponsor,
		entry.Server.Country,
		entry.Server.Name,
	)
}
