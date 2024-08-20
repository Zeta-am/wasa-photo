package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) likePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Get the id of the photo
	pid, err := strconv.Atoi(ps.ByName("idPhoto"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Like the photo
	like, res, err := rt.db.LikePhoto(uid, pid)
	switch res {
	case database.SUCCESS:
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(like)
		if err != nil {
			http.Error(w, "can't enconde the response: "+err.Error(), http.StatusInternalServerError)
			return
		}
		return
	case database.NO_ROWS:
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	case database.UNIQUE_FAILED:
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(like)
		if err != nil {
			http.Error(w, "can't enconde the response: "+err.Error(), http.StatusInternalServerError)
			return
		}
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
}