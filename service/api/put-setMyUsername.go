package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/Zeta-am/wasa-photo/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) setMyUserName(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Read the request body
	var usrUpdated utils.User
	err = json.NewDecoder(r.Body).Decode(&usrUpdated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Validate the username
	if !utils.HttpValidateUsername(w, usrUpdated.Username) {
		return
	}

	// Check if the username exists
	exist, _, _ := rt.db.IsUsernameExists(usrUpdated.Username)
	if exist {
		http.Error(w, "the username already exists", http.StatusBadRequest)
		return
	}

	// Change the username
	res, err := rt.db.SetMyUsername(usrUpdated.Username, uid)
	switch res {
	case database.NO_ROWS:
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the user updated
	usr, res, err := rt.db.GetUserById(uid)
	switch res {
	case database.NO_ROWS:
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(usr)
	if err != nil {
		http.Error(w, "can't encode the response", http.StatusInternalServerError)
		return
	}
}