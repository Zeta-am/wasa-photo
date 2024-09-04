package api

import (
	"net/http"
)

// Handler returns an instance of httprouter.Router that handle APIs registered here
	func (rt *_router) Handler() http.Handler {
		// User
		rt.router.POST("/users", rt.wrap(rt.doLogin, false))
		rt.router.GET("/users/:idUser", rt.wrap(rt.getUserProfile, true))
		rt.router.GET("/users/:idUser/stream", rt.wrap(rt.getMyStream, true))
		rt.router.PUT("/users/:idUser/edit", rt.wrap(rt.setMyUserName, true))
		rt.router.GET("/users/:idUser/posts", rt.wrap(rt.getUserPhotos, true))
	
		// Post 
		rt.router.POST("/users/:idUser/posts", rt.wrap(rt.uploadPhoto, true))
		rt.router.DELETE("/users/:idUser/posts/:idPhoto", rt.wrap(rt.deletePhoto, true))

		// Comment 
		rt.router.POST("/users/:idUser/posts/:idPhoto/comments", rt.wrap(rt.commentPhoto, true))
		rt.router.DELETE("/users/:idUser/posts/:idPhoto/comments/:idComment", rt.wrap(rt.uncommentPhoto, true))	
		rt.router.GET("/users/:idUser/posts/:idPhoto/comments", rt.wrap(rt.getComments, true))

		// Like
		rt.router.PUT("/users/:idUser/posts/:idPhoto/likes", rt.wrap(rt.likePhoto, true))
		rt.router.DELETE("/users/:idUser/posts/:idPhoto/likes", rt.wrap(rt.unlikePhoto, true))
		rt.router.GET("/users/:idUser/posts/:idPhoto/likes", rt.wrap(rt.getLikes, true))	

		// Follow
		rt.router.PUT("/users/:idUser/followings/:idFollowed", rt.wrap(rt.followUser, true))
		rt.router.DELETE("/users/:idUser/followings/:idFollowed", rt.wrap(rt.unfollowUser, true))
		rt.router.GET("/users/:idUser/followers", rt.wrap(rt.listFollowers, true))
		rt.router.GET("/users/:idUser/followings", rt.wrap(rt.listFollowings, true))

		// Ban 
		rt.router.PUT("/users/:idUser/banList/:idUserBlocked", rt.wrap(rt.banUser, true))
		rt.router.DELETE("/users/:idUser/banList/:idUserBlocked", rt.wrap(rt.unbanUser, true))
		rt.router.GET("/users/:idUser/banList", rt.wrap(rt.getBannedList, true))

		// Special routes
		rt.router.GET("/liveness", rt.liveness)


		return rt.router
	}
