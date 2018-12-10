package main

import (
	"flag"
	"log"

	"github.com/therusetiawan/golang-simple-api/internal/api"
	"github.com/therusetiawan/golang-simple-api/pkg/db"
)

var (
	listenPort = flag.String("listen-port", "9000", "Port where app listen to")
	dbUrl      = flag.String("db-url", "postgres://postgres:postgres@localhost:5432/golang-simple-crud?sslmode=disable", "Connection string to postgres")
	debug      = flag.Bool("debug", true, "Want to verbose query or not")
)

func main() {
	flag.Parse()

	// database configurations
	dbConfig := db.Config{
		URL:   *dbUrl,
		Debug: *debug,
	}

	// database connection
	dbConn, err := db.NewConnection(&dbConfig)
	if err != nil {
		log.Fatalf("database connection failed")
	}
	defer dbConn.Close()

	// run migrations
	err = db.Migrate(*dbUrl)
	if err != nil {
		log.Fatalf(err.Error())
	}

	// server configurations
	apiConfig := api.Config{
		ListenPort: *listenPort,
	}

	apiConfig.Start() // start server
}
