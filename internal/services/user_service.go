package services

import (
	"errors"

	"eduscale/internal/models"
	"eduscale/internal/repository"
)

func RegisterUser(user models.User) error {

	return repository.CreateUser(user)

}

func LoginUser(email string, password string) (models.User, error) {

	user, err := repository.GetUserByEmail(email)

	if err != nil {
		return user, err
	}

	if user.Password != password {
		return user, errors.New("invalid password")
	}

	return user, nil

}

func GetAllUsers() ([]models.User, error) {

	return repository.GetUsers()

}