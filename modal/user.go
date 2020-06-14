package modal

import (
	"fmt"
	"karsingh991/cns-auth/db"

	log "github.com/sirupsen/logrus"
)

type User struct {
	Name         string `json:"username"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	MobileNumber int    `json:"mobilenumber"`
}

func (u User) Create() error {
	//insert into db
	log.Infof("User: %v is inserted", u)
	sql := fmt.Sprintf("INSERT INTO users VALUES ('%s', '%s', '%s')", u.Name, u.Email, u.Password)
	err := db.Insert(sql)
	if err != nil {
		log.Errorf("error while inserting user recode to db %q", err.Error())
		return err
	}
	return nil
}

func GetUsers() ([]User, error) {
	var users []User
	db := db.GetCnsDB()
	sql := `SELECT * FROM USERS`
	rows, err := db.Query(sql)
	if err != nil {
		return users, err
	}

	for rows.Next() {
		var name, email, password string
		err := rows.Scan(&name, &email, &password)
		if err != nil {
			return users, err
		}

		users = append(users, User{
			Name:     name,
			Email:    email,
			Password: password,
		})
	}
	log.Info("return user: ", users)
	return users, nil
}
