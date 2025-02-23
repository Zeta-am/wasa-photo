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
	rows, err := db.c.Query(`SELECT u.user_id, u.username,
								(SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = ? AND followed_id = u.user_id)) AS followed
								FROM users u
								INNER JOIN follows f ON u.user_id = f.follower_id
								WHERE f.followed_id = ?`, uid, uid)
	if err != nil {
		return nil, ERROR, err
	}

	defer rows.Close()

	var users []utils.User
	for rows.Next() {
		var user utils.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Followed)
		if err != nil {
			return nil, ERROR, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, ERROR, err
	}

	return users, SUCCESS, nil
}

func (db *appdbimpl) GetListFollowings(uid int) ([]utils.User, int, error) {
	rows, err := db.c.Query(`SELECT u.user_id, u.username,
								(SELECT EXISTS(SELECT 1 FROM follows WHERE follower_id = ? AND followed_id = u.user_id)) AS followed
								FROM users u
								INNER JOIN follows f ON u.user_id = f.followed_id
								WHERE f.follower_id = ?`, uid, uid)
	if err != nil {
		return nil, ERROR, err
	}

	defer rows.Close()

	var users []utils.User
	for rows.Next() {
		var user utils.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Followed)
		if err != nil {
			return nil, ERROR, err
		}
		users = append(users, user)
	}

	if err = rows.Err(); err != nil {
		return nil, ERROR, err
	}

	return users, SUCCESS, nil
}
