package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
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

	// Check if the ID of the path is equal to the ID of the authorization
	if uid != ctx.UserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the user from database
	dbUser, err := rt.db.GetUserProfile(uid)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't get the user from database")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	//TODO: check if the user is banned

	// Send the response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dbUser)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
