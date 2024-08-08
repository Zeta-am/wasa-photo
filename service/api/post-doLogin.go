package api

import (
	"encoding/json"
	"net/http"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"github.com/Zeta-am/wasa-photo/service/utils"
)

func (rt *_router) doLogin(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {

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
	
	// Check if the user exists, if not create it 
	
}