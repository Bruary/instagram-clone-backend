package db

import (
	"fmt"

	"github.com/Bruary/instagram-clone-backend/models"
	"github.com/google/uuid"
)

func CreateNewUser(req models.NewUser) error {

	newUser := models.User{
		UUID:     uuid.NewString(),
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}

	insertNewUserCmd := fmt.Sprintf("INSERT INTO users(uid, name, email, password) VALUES('%s', '%s', '%s', '%s');", newUser.UUID, newUser.Name, newUser.Email, newUser.Password)

	fmt.Println("query: ", insertNewUserCmd)

	db := GetDb()

	_, err := db.Exec(insertNewUserCmd)
	if err != nil {
		return err
	}

	return nil
}
