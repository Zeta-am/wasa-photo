package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/Zeta-am/wasa-photo/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) commentPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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
	post, res, err := rt.db.GetPostById(pid)
	if res == database.NO_ROWS {
		http.Error(w, "post not found", http.StatusNotFound)
		return
	} else if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// TODO: Check that the user who posted the photo has not banned the user who wants to comment
	banned, _, err := rt.db.IsBanned(uid, post.UserID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if banned {
		http.Error(w, "access denied. You cannot perform this action because the user is banned.", http.StatusForbidden)
		return
	}

	// Read the request body
	var comm utils.Comment
	err = json.NewDecoder(r.Body).Decode(&comm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Fill the comment object
	comm.PostID = pid
	comm.UserID = uid
	comm.Timestamp = time.Now().Format(time.RFC3339)

	// Create the comment
	cid, res, err := rt.db.CreateComment(comm)
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
	comm.CommentID = cid
	err = json.NewEncoder(w).Encode(comm)
	if err != nil {
		http.Error(w, "can't encode the response: "+err.Error(), http.StatusInternalServerError)
		return
	}

}
