package database

import "github.com/Zeta-am/wasa-photo/service/utils"


func (db *appdbimpl) FollowUser(uid int, followedId int) (int, error) {
	_, err := db.c.Exec(`INSERT
							INTO follows (follower_id, followed_id)
							VALUES (?, ?)`, uid, followedId)
	res := checkResults(err)
	if res != SUCCESS {
		return res, err
	} 
	return res, nil
}

func (db *appdbimpl) UnfollowUser(uid int, unfollowedId int) (int, error) {
	_, err := db.c.Exec(`DELETE
							FROM follows 
							WHERE follower_id = ? AND followed_id = ?`, uid, unfollowedId)
	res := checkResults(err)
	if res != SUCCESS {
		return res, err
	} 
return res, nil 			
}

func (db *appdbimpl) GetListFollowers(uid int) ([]utils.User, int, error) {
	rows, err := db.c.Query(`SELECT users.user_id, users.username
								FROM users
								INNER JOIN follows ON users.user_id = follows.follower_id
								WHERE follows.followed_id = ?`, uid)
	if err != nil {
		return nil, ERROR, err
	}

	return getUsers(rows)
}

func (db *appdbimpl) GetListFollowings(uid int) ([]utils.User, int, error) {
	rows, err := db.c.Query(`SELECT users.user_id, users.username
								FROM users
								INNER JOIN follows ON users.user_id = follows.followed_id
								WHERE follows.follower_id = ?`, uid)
	if err != nil {
		return nil, ERROR, err
	}

	return getUsers(rows)
}