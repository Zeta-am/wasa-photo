package database

import (
	
	"github.com/Zeta-am/wasa-photo/service/utils"
)

// TODO: Rivedi tutta la struttura di aggiornamento dei dati e di controllo degli errori di tutti
// i metodi scritti		
// func (db *appdbimpl) fillUser(user utils.User) (utils.User, int, error) 

func (db *appdbimpl) GetUserByName(username string) (utils.User, int, error) {
	var user utils.User
	err := db.c.QueryRow(`SELECT user_id, username 
							FROM users 
							WHERE username = ?;`, username).Scan(&user.UserID, &user.Username)
	res := checkResults(err)
	if res != SUCCESS {
		return user, res, err
	}

	return user, res, nil
}

func (db *appdbimpl) GetUserById(id int) (utils.User, int, error) {
	var user utils.User

	err := db.c.QueryRow(`SELECT user_id, username 
							FROM users
							WHERE user_id = ?`, id).Scan(&user.UserID, &user.Username)
	res := checkResults(err)
	if res != SUCCESS {
		return utils.User{}, res, err
	}			
	return user, res, nil
}

func (db *appdbimpl) IsUsernameExists(username string) (bool, int, error) {
	var usern string
	err := db.c.QueryRow(`SELECT username 
							FROM users 
							WHERE username = ?;`, username).Scan(&usern)
	res := checkResults(err)
	if res != SUCCESS {
		if res == NO_ROWS {
			return false, res, nil
		}
		return false, res, err
	}
	return true, res, nil
} 

func (db *appdbimpl) CreateUser(u utils.User) (utils.User, int, error) {
	_, err := db.c.Exec(`INSERT 
								INTO users (username, user_name, user_surname) 
								VALUES (?, ?, ?);`, u.Username, u.Name, u.Surname)
	res := checkResults(err)
	if res != SUCCESS {
		return utils.User{}, res, err
	}
	return u, res, nil
}

func (db *appdbimpl) GetUserProfile(userId int) (utils.User, int, error) {
	var user utils.User
	// Get the id, username, name and surname
	err := db.c.QueryRow(`SELECT user_id, username, user_name, user_surname 
							FROM users 
							WHERE user_id = ?;`, userId).Scan(&user.UserID, &user.Username, &user.Name, &user.Surname)
	if res := checkResults(err); res != SUCCESS {
		return user, res, err
	}

	// Get the number of followers
	err = db.c.QueryRow(`SELECT COUNT(follower_id)
							FROM follows
							WHERE followed_id = ?;`, userId).Scan(&user.FollowerCount)
	if res := checkResults(err); res != SUCCESS {
		return user, res, err
	}

	// Get the number of followed
	err = db.c.QueryRow(`SELECT COUNT(followed_id)
							FROM follows
							WHERE follower_id = ?;`, userId).Scan(&user.FollowingCount)
	if res := checkResults(err); res != SUCCESS {
		return user, res, err
	}
	// Get the number of post
	err = db.c.QueryRow(`SELECT COUNT(post_id)
							FROM posts
							WHERE user_id = ?;`, userId).Scan(&user.PostCount)
	res := checkResults(err)
	if res != SUCCESS {
		return user, res, err
	}

	//TODO: Check if is banned
	return user, res, nil
}
