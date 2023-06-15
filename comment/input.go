package comment

import "hacktiv-final2/user"

type Input struct {
	Message string `json:"message" binding:"required"`
	PhotoID int    `json:"photo_id" binding:"required"`
	User    user.User
}

type UpdateInput struct {
	Message string `json:"message" binding:"required"`
	User    user.User
}

type GetCommentInput struct {
	ID int `uri:"id" binding:"required"`
}
