package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) getUserPhotos(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")

	// Get the uid from the url
	uid, err := strconv.Atoi(ps.ByName("idUser"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Get the posts from database
	posts, res, err := rt.db.GetUserPhotos(uid)
	if res == database.ERROR || err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response
	w.WriteHeader(http.StatusOK)
	if err = json.NewEncoder(w).Encode(posts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
