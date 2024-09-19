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

func (rt *_router) uncommentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Check if the photo exists
	_, res, err := rt.db.GetPostById(pid)
	if res == database.NO_ROWS {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	} else if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Get the id of the comments
	cid, err := strconv.Atoi(ps.ByName("idComment"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Delete the comment
	res, err = rt.db.DeleteComment(cid, pid, uid)
	switch res {
	case database.NO_ROWS:
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	case database.ERROR:
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the deleted comment
	w.WriteHeader(http.StatusOK)
	var delComm = utils.Comment{
		CommentID: cid,
		PostID:    pid,
		UserID:    uid,
	}
	err = json.NewEncoder(w).Encode(delComm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
