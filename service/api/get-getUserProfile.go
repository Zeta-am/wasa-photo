package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserProfile(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")

	// Get the uid from the url
	uid, err := strconv.Atoi(ps.ByName("idUser"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the user is banned
	banned, _, err := rt.db.IsBanned(ctx.UserID, uid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if banned {
		http.Error(w, "access denied. You cannot perform this action because the user is banned.", http.StatusForbidden)
		return
	}

	// Get the user from database
	currentUserId := ctx.UserID
	dbUser, res, err := rt.db.GetUserById(uid, currentUserId)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		ctx.Logger.WithError(err).Error("can't get the user from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dbUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
