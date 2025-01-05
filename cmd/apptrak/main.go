package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/JalalA984/apptrak/internal/models"
	"github.com/JalalA984/apptrak/pkg/config"
	"github.com/joho/godotenv"

	_ "github.com/go-sql-driver/mysql" // we need the driver’s init() function to run so that it can register itself with the database/sql package
)

// Define a wrapper around config.Application
type applicationConf struct {
	*config.ApplicationConfig
}

func main() {
	godotenv.Load(".env")

	port := flag.String("port", ":5000", "HTTP Network Address")
	mysqlDSN := os.Getenv("DSN")
	dsn := flag.String("dsn", mysqlDSN, "MySQL data source name")

	flag.Parse()

	infoLog := log.New(os.Stdout, "[INFO]\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "[ERROR]\t", log.Ldate|log.Ltime|log.Llongfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}

	// Make sure connection pool is closed before main terminates
	defer db.Close()

	appConfig := &applicationConf{
		ApplicationConfig: &config.ApplicationConfig{
			ErrorLog: errorLog,
			InfoLog: infoLog,
			Applications: &models.ApplicationModel{DB: db},
		},
	}

	server := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  appConfig.routes(),
	}

	infoLog.Printf("Application starting on %s", *port)
	err = server.ListenAndServe()
	errorLog.Fatal(err)

}

// The openDB() function wraps sql.Open() and returns a sql.DB connection pool for a given DSN.
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
