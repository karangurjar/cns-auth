package main

import (
	"fmt"
	"karsingh991/cns-auth/db"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func initRestAPIs() {
}

func main() {

	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "password"
		dbname   = "cns"
	)

	dbDriverName := "postgres"
	connectionStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	err := db.InitDB(dbDriverName, connectionStr)
	if err != nil {
		log.Errorf("error while connecting to db %q", err.Error())
		os.Exit(0)
	}

	log.Info("CNS conneted to db. Starting app ...")

	//register all the apis with the endpoint and handlers
	registerRestApis()

	serverPort := "8080"
	address := fmt.Sprintf(":%s", serverPort)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Errorf("failed starting server on adress : %s", address)
		os.Exit(0)
	}

	log.Infof("Server started on address %s", address)
	log.Info("end of cns-auth server")
}
