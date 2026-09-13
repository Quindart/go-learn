package models

import (
	"fmt"

	"example.com/app-event-booking/db"
	"example.com/app-event-booking/utils"
)

type User struct {
	ID       int64
	Email    string `binding: required`
	Password string `binding: required`
}

func (u *User) Save() error {
	query := "insert into users(email, password) values (? , ? )"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	passhash, _ := utils.HashPassword(u.Password)
	fmt.Println(passhash)

	u.Password = passhash
	result, err := stmt.Exec(
		u.Email,
		u.Password,
	)
	if err != nil {
		return err
	}
	id, err := result.LastInsertId()
	u.ID = id
	if err != nil {
		return err
	}
	return nil
}

func ValidateCredentials(user *User) (*User, bool) {
	query := `select * from users where email = ?`
	var userDB User
	result := db.DB.QueryRow(query, user.Email)
	err := result.Scan(
		&userDB.ID,
		&userDB.Email,
		&userDB.Password,
	)
	fmt.Println("result:::", userDB.Password)
	fmt.Println("result:::", user.Password)

	if err != nil {
		return nil, false
	}
	checked := utils.CheckPasswordHash(user.Password, userDB.Password)
	return &userDB, checked
}
