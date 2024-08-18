package api

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/database"
	"github.com/Zeta-am/wasa-photo/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	w.Header().Set("Content-type", "application/json")

	var post utils.Post		// The post that will be created

	//Get the uid from the url
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
	post.UserID = uid

	// Encode the request body as a multipart/form-data
	err = r.ParseMultipartForm(20 << 20)	// Max 20 MiB
	if err != nil {
		http.Error(w, "the photo exceeds the maximum length", http.StatusBadRequest)
		return
	}
	
	// Takes the caption from the URL
	post.Caption = r.URL.Query().Get("caption")
	if len(post.Caption) > 200 {
		http.Error(w, "the caption exceeds the maximum length", http.StatusBadRequest)
		return
	}

	// Get the image from the form
	imgFile, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer imgFile.Close()

	// Read the file
	image, err := io.ReadAll(imgFile)
	if err != nil {
		ctx.Logger.WithError(err).Error("error read image")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	post.Image = base64.StdEncoding.EncodeToString(image)
	post.Timestamp = time.Now().Format("2017-07-21T17:32:28")

	// Create the post in the database
	pid, res, err := rt.db.CreatePost(post)

	// Check for errors
	if res == database.ERROR {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err != nil {
		ctx.Logger.WithError(err).Error("error creating post")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	post.PostID = pid

	// Encode the response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(post)
	if err != nil {
		ctx.Logger.WithError(err).Error("error encoding the response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}		
}