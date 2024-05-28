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
	ID         int64
	Timestamp  time.Time `json:"timestamp"`
	ISP        string    `json:"isp"`
	PackatLoss string    `json:"packatLoss"`
	Ping       struct {
		Latency float32 `json:"latency"`
		Jitter  float32 `json:"jitter"`
	}
	Interface struct {
		ExternalIp  string `json:"externalIp"`
		ContainerIp string `json:"internalIp"`
		IsVPN       bool   `json:"isVpn"`
	}
	Server struct {
		Name     string `json:"name"`
		Location string `json:"location"`
		Country  string `json:"country"`
	}
	Download struct {
		Bandwidth float64 `json:"bandwidth"`
		Bytes     float64 `json:"bytes"`
		Latency   struct {
			Ping   float32 `json:"iqm"`
			Jitter float32 `json:"jitter"`
		}
	}
	Upload struct {
		Bandwidth float64 `json:"bandwidth"`
		Bytes     float64 `json:"bytes"`
		Latency   struct {
			Ping   float32 `json:"iqm"`
			Jitter float32 `json:"jitter"`
		}
	}
	Result struct {
		SpeedtestResponseUrl string `json:"url"`
		ID                   string `json:"id"`
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
		log.Fatalln("TIMEZONE is not set")
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
	max_attempts := 5

	var body []byte

	for ok := true; ok; ok = (max_attempts != 0) {
		fmt.Println("Trying to run speedtest-cli command...")

		response, err := exec.Command("speedtest", "--accept-license", "--accept-gdpr", "--format=json").Output()
		if err != nil {
			log.Panicf("Could not run speedtest = %v... %v attempts left\n", err, max_attempts)
			max_attempts--
			continue
		}

		body = response
		max_attempts = 0
	}

	return body
}

func createTable(db *sql.DB) {
	query := `
		CREATE TABLE IF NOT EXISTS speedtest (
			id INTEGER PRIMARY KEY,
			timestamp TIMESTAMP,
			isp TEXT,
			ping_latency FLOAT,
			ping_jitter FLOAT,
			is_vpn INTEGER,
			container_ip STRING,
			external_ip STRING,
			name_server STRING,
			location_server STRING,
			country_server STRING,
			bandwidth_download FLOAT,
			bytes_download FLOAT,
			ping_download FLOAT,
			jitter_download FLOAT,
			bandwidth_upload FLOAT,
			bytes_upload FLOAT,
			ping_upload FLOAT,
			jitter_upload FLOAT,
			speedtest_response_url TEXT,
			speedtest_id TEXT
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
			timestamp,
			isp,
			ping_latency,
			ping_jitter,
			is_vpn,
			container_ip,
			external_ip,
			name_server,
			location_server,
			country_server,
			bandwidth_download,
			bytes_download,
			ping_download,
			jitter_download,
			bandwidth_upload,
			bytes_upload,
			ping_upload,
			jitter_upload,
			speedtest_response_url,
			speedtest_id
		) VALUES ( ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? )
	`

	statement, err := db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	statement.Exec(
		entry.Timestamp,
		entry.ISP,
		entry.Ping.Latency,
		entry.Ping.Jitter,
		entry.Interface.IsVPN,
		entry.Interface.ContainerIp,
		entry.Interface.ExternalIp,
		entry.Server.Name,
		entry.Server.Location,
		entry.Server.Country,
		entry.Download.Bandwidth,
		entry.Download.Bytes,
		entry.Download.Latency.Ping,
		entry.Download.Latency.Jitter,
		entry.Upload.Bandwidth,
		entry.Upload.Bytes,
		entry.Upload.Latency.Ping,
		entry.Upload.Latency.Jitter,
		entry.Result.SpeedtestResponseUrl,
		entry.Result.ID,
	)
}
