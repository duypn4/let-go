package models

import (
	"errors"
	"eventsapi/db"
	"eventsapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(email, password)
	VALUES(?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}

	user.ID, err = result.LastInsertId()
	return err
}

func (user User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)

	var retriedPassword string
	err := row.Scan(&user.ID, &retriedPassword)
	if err != nil {
		return errors.New("could not authenticate")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retriedPassword)
	if !passwordIsValid {
		return errors.New("could not authenticate")
	}

	return nil
}
