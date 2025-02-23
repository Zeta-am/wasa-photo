package database

import "github.com/Zeta-am/wasa-photo/service/utils"

func (db *appdbimpl) FollowUser(uid int, followedId int) (int, error) {
	var exists int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follows WHERE follower_id = ? AND followed_id = ?", uid, followedId).Scan(&exists)
	if err != nil {
		return ERROR, err
	}

	if exists > 0 {
		return UNIQUE_FAILED, err
	}

	_, err = db.c.Exec(`INSERT INTO follows (follower_id, followed_id) VALUES (?, ?)`,
		uid, followedId)

	if err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}

func (db *appdbimpl) UnfollowUser(uid int, unfollowedId int) (int, error) {
	var exists int
	err := db.c.QueryRow("SELECT COUNT(*) FROM follows WHERE follower_id = ? AND followed_id = ?", uid, unfollowedId).Scan(&exists) 
	if err != nil {
		return ERROR, err
	}

	if exists == 0 {
		return UNIQUE_FAILED, nil
	}

	result, err := db.c.Exec(`DELETE FROM follows WHERE follower_id = ? AND followed_id = ?`,
		uid, unfollowedId)

	if err != nil {
		return ERROR, err
	}

	rows, err := result.RowsAffected()
	if err != nil {
		return ERROR, err
	}
	if rows == 0 {
		return NO_ROWS, nil
	}

	return SUCCESS, nil
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
