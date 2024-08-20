package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) banUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")

	// Get the user id from the URL
	uid, err := strconv.Atoi(ps.ByName("idUser"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is authorized
	if uid != ctx.UserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the user id you want to ban
	banId, err := strconv.Atoi(ps.ByName("idUserBlocked"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is trying to ban himself
	if uid == banId {
		http.Error(w, "you can't ban yourself", http.StatusBadRequest)
		return
	}

	// Ban the user
	res, err := rt.db.BanUser(uid, banId)
	switch res {
	case database.UNIQUE_FAILED:	// The user was already banned
		w.WriteHeader(http.StatusOK)
		return
	case database.NO_ROWS:
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the followed user 
	user, res, err := rt.db.GetUserById(banId)

	// Check for errors
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}