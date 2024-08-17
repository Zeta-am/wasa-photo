package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
	func (rt *_router) Handler() http.Handler {
		// User
		rt.router.POST("/users", rt.wrap(rt.doLogin, false))
		rt.router.GET("/users/:idUser", rt.wrap(rt.getUserProfile, true))
		rt.router.GET("/users/:idUser/followers", rt.wrap(rt.listFollowers, true))
		rt.router.GET("/users/:idUser/stream", rt.wrap(rt.getMyStream, true))
	
		// Post 
		rt.router.POST("/users/:idUser/posts", rt.wrap(rt.uploadPhoto, true))

		// Comment 

		// Like

		// Follow

		// Ban 

		// Special routes
		rt.router.GET("/liveness", rt.liveness)


		return rt.router
	}
