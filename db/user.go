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

func GetUser(req models.GetUserRequest) (models.User, error) {

	getUserCmd := fmt.Sprintf("SELECT * from users where uid = '%s';", req.Uid)

	fmt.Println("query: ", getUserCmd)

	db := GetDb()

	resp := models.User{}

	err := db.QueryRow(getUserCmd).Scan(&resp.Id, &resp.UUID, &resp.Name, &resp.Email, &resp.Password, &resp.Profile_image_url, &resp.Profile_public_id, &resp.Active, &resp.Verified_at, &resp.Created_at, &resp.Updated_at, &resp.Deleted_at)
	if err != nil {
		return resp, err
	}

	return resp, nil
}

func DeleteUser(req models.DeleteUserRequest) error {

	deleteUserCmd := fmt.Sprintf("DELETE FROM users WHERE uid = '%s';", req.Uid)

	fmt.Println("query: ", deleteUserCmd)

	db := GetDb()

	_, err := db.Exec(deleteUserCmd)
	if err != nil {
		return err
	}

	return nil
}

func DeactivateUser(req models.DeactivateUserRequest) error {

	deactivateUserCmd := fmt.Sprintf("UPDATE users SET active = false WHERE uid = '%s';", req.Uid)

	fmt.Println("query: ", deactivateUserCmd)

	db := GetDb()

	_, err := db.Exec(deactivateUserCmd)
	if err != nil {
		return err
	}

	return nil
}
