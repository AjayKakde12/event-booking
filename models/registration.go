package models

import (
	"errors"
	"fmt"
	"strconv"

	"event-booking.com/root/db"
)

type Registration struct {
	ID      string
	UserId  int64 `binding:"required"`
	EventId int64 `binding:"required"`
}

func (r *Registration) RegisterForEvent() (int64, error) {
	query := `INSERT INTO registrations(event_id, user_id) VALUES(?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	registration, err := stmt.Exec(r.EventId, r.UserId)
	if err != nil {
		return 0, err
	}
	registrationId, err := registration.LastInsertId()
	if err != nil {
		return 0, err
	}
	return registrationId, nil
}

func (r *Registration) GetRegistration() (int64, error) {
	fmt.Println(r.EventId, r.UserId)
	query := `SELECT id FROM registrations WHERE user_id = ? AND event_id = ?`
	row := db.DB.QueryRow(query, r.UserId, r.EventId)
	err := row.Scan(&r.ID)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}
	id, err := strconv.ParseInt(r.ID, 10, 64)
	if err != nil {
		fmt.Println(err)
		return 0, errors.New("cannot parse id")
	}
	return id, nil
}

func CancelRegistration(id int64) error {
	query := `DELETE FROM registrations WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(id)

	return err
}
