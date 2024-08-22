package database

func (db *appdbimpl) BanUser(uid int, bannedId int) (int, error) {
	// Begin a transaction
	tx, err := db.c.Begin()
	if err != nil {
		return ERROR, err
	}

	// Ban the user
	_, err = tx.Exec(`INSERT
						INTO bans (user_id, banned_id)
						VALUES (?, ?)`, uid, bannedId)					
	res := checkResults(err)
	if res != SUCCESS {
		erro := tx.Rollback()
		if erro != nil {
			err = erro
		}
		return res, err
	}

	// Unfollow the user
	_, err = tx.Exec(`DELETE
						FROM follows
						WHERE (follower_id = ? AND followed_id = ?) 
						OR (follower_id = ? AND followed_id = ?)`, uid, bannedId, bannedId, uid)
	res = checkResults(err)
	if res != SUCCESS {
		erro := tx.Rollback()
		if erro != nil {
			err = erro
		}
		return res, err
	}

	// Commit the transaction
	err = tx.Commit()
	if err != nil {
		erro := tx.Rollback()
		if erro != nil {
			err = erro
		}
		return res, err
	}

	return SUCCESS, nil
}

func (db appdbimpl) UnbanUser(uid int, unbannedId int) (int, error) {
	_, err := db.c.Exec(`DELETE
							FROM bans
							WHERE user_id = ? AND banned_id = ?`, uid, unbannedId)
	if res := checkResults(err); res != SUCCESS {
		return res, err
	}
	return SUCCESS, nil
}


// Check if uid is banned by bannerId
func (db appdbimpl) IsBanned(uid int, bannerId int) (bool, int, error) {
	var value int 
	err := db.c.QueryRow(`SELECT EXISTS(
							SELECT 1
							FROM bans
							WHERE user_id = ? AND banned_id = ?)`, bannerId, uid).Scan(&value)
	if res := checkResults(err); res != SUCCESS || value == 0{
		return false, res, err
	}
	return true, SUCCESS, nil
}