package api

import (
	"net/http"
	"strconv"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) listFollowers(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")

	// Get uid from the url
	uid, err := strconv.Atoi(ps.ByName("idUser"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}	

	// Check if the ID of the path is equal to the ID of the authorization
	if uid != ctx.UserID {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	// Get the list of the followers from the database
		

}
