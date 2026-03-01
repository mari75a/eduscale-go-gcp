package services

import (
	"errors"
	"golang.org/x/crypto/bcrypt"

	"eduscale/internal/models"
	"eduscale/internal/repository"
)

func RegisterUser(user models.User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)

	return repository.CreateUser(user)
}

func LoginUser(email string, password string) (models.User, error) {

	user, err := repository.GetUserByEmail(email)

	if err != nil {
		return user, err
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(password),
	)

	if err != nil {
		return user, errors.New("invalid password")
	}

	return user, nil

}

func GetAllUsers() ([]models.User, error) {

	return repository.GetUsers()

}