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

func (db *appdbimpl) GetComments(uid, pid int) ([]utils.Comment, int, error) {
	rows, err := db.c.Query(`
        SELECT c.comment_id, c.post_id, c.user_id, c.caption, c.timestamp, u.username
        FROM comments c
        JOIN users u ON c.user_id = u.user_id
        WHERE c.post_id = ?
        ORDER BY c.timestamp DESC`, pid)
	if err != nil {
		return nil, ERROR, err
	}
	defer rows.Close()

	var comments []utils.Comment
	for rows.Next() {
		var comment utils.Comment
		err := rows.Scan(
			&comment.CommentID,
			&comment.PostID,
			&comment.UserID,
			&comment.Caption,
			&comment.Timestamp,
			&comment.Username,
		)
		if err != nil {
			return nil, ERROR, err
		}
		comments = append(comments, comment)
	}

	if err = rows.Err(); err != nil {
		return nil, ERROR, err
	}
	
	return comments, SUCCESS, nil
}
