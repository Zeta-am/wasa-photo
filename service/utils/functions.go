package utils

import (
	"errors"
	"net/http"
	"regexp"
	"strconv"	
)

func GetAuthorization(w http.ResponseWriter, r *http.Request) (int, error) {
	// auth := strings.Split(r.Header.Get("Authorization"), " ")
	// if len(auth) <= 1 {
	// 	return 0, ErrUnauthorized
	// }
	// // Get the header authorization value
	// uid, err := strconv.Atoi(auth[1])
	// if err != nil {
	// 	return 0, ErrUnauthorized
	// }

	auth, err := strconv.Atoi(r.Header.Get("Authorization"))
	if err != nil {
		return 0, nil
	}
	return auth, nil	
}

func SetHeaderJson(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
}

func SetHeaderText(w http.ResponseWriter) {
	w.Header().Set("Content-type", "text/plain")
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
