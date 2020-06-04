package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"karsingh991/cns-auth/db"
	"karsingh991/cns-auth/modal"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("true"))
	if err != nil {
		log.Error("error while sending reponse.")
	}
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	rData, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()

	if err != nil {
		log.Errorf("error in reading request %s body while creating user.", r.URL)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user modal.User
	//unmarshal request data
	err = json.Unmarshal(rData, &user)
	if err != nil {
		log.Errorf("unmarshaling error while inserting user")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request!"))
		return
	}

	err = user.Create()
	if err != nil {
		log.Errorf("error while creatng new user: %s error: %q", user.Name, err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request!"))
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created!"))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	users, err := modal.GetUsers()
	if err != nil {
		log.Errorf("error while getting users details from db Error %q", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request!"))
		return
	}

	jsonData, err := json.Marshal(users)
	if err != nil {
		log.Errorf("error while marshling users Error %q", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Bad request!"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func initRestAPIs() {
	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/user/create", createUserHandler)
	http.HandleFunc("/user", getUserHandler)
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

	log.Info("Db connection stablished, starting server...")

	initRestAPIs()

	serverPort := "8080"
	address := fmt.Sprintf(":%s", serverPort)

	if err := http.ListenAndServe(address, nil); err != nil {
		log.Errorf("failed starting server on adress : %s", address)
		os.Exit(0)
	}

	log.Infof("Server started on address %s", address)
}
