package database

import (
	"database/sql"

	"github.com/Zeta-am/wasa-photo/service/utils"
)

var GET_USER_STREAM = `SELECT
							p.*, 
							u.username,
							COUNT(DISTINCT l.post_id) AS like_count,
							COUNT(DISTINCT c.comm_id) AS comment_count,
							MAX(l.user_id IS NOT NULL AND l.user_id = ?) AS liked	
						FROM posts p
						INNER JOIN follows f ON p.user_id = f.followed_id
						INNER JOIN users u ON p.user_id = u.user_id
						LEFT JOIN likes l ON p.post_id = l.post_id
						LEFT JOIN comments c ON p.post_id = c.post_id
						WHERE f.follower_id = ?
						GROUP BY p.post_id
						ORDER BY p.timestamp DESC;
`

func (db *appdbimpl) GetUserByName(username string) (utils.User, int, error) {
	var user utils.User
	err := db.c.QueryRow(`SELECT user_id, username 
							FROM users 
							WHERE username = ?;`, username).Scan(&user.UserID, &user.Username)
	res := checkResults(err)
	if res != SUCCESS {
		return user, res, err
	}

	return db.fillUser(user)
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
	return db.fillUser(user)
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

func (db *appdbimpl) fillUser(u utils.User) (utils.User, int, error) {
	// Fill FollowerCount field
	err := db.c.QueryRow(`SELECT COUNT(*)
						FROM follows
						WHERE followed_id = ?`, u.UserID).Scan(&u.FollowerCount)
	if err != nil {
		return utils.User{}, ERROR, err
	}

	// Fill FollowingCount field
	err = db.c.QueryRow(`SELECT COUNT(*)
						FROM follows
						WHERE follower_id = ?`, u.UserID).Scan(&u.FollowingCount)
	if err != nil {
		return utils.User{}, ERROR, err
	}

	// Fill PostCount field
	err = db.c.QueryRow(`SELECT COUNT(*)
						FROM posts
						WHERE user_id = ?`, u.UserID).Scan(&u.PostCount)
	if err != nil {
		return utils.User{}, ERROR, err
	}

	return u, SUCCESS, nil
}

func (db *appdbimpl) SetMyUsername(username string, uid int) (int, error) {
	_, err := db.c.Exec(`UPDATE users
							SET username = ?
							WHERE user_id = ?`, username, uid)
	if res := checkResults(err); res != SUCCESS {
		return res, err
	}
	return SUCCESS, nil
}

func (db *appdbimpl) GetMyStream(uid int) ([]utils.Post, int, error) {
	rows, err := db.c.Query(GET_USER_STREAM, uid, uid)
	if err != nil {
		return nil, ERROR, err
	}
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	var stream []utils.Post
	for rows.Next() {
		var post utils.Post
		var nullCapt sql.NullString
		if err := rows.Scan(&post.PostID, &post.UserID, &post.Image, &post.Timestamp, &nullCapt, &post.Username, &post.LikeCount, &post.CommentCount, &post.Liked); err != nil {
			return nil, ERROR, err
		}
		post.Caption = nullCapt.String
		stream = append(stream, post)
	}

	if err = rows.Err(); err != nil {
		return nil, ERROR, err
	}

	return stream, SUCCESS, nil
}

func (db *appdbimpl) GetUserPhotos(uid int) ([]utils.Post, int, error) {
	rows, err := db.c.Query(`SELECT p.*,
									COUNT(DISTINCT l.post_id) AS like_count,
									COUNT(DISTINCT c.comm_id) AS comment_count
							FROM posts p
							LEFT JOIN likes l ON p.post_id = l.post_id
							LEFT JOIN comments c ON p.post_id = c.post_id
							WHERE p.user_id = ?
							GROUP BY p.post_id
							ORDER BY p.timestamp DESC`, uid)
	if err != nil {
		return nil, ERROR, err
	}

	defer func() {
		if errow := rows.Close(); errow != nil {
			err = errow
		}
	}()

	var posts []utils.Post
	for rows.Next() {
		var post utils.Post
		if err = rows.Scan(&post.PostID, &post.UserID, &post.Image, &post.Timestamp, &post.Caption, &post.LikeCount, &post.CommentCount); err != nil {
			return nil, ERROR, err
		}
		posts = append(posts, post)
	}
	if err = rows.Err(); err != nil {
		return nil, ERROR, err
	}
	return posts, SUCCESS, nil
}
