package database

import "github.com/Zeta-am/wasa-photo/service/utils"

func (db *appdbimpl) LikePhoto(uid int, pid int) (utils.Like, int, error) {
	var like utils.Like
	_, err := db.c.Exec(`INSERT
							INTO likes(user_id, post_id)
							VALUES (?, ?)`, uid, pid)
	if res := checkResults(err); res != SUCCESS {
		return like, res, err
	}
	like.PostID = pid
	like.UserID = uid
	return like, SUCCESS, nil
}

func (db *appdbimpl) UnlikePhoto(uid int, pid int) (int, error) {
	_, err := db.c.Exec(`DELETE
							FROM likes
							WHERE user_id = ? AND post_id = ?`, uid, pid)
	if res := checkResults(err); res != SUCCESS {
		return res, err
	}
	return SUCCESS, nil
}

func (db *appdbimpl) GetLikes(uid int, pid int) ([]utils.Like, int, error) {
	rows, err := db.c.Query(`SELECT * 
								FROM likes
								WHERE user_id = ? AND post_id = ?`, uid, pid)
	if res := checkResults(err); res != SUCCESS {
		return nil, res, err
	}
	defer func() {
		if errow := rows.Close(); errow != nil {
			err = errow
		}
	}()
	var likes []utils.Like
	for rows.Next() {
		var like utils.Like
		if err = rows.Scan(&like.UserID, &like.PostID); err != nil {
			return nil, ERROR, err
		}
		likes = append(likes, like)
	}
	if err = rows.Err(); err != nil {
		return nil, ERROR, err
	}
	return likes, SUCCESS, nil
}
