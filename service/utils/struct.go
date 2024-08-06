package utils



type User struct {
	UserID         int    `json:"id"`
	Username       string `json:"username"`
	Name           string `json:"name"`
	Surname        string `json:"surname"`
	PostCount      int    `json:"postNo"`
	FollowerCount  int    `json:"followerNo"`
	FollowingCount int    `json:"followingNo"`
}

type Post struct {
	PostID int `json:"id"`
	Image string `json:"image"`
	UserID int `json:"user-id"`
	LikeCount int `json:"like-numbers"`
	CommentCount int `json:"comment-numbers"`
	Timestamp string `json:"upload-time"`
}

type Comment struct {
	CommentID int `json:"id"`
	UserID int `json:"userId"`
	PostID int `json:"photoId"`
	Timestamp string `json:"upload-time"`
	Description string `json:"comment"`
}

type Like struct {
	UserID int `json:"userId"`
	PostID int `json:"photoId"`
}