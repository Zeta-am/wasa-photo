package database

import "github.com/Zeta-am/wasa-photo/service/utils"


func (db *appdbimpl) CreatePost(p utils.Post) (int, int, error) { 
	var pid int
	err := db.c.QueryRow(`INSERT 
							INTO posts (user_id, image, timestamp, caption)
							VALUES (?, ?, ?, ?)`, p.UserID, p.Image, p.Timestamp, p.Caption).Scan(&pid)
	return pid, checkResults(err), err
}