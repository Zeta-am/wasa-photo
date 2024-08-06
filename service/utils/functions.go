package utils

import (
	"net/http"
	"strconv"
	"strings"
)


func GetAuthorization(w http.ResponseWriter, r *http.Request) (int, error){
	auth := strings.Split(r.Header.Get("Authorization"), " ")
	if len(auth) <= 1 {
		return 0, ErrUnauthorized
	} 
	// Get the header authorization value
	uid, err := strconv.Atoi(auth[1])
	if err != nil {
		return 0, ErrUnauthorized
	}
	return uid, nil
}

func SetHeaderJson(w http.ResponseWriter) {
	w.Header().Set("Content-type", "application/json")
}

func SetHeaderText(w http.ResponseWriter) {
	w.Header().Set("Content-type", "text/plain")
}