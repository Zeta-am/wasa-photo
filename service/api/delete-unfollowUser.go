package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) unfollowUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	unfollowUid, err := strconv.Atoi(ps.ByName("idFollowed"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Unfollow the user
	res, err := rt.db.UnfollowUser(uid, unfollowUid)
	if res != database.SUCCESS {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get updated user profile with correct followed status
	user, res, err := rt.db.GetUserById(unfollowUid, uid)
	if res != database.SUCCESS {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return response
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(user)
}
