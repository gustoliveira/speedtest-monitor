package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"gustoliveira/speedtest-monitor/internal"

	_ "github.com/mattn/go-sqlite3"
)

type Service interface {
	Close() error
	InsertSpeedtest(internal.SpeedtestResponse) (string, error)
}

type service struct {
	db *sql.DB
}

var (
	dburl = os.Getenv("DB_URL")

	dbInstance *service
)

func (s *service) GetDB() *sql.DB {
	return s.db
}

func New() Service {
	if dbInstance != nil {
		return dbInstance
	}

	fmt.Println(dburl)

	db, err := sql.Open("sqlite3", dburl)
	if err != nil {
		log.Fatal(err)
	}

	err = startDatabase(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Database started: %s", dburl)

	dbInstance = &service{
		db: db,
	}

	return dbInstance
}

func (s *service) Close() error {
	log.Printf("Disconnected from database: %s", dburl)
	return s.db.Close()
}

func startDatabase(db *sql.DB) error {
	statement, err := createTable(db)
	if err != nil {
		return err
	}

	_, err = statement.Exec()
	return err
}

func createTable(db *sql.DB) (*sql.Stmt, error) {
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

	return db.Prepare(query)
}

func (s *service) InsertSpeedtest(entry internal.SpeedtestResponse) (string, error) {
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

	statement, err := s.db.Prepare(query)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = statement.Exec(
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

	return entry.Result.ID, err
}
