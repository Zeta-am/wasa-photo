package database

import (
	"database/sql"
	"errors"

	"github.com/Zeta-am/wasa-photo/service/utils"
)

func (db *appdbimpl) GetUserByName(username string) (utils.User, error) {
	var user utils.User
	err := db.c.QueryRow("SELECT user_id, username FROM users WHERE username = ?;", username).Scan(&user.UserID, &user.Username)
	return user, err
}

func (db *appdbimpl) IsUsernameExists(username string) (bool, error) {
	var usern string
	err := db.c.QueryRow("SELECT username FROM users WHERE username = ?;", username).Scan(&usern)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}
	return usern != "", err
}


func (db *appdbimpl) CreateUser(u utils.User) (utils.User, error) {
	result, err := db.c.Exec("INSERT INTO users (username, user_name, user_surname) VALUES (?, ?, ?);", u.Username, u.Name, u.Surname)
	if err != nil {
		return utils.User{}, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return utils.User{}, err
	}
	u.UserID = int(id)
	return u, nil
}
