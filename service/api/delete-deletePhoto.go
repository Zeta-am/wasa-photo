package api

import (
	"encoding/json"	
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/julienschmidt/httprouter"
)


func (rt *_router) deletePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check if the photo exists
	post, res, err := rt.db.GetPostById(pid)
	if res == database.NO_ROWS {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	} else if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return						
	}

	// Delete the photo
	res, err = rt.db.DeletePost(pid)
	switch res {
	case database.NO_ROWS:
		http.Error(w, "post to be deleted was not found", http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	return		
	}
	ctx.Logger.Info("Post deleted")
}