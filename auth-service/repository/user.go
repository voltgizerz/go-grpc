package repository

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/auth-service/models"
	"github.com/auth-service/utils"
)

// UserRepository - interface for user repository.
type UserRepository interface {
	GetUserByID(id int64) (*models.User, error)
	GetUserByUsernameAndPassword(username string, password string) (*models.User, error)
}

// ReadDataUsers - read data from file bcz we are not using db conn.
func ReadDataUsers() []models.User {
	// temporary database using json
	// golang open file
	file, err := os.Open("data/users.json")
	if err != nil {
		log.Println(err)
	}

	// golang read file
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println(err)
	}

	// golang unmarshal json
	var users []models.User
	err = json.Unmarshal(data, &users)
	if err != nil {
		log.Println(err)
	}

	return users
}

// NewUserRepository - create new user repository.
func NewUserRepository() UserRepository {
	return UserRepositoryImpl{
		DataUsers: ReadDataUsers(),
	}
}

// UserRepositoryImpl - implementation of user repository.
type UserRepositoryImpl struct {
	DataUsers []models.User
}

// GetUserByID - get user by id.
func (u UserRepositoryImpl) GetUserByID(id int64) (*models.User, error) {
	for _, v := range u.DataUsers {
		if v.ID == id {
			return &v, nil
		}
	}

	return nil, nil
}

// GetUserByUsernameAndPassword - get user by username and password.
func (u UserRepositoryImpl) GetUserByUsernameAndPassword(username string, password string) (*models.User, error) {
	for _, v := range u.DataUsers {
		if v.Username == username && v.Password == utils.HashPassword(password) {
			return &v, nil
		}
	}

	return nil, nil
}

