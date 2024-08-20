package database

func (db *appdbimpl) BanUser(uid int, bannedId int) (int, error) {
	_, err := db.c.Exec(`INSERT
						INTO bans (user_id, banned_id)
						VALUES (?, ?)`, uid, bannedId)					
	res := checkResults(err)
	if res != SUCCESS {
		return res, err
	}
	return res, nil

	// TODO: L'utente bannato (bannedId) non potra' piu' visualizzare foto, like e commenti dell'utente 
	// che lo sta bannando; inoltre il follow viene rimosso
}