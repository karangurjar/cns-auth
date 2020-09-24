package handlers

import (
	"encoding/json"
	"fmt"
	"karsingh991/cns-auth/common"
	"karsingh991/cns-auth/modal"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

//HealthHandler will send a true message if app is reachable.
func HealthHandler(c echo.Context) error {
	return c.String(http.StatusOK, "True")
}

func createUserHandler(c echo.Context) error {
	//check if body is not empty
	if c.Request().Body == nil {
		errMsg := fmt.Sprint("request body is empty for request: %s", c.Request().RequestURI)
		log.Error(errMsg)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": errMsg,
		})
	}

	//read request data and bind to user
	var user modal.User
	if err := c.Bind(&user); err != nil {
		errMsg := fmt.Sprintf("error while readin request, err: %s", err.Error())
		log.Error(errMsg)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"error": errMsg,
		})
	}

	//convert plain password to hash before storing
	hashedPassword, err := common.HashPassword(user.Password)
	if err == nil {
		user.Password = hashedPassword
	}

	err = user.Create()
	if err != nil {
		errMsg := fmt.Sprintf("error while creatng new user: %s Error: %q", user.Name, err.Error())
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": errMsg,
		})
	}

	//finaly send success message
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "user created successfully",
		"error":   "",
		"user":    user,
	})
}

func getUserHandler(c echo.Context) {
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
