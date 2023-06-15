package socialmedia

import "time"

type PostFormatter struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaURL string    `json:"social_media_url"`
	UserID         int       `json:"user_id"`
	CreatedAt      time.Time `json:"created_at"`
}

func FormatSocialMedia(socialmedia SocialMedia) PostFormatter {
	formatter := PostFormatter{
		ID:             socialmedia.ID,
		Name:           socialmedia.Name,
		SocialMediaURL: socialmedia.SocialMediaURL,
		UserID:         socialmedia.UserID,
		CreatedAt:      socialmedia.CreatedAt,
	}

	return formatter
}

func FormatSocialMedias(socialmedias []SocialMedia) []PostFormatter {
	if len(socialmedias) == 0 {
		return []PostFormatter{}
	}

	var socialmediasFormatter []PostFormatter
	for _, value := range socialmedias {
		socialmediasFormatter = append(socialmediasFormatter, FormatSocialMedia(value))
	}

	return socialmediasFormatter
}

type GetFormatter struct {
	SocialMedias []struct {
		Id             int       `json:"id"`
		Name           string    `json:"name"`
		SocialMediaUrl string    `json:"social_media_url"`
		UserId         int       `json:"UserId"`
		CreatedAt      time.Time `json:"createdAt"`
		UpdatedAt      time.Time `json:"updatedAt"`
		User           struct {
			Id       int    `json:"id"`
			Username string `json:"username"`
		} `json:"User"`
	} `json:"social_medias"`
}

func FormatGetSocialMedia(socialmedia []SocialMedia) GetFormatter {
	var formatter GetFormatter
	for _, value := range socialmedia {
		formatter.SocialMedias = append(formatter.SocialMedias, struct {
			Id             int       `json:"id"`
			Name           string    `json:"name"`
			SocialMediaUrl string    `json:"social_media_url"`
			UserId         int       `json:"UserId"`
			CreatedAt      time.Time `json:"createdAt"`
			UpdatedAt      time.Time `json:"updatedAt"`
			User           struct {
				Id       int    `json:"id"`
				Username string `json:"username"`
			} `json:"User"`
		}{
			Id:             value.ID,
			Name:           value.Name,
			SocialMediaUrl: value.SocialMediaURL,
			UserId:         value.UserID,
			CreatedAt:      value.CreatedAt,
			UpdatedAt:      value.UpdatedAt,
			User: struct {
				Id       int    `json:"id"`
				Username string `json:"username"`
			}{
				Id:       value.User.ID,
				Username: value.User.Username,
			},
		})
	}

	return formatter
}

type UpdateFormatter struct {
	Id             int       `json:"id"`
	Name           string    `json:"name"`
	SocialMediaUrl string    `json:"social_media_url"`
	UserId         int       `json:"UserId"`
	UpdateAt       time.Time `json:"updateAt"`
}

func FormatUpdateSocialMedia(socialmedia SocialMedia) UpdateFormatter {
	formatter := UpdateFormatter{
		Id:             socialmedia.ID,
		Name:           socialmedia.Name,
		SocialMediaUrl: socialmedia.SocialMediaURL,
		UserId:         socialmedia.UserID,
		UpdateAt:       socialmedia.UpdatedAt,
	}

	return formatter
}

type DeleteFormatter struct {
	Message string `json:"message"`
}

func FormatDeleteSocialMedia() DeleteFormatter {
	formatter := DeleteFormatter{
		Message: "Your social media has been successfully deleted",
	}

	return formatter
}
