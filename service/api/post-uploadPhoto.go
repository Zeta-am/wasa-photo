package api

import (
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/Zeta-am/wasa-photo/service/api/reqcontext"
	"github.com/Zeta-am/wasa-photo/service/utils"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) uploadPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
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

	// Encode the request body as a multipart/form-data
	err = r.ParseMultipartForm(20 << 20)	// Max 20 MiB
	if err != nil {
		http.Error(w, "the photo exceeds the maximum length", http.StatusBadRequest)
		return
	}
	
	// Takes the caption from the URL
	caption := r.URL.Query().Get("caption")
	if len(caption) > 200 {
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

	// Create the post
	var post utils.Post
	post.UserID = uid
	post.Image = base64.StdEncoding.EncodeToString(image)
	post.Caption = caption
	post.Timestamp = time.Now().Format("2017-07-21T17:32:28")

	dbPost, err := rt.db.CreatePost(post)

	if err != nil {
		ctx.Logger.WithError(err).Error("error creating post")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Encode the response
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(dbPost)
	if err != nil {
		ctx.Logger.WithError(err).Error("error encoding the response")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}		
}