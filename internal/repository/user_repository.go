package repository

import (
	"eduscale/configs"
	"eduscale/internal/models"
)

func CreateUser(user models.User) error {

	query := `
	INSERT INTO users(name,email,password,role)
	VALUES($1,$2,$3,$4)
	`

	_, err := configs.DB.Exec(
		query,
		user.Name,
		user.Email,
		user.Password,
		user.Role,
	)

	return err
}

func GetUsers() ([]models.User, error) {

	rows, err := configs.DB.Query(
		"SELECT id,name,email,role FROM users",
	)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {

		var user models.User

		rows.Scan(
			&user.ID,
			&user.Name,
			&user.Email,
			&user.Role,
		)

		users = append(users, user)
	}

	return users, nil
}

func GetUserByEmail(email string) (models.User, error) {

	var user models.User

	query := `
	SELECT id,name,email,password,role
	FROM users
	WHERE email=$1
	`

	err := configs.DB.QueryRow(query, email).Scan(
		&user.ID,
		&user.Name,
		&user.Email,
		&user.Password,
		&user.Role,
	)

	return user, err
}