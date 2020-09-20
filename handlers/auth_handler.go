package handlers

import (
	"encoding/json"
	"io/ioutil"
	"karsingh991/cns-auth/common"
	"karsingh991/cns-auth/modal"
	"net/http"

	"github.com/labstack/echo/v4"
	"google.golang.org/appengine/log"
)

//HealthHandler will send a true message if app is reachable.
func HealthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "True")
}

func createUserHandler(c echo.Context) {
	body := c.Request().Body

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

	//convert plain password to hash before storing
	hashedPassword, err := common.HashPassword(user.Password)
	if err == nil {
		user.Password = hashedPassword
	}

	err = user.Create()
	if err != nil {
		log.Errorf("error while creatng new user: %s Error: %q", user.Name, err.Error())
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
