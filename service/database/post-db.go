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

func (db *appdbimpl) GetPostById(pid int) (utils.Post, int, error) {
	var post utils.Post
	err := db.c.QueryRow(`SELECT * 
							FROM posts
							WHERE post_id = ?`, pid).Scan(&post.PostID, &post.UserID, &post.Image, &post.Timestamp, &post.Caption)	
	if res := checkResults(err); res != SUCCESS {
		return utils.Post{}, res, err
	}
	return post, SUCCESS, nil
}