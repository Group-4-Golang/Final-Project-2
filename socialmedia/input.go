package socialmedia

import "hacktiv-final2/user"

type Input struct {
	Name           string `json:"name" binding:"required"`
	SocialMediaURL string `json:"social_media_url" binding:"required"`
	User           user.User
}

type GetSocialMediaInput struct {
	ID int `uri:"id" binding:"required"`
}
