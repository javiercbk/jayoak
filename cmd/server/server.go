package main

//go:generate sqlboiler psql

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/javiercbk/jayoak/http"
)

const defaultFilesFolder = "jayoak-files"
const defaultLogFilePath = "jayoak-server.log"
const defaultAddress = "0.0.0.0"
const defaultDBName = "jayoak"
const defaultDBUser = "jayoak"

func main() {
	var logFilePath, filesFolder, address, dbName, dbHost, dbUser, dbPass, rAddress, rPass, rSecret string
	flag.StringVar(&logFilePath, "l", defaultLogFilePath, "the log file location")
	flag.StringVar(&filesFolder, "f", defaultFilesFolder, "the path to the application uploads folder")
	flag.StringVar(&address, "a", defaultAddress, "the http server address")
	flag.StringVar(&dbName, "dbn", defaultDBName, "the database name")
	flag.StringVar(&dbHost, "dbh", defaultDBUser, "the database host")
	flag.StringVar(&dbUser, "dbu", "", "the database user")
	flag.StringVar(&dbPass, "dbp", "", "the database password")
	flag.StringVar(&rAddress, "ra", "", "the redis address")
	flag.StringVar(&rPass, "rp", "", "the redis password")
	flag.StringVar(&rSecret, "rs", "", "the redis secret")
	flag.Parse()
	logFile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		fmt.Printf("error opening lof file: %s", err)
		os.Exit(1)
	}
	defer logFile.Close()
	logger := log.New(logFile, "applog: ", log.Lshortfile|log.LstdFlags)
	cnf := http.Config{
		Address:       address,
		FilesFolder:   filesFolder,
		DBName:        dbName,
		DBHost:        dbHost,
		DBUser:        dbUser,
		DBPass:        dbPass,
		RedisAddress:  rAddress,
		RedisPassword: rPass,
		RedisSecret:   rSecret,
	}
	err = http.Serve(cnf, logger)
	if err != nil {
		logger.Fatalf("could not start server %s\n", err)
	}
}
