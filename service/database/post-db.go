package database

import "github.com/Zeta-am/wasa-photo/service/utils"


func (db *appdbimpl) CreatePost(p utils.Post) (utils.Post, error) { 
	res, err := db.c.Exec(`INSERT 
							INTO posts (user_id, image, timestamp, caption)
							VALUES (?, ?, ?, ?)`, p.UserID, p.Image, p.Timestamp, p.Caption)
	if err != nil {
		return utils.Post{}, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return utils.Post{}, err
	}
	p.PostID = int(id)
	return p, nil
}