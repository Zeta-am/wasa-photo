package database

import (
	"database/sql"
	"errors"
	"strings"

	"github.com/Zeta-am/wasa-photo/service/utils"
)

const (
	SUCCESS = iota
	ERROR
	NO_ROWS
	UNIQUE_FAILED
)

func checkResults(err error) int {
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) || strings.Contains(err.Error(), "KEY constraint failed") {
			return NO_ROWS
		}
		if strings.Contains(err.Error(), "UNIQUE constraint failed") {
			return UNIQUE_FAILED
		}
		if strings.Contains(err.Error(), "converting NULL") {
			return SUCCESS
		}
		return ERROR
	}

	return SUCCESS
}

func getUsers(rows *sql.Rows) ([]utils.User, int, error) {
	var err error = nil
	var users []utils.User
	defer func() {
		if closeErr := rows.Close(); closeErr != nil {
			err = closeErr
		}
	}()

	// Get the user_id and username from the rows selected
	for rows.Next() {
		var usr utils.User
		if err = rows.Scan(&usr.UserID, &usr.Username); err != nil {
			return nil, ERROR, err
		}
		users = append(users, usr)
	}

	if err = rows.Err(); err != nil {
		return nil, ERROR, err
	}
	return users, SUCCESS, nil
}
