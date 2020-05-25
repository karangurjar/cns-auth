package modal

import log "github.com/sirupsen/logrus"

type User struct {
	Name     string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u User) Create() error {
	//insert into db
	log.Infof("User: %v is inserted", u)
	return nil
}
