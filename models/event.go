package models

import (
	"time"

	"github.com/kaustubsingh09/go-auth-gin/db"
)

type Event struct {
	ID          int64
	Name        string
	Description string
	Location    string
	DateTime    time.Time
	UserId      int64
}

func (e *Event) Save() error {
	//add it to a db
	query := `INSERT INTO events(name, description, location, dateTime, user_id)
	VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	e.ID = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := `
	SELECT * FROM events 
	`
	result, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer result.Close()

	var events []Event

	for result.Next() {
		var event Event
		err := result.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}

	return events, nil

}

func GetUniqueEvent(id int64) (*Event, error) {
	query := `SELECT * FROM events WHERE id = ?`
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (e *Event) UpdateEvent() (*Event, error) {
	query := `UPDATE events SET name = ?, description = ?, location = ?, dateTime = ?, user_id = ? WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// var event Event
	_, err = stmt.Exec(&e.Name, &e.Description, &e.Location, &e.DateTime, &e.UserId, &e.ID)

	if err != nil {
		return nil, err
	}
	return e, nil
}

func (e *Event) DeleteEvent() error {
	query := `DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(&e.ID)

	if err != nil {
		return err
	}
	return nil

}
