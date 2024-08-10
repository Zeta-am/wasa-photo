package database

import (
	"database/sql"
	"errors"

	"github.com/Zeta-am/wasa-photo/service/utils"
)

func (db *appdbimpl) GetUserByName(username string) (utils.User, error) {
	var user utils.User
	err := db.c.QueryRow(`SELECT user_id, username 
							FROM users 
							WHERE username = ?;`, username).Scan(&user.UserID, &user.Username)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (db *appdbimpl) IsUsernameExists(username string) (bool, error) {
	var usern string
	err := db.c.QueryRow(`SELECT username 
							FROM users 
							WHERE username = ?;`, username).Scan(&usern)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, nil	
		}
		return false, err
	}
	return true, nil
} 

func (db *appdbimpl) CreateUser(u utils.User) (utils.User, error) {
	result, err := db.c.Exec(`INSERT 
								INTO users (username, user_name, user_surname) 
								VALUES (?, ?, ?);`, u.Username, u.Name, u.Surname)
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

func (db *appdbimpl) GetUserProfile(userId int) (utils.User, error) {
	var user utils.User
	// Get the id, username, name and surname
	err := db.c.QueryRow(`SELECT user_id, username, user_name, user_surname 
							FROM users 
							WHERE user_id = ?;`, userId).Scan(&user.UserID, &user.Username, &user.Name, &user.Surname)
	if err != nil {
		return user, err		
	}

	// Get the number of followers
	err = db.c.QueryRow(`SELECT COUNT(follower_id)
							FROM follows
							WHERE followed_id = ?;`, userId).Scan(&user.FollowerCount)
	if err != nil {
		return user, err
	}

	// Get the number of followed
	err = db.c.QueryRow(`SELECT COUNT(followed_id)
							FROM follows
							WHERE follower_id = ?;`, userId).Scan(&user.FollowingCount)
	if err != nil {
		return user, err
	}
	// Get the number of post
	err = db.c.QueryRow(`SELECT COUNT(post_id)
							FROM posts
							WHERE user_id = ?;`, userId).Scan(&user.PostCount)
	if err != nil {
		return user, err
	}

	//TODO: Check if is banned

	return user, nil
}
