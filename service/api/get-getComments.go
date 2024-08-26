package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Get the post_id from the url
	pid, err := strconv.Atoi(ps.ByName("idPhoto"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return 
	}

	// Get the comments from database
	comms, res, err := rt.db.GetComments(uid, pid)
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(comms); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}