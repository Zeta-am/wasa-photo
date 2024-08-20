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
		rt.router.POST("/users/:idUser/posts/:idPhoto/comments", rt.wrap(rt.commentPhoto, true))

		// Like
		rt.router.PUT("/users/:idUser/posts/:idPhoto/likes", rt.wrap(rt.likePhoto, true))
		rt.router.DELETE("/users/:idUser/posts/:idPhoto/likes", rt.wrap(rt.unlikePhoto, true))

		// Follow
		rt.router.PUT("/users/:idUser/followings/:idFollowed", rt.wrap(rt.followUser, true))
		rt.router.DELETE("/users/:idUser/followings/:idFollowed", rt.wrap(rt.unfollowUser, true))

		// Ban 
		rt.router.PUT("/users/:idUser/banList/:idUserBlocked", rt.wrap(rt.banUser, true))

		// Special routes
		rt.router.GET("/liveness", rt.liveness)


		return rt.router
	}
