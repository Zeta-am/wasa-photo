// get-getUserProfileByUsername.go
package api

import (
	"encoding/json"
	"net/http"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/Zeta-am/wasa-photo/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")

	// Get username from query parameter
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "Missing username parameter", http.StatusBadRequest)
		return
	}

	// Get users from database
	users, res, err := rt.db.GetUsersByPattern(username, ctx.UserID)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err != nil {
		ctx.Logger.WithError(err).Error("can't get users from database")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Build response
	response := struct {
		Users []utils.User `json:"users"`
	}{
		Users: users,
	}

	// Send response
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
