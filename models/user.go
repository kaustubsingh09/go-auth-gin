package models

import (
	"errors"

	"github.com/kaustubsingh09/go-auth-gin/db"
	"github.com/kaustubsingh09/go-auth-gin/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users (email, password)
	VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	encryptedPassword, err := utils.HashPassword(&u.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(&u.Email, encryptedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	if err != nil {
		return err
	}

	u.ID = id
	return nil
}

func GetAllUsers() ([]User, error) {
	query := `SELECT * FROM users`

	result, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer result.Close()

	var users []User

	for result.Next() {
		var user User
		err = result.Scan(&user.ID, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil

}

func (u *User) LoginValidationCred() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, &u.Email)

	var retrievedPassword string

	//to get the password from scan method
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}
	passwordIsValid := utils.DecryptPassword(&u.Password, &retrievedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentials")
	}
	return nil
}
