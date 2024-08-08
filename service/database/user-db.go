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
	var user utils.User
	user.Username = u.Username
	err := db.c.QueryRow("INSERT INTO users (user_id, username) VALUES (?, ?);").Scan(&user.UserID)
	return user, err
}