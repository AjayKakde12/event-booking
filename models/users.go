package models

import (
	"errors"
	"fmt"

	"event-booking.com/root/db"
	"event-booking.com/root/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES(?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return err
	}
	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return err
	}
	user.ID = userId
	return err
}

func (user *User) ValidateCredentials() error {
	var retrievedPassword string
	query := "SELECT password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)
	err := row.Scan(&retrievedPassword)
	if err != nil {
		return err
	}
	isValidaPassword := utils.CheckHashPassword(user.Password, retrievedPassword)
	if !isValidaPassword {
		return errors.New("credentials invalid")
	}
	return nil
}
