package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
func (rt *_router) Handler() http.Handler {
	// User
	rt.router.POST("/users", rt.wrap(rt.doLogin, false))
	rt.router.GET("/users/:idUser", rt.wrap(rt.getUserProfile, true))

	// Post 

	// Comment 

	// Like

	// Follow

	// Ban 

	// Special routes
	rt.router.GET("/liveness", rt.liveness)


	return rt.router
}
