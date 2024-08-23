package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) followUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Get the user id you want to follow
	followUid, err := strconv.Atoi(ps.ByName("idFollowed"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if userId is banned from followUid
	banned, _, err := rt.db.IsBanned(uid, followUid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if banned {
		http.Error(w, "access denied. You cannot perform this action because the user is banned.", http.StatusForbidden)
		return
	}

	// Check that you are not trying to follow a user banned by us
	banned, _, err = rt.db.IsBanned(followUid, uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if banned {
		http.Error(w, "access denied. You cannot perform this action because the user is banned.", http.StatusForbidden)
		return
	}
	
	// Follow the user 
	res, err := rt.db.FollowUser(uid, followUid)
	switch res {
	case database.UNIQUE_FAILED:	// The user was already followed
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
	user, res, err := rt.db.GetUserById(followUid)

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
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	
}