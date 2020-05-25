package main

import (
	"fmt"
	"io/ioutil"
	"karsingh991/cns-auth/modal"
	"net/http"
	"os"

	"github.com/segmentio/encoding/json"
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
	if err != nil {
		log.Errorf("error in reading request %s body while creating user.", r.URL)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

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
		log.Errorf("error while creatng new user: %s", user.Name)
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("user created!"))
}

func main() {
	port := "8080"
	address := fmt.Sprintf(":%s", port)

	http.HandleFunc("/", healthHandler)
	http.HandleFunc("/user/create", createUserHandler)
	if err := http.ListenAndServe(address, nil); err != nil {
		log.Errorf("failed starting server on adress : %s", address)
		os.Exit(0)
	}

	log.Infof("Server started on address %s", address)
}
