package api

import (
	"encoding/json"
	"net/http"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	// Set the content of the request as an application/json
	w.Header().Set("Content-type", "application/json")

	var user utils.User
	// Read the request body
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Check if the username is valid
	if !utils.HttpValidateUsername(w, user.Username) {
		return
	}

	// Check if the user exists
	exist, err := rt.db.IsUsernameExists(user.Username)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't check if the user exists")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// If not exist create a new user
	if !exist {
		user, err = rt.db.CreateUser(user)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't create the user")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusCreated)
	} else { // Else get the user from database
		user, err = rt.db.GetUserByName(user.Username)
		if err != nil {
			ctx.Logger.WithError(err).Error("can't get the user from database")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
	}

	// Encode the user and send it to the client
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		ctx.Logger.WithError(err).Error("can't encode the response")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	ctx.Logger.Info("Login successful")
}
