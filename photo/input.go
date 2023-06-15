package photo

import "hacktiv-final2/user"

type Input struct {
	Title    string `json:"title" binding:"required"`
	Caption  string `json:"caption"`
	PhotoURL string `json:"photo_url" binding:"required"`
	User     user.User
}

type GetPhotoInput struct {
	ID int `uri:"id" binding:"required"`
}
