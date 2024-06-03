package monitor

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
	"time"

	"gustoliveira/speedtest-monitor/internal"
	"gustoliveira/speedtest-monitor/internal/database"

	_ "github.com/mattn/go-sqlite3"
	"github.com/robfig/cron"
)

func RunSpeedtestCronJob(dbService database.Service) {
	startCronJob(func() {
		newSpeedtest := getSpeedtest()

		fmt.Println(string(newSpeedtest))

		body := getSpeedtest()

		var response internal.SpeedtestResponse
		err := json.Unmarshal(body, &response)
		if err != nil {
			log.Fatalln(err)
		}

		dbService.InsertSpeedtest(response)
		if err != nil {
			log.Fatalln(err)
		}
	})
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
		fmt.Println("Speedtest command ran successfully!")
	}

	return body
}
