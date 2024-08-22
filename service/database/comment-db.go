package database

import "github.com/Zeta-am/wasa-photo/service/utils"

func (db *appdbimpl) CreateComment(c utils.Comment) (int, int, error) {
	var cid int
	err := db.c.QueryRow(`INSERT
							INTO comments(user_id, post_id, timestamp, caption)
							VALUES (?, ?, ?, ?)
							RETURNING comm_id`, &c.UserID, &c.PostID, &c.Timestamp, &c.Caption).Scan(&cid)
	if res := checkResults(err); res != SUCCESS {
		return cid, res, err
	}
	return cid, SUCCESS, nil	
}

func (db *appdbimpl) DeleteComment(cid int, pid int, uid int) (int, error) {
	_, err := db.c.Exec(`DELETE
							FROM comments
							WHERE comm_id = ? AND post_id = ? AND user_id = ?`, cid, pid, uid)
	if res := checkResults(err); res != SUCCESS {
		return res, err
	}
	return SUCCESS, nil
}
	