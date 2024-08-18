package database


func (db *appdbimpl) FollowUser(uid int, followedId int) (int, error) {
	_, err := db.c.Exec(`INSERT
							INTO follows (follower_id, followed_id)
							VALUES (?, ?)`, uid, followedId)
	res := checkResults(err)
	if res != SUCCESS {
		return res, err
	} 
	return res, nil
}