package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("true"))
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	rData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("error while reading request body %s while creating user.", r.URL)
		w.WriteHeader(http.StatusBadRequest)
	}

	log.Info(string(rData))
	defer r.Body.Close()
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
