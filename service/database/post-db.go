package database

import "github.com/Zeta-am/wasa-photo/service/utils"

func (db *appdbimpl) CreatePost(p utils.Post) (int, int, error) {
	var pid int
	err := db.c.QueryRow(`INSERT 
								INTO posts (user_id, image, timestamp, caption)
								VALUES (?, ?, ?, ?)
								RETURNING post_id`, p.UserID, p.Image, p.Timestamp, p.Caption).Scan(&pid)
	return pid, checkResults(err), err
}

func (db *appdbimpl) DeletePost(pid int) (int, error) {
	// Inizia una transazione
	tx, err := db.c.Begin()
	if err != nil {
		return ERROR, err
	}

	// Elimina prima i commenti collegati
	_, err = tx.Exec(`DELETE FROM comments WHERE post_id = ?`, pid)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = rollbackErr
		}
		return ERROR, err
	}

	// Elimina i like collegati
	_, err = tx.Exec(`DELETE FROM likes WHERE post_id = ?`, pid)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = rollbackErr
		}
		return ERROR, err
	}

	// Elimina il post
	_, err = tx.Exec(`DELETE FROM posts WHERE post_id = ?`, pid)
	if err != nil {
		if rollbackErr := tx.Rollback(); rollbackErr != nil {
			err = rollbackErr
		}
		return ERROR, err
	}

	// Commit della transazione
	if err = tx.Commit(); err != nil {
		return ERROR, err
	}

	return SUCCESS, nil
}

func (db *appdbimpl) GetPostById(pid int) (utils.Post, int, error) {
	var post utils.Post
	err := db.c.QueryRow(`SELECT * 
							FROM posts
							WHERE post_id = ?`, pid).Scan(&post.PostID, &post.UserID, &post.Image, &post.Timestamp, &post.Caption)
	if res := checkResults(err); res != SUCCESS {
		return utils.Post{}, res, err
	}
	return db.fillPost(post)
}

func (db *appdbimpl) fillPost(p utils.Post) (utils.Post, int, error) {
	// Fill the field CommentCount
	err := db.c.QueryRow(`SELECT COUNT(*)
							FROM comments
							WHERE post_id = ?`, p.PostID).Scan(&p.CommentCount)
	if err != nil {
		return p, ERROR, err
	}

	// Fill the field LikeCount
	err = db.c.QueryRow(`SELECT COUNT(*)
							FROM likes
							WHERE post_id = ?`, p.PostID).Scan(&p.LikeCount)
	if err != nil {
		return p, ERROR, err
	}
	return p, SUCCESS, nil
}
