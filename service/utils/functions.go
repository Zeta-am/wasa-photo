package utils

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"
)

func GetAuthorization(w http.ResponseWriter, r *http.Request) (int, error) {
	auth, err := strconv.Atoi(r.Header.Get("Authorization"))
	if err != nil {
		return 0, nil
	}
	return auth, nil
}

func ValidateUsername(username string) error {
	if len(username) == 0 {
		return ErrMissingUsername
	}
	// Check that the username is alphanumeric between 3 and 16
	isMatch, err := regexp.MatchString("^[a-zA-Z0-9._]{3,16}$", username)
	if !isMatch || err != nil {
		return ErrUsernameNotValid
	}
	return nil
}

func HttpValidateUsername(w http.ResponseWriter, username string) bool {
	err := ValidateUsername(username)
	if errors.Is(err, ErrUsernameNotValid) {
		http.Error(w, err.Error(), http.StatusNotAcceptable)
		return false
	}
	if errors.Is(err, ErrMissingUsername) {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return false
	}
	return true
}
