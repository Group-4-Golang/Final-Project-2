package photo

import "time"

type PostFormatter struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoURL  string    `json:"photo_url"`
	UserID    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatPhoto(photo Photo) PostFormatter {
	formatter := PostFormatter{
		ID:        photo.ID,
		Title:     photo.Title,
		Caption:   photo.Caption,
		PhotoURL:  photo.PhotoURL,
		UserID:    photo.UserID,
		CreatedAt: photo.CreatedAt,
	}

	return formatter
}

type GetFormatter struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Caption   string    `json:"caption"`
	PhotoUrl  string    `json:"photo_url"`
	UserId    int       `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	User      struct {
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"User"`
}

func FormatGetPhoto(photo []Photo) GetFormatter {
	var formatter GetFormatter
	for _, value := range photo {
		formatter.Id = value.ID
		formatter.Title = value.Title
		formatter.Caption = value.Caption
		formatter.PhotoUrl = value.PhotoURL
		formatter.UserId = value.UserID
		formatter.CreatedAt = value.CreatedAt
		formatter.UpdatedAt = value.UpdatedAt
		formatter.User.Email = value.User.Email
		formatter.User.Username = value.User.Username
	}
	return formatter
}

func FormatGetAllPhoto(photo []Photo) []GetFormatter {
	var formatter []GetFormatter
	for _, value := range photo {
		formatter = append(formatter, GetFormatter{
			Id:        value.ID,
			Title:     value.Title,
			Caption:   value.Caption,
			PhotoUrl:  value.PhotoURL,
			UserId:    value.UserID,
			CreatedAt: value.CreatedAt,
			UpdatedAt: value.UpdatedAt,
			User: struct {
				Email    string `json:"email"`
				Username string `json:"username"`
			}{
				Email:    value.User.Email,
				Username: value.User.Username,
			},
		})
	}
	return formatter
}

type UpdateFormatter struct {
	Id       int       `json:"id"`
	Title    string    `json:"title"`
	Caption  string    `json:"caption"`
	PhotoUrl string    `json:"photo_url"`
	UserId   int       `json:"user_id"`
	UpdateAt time.Time `json:"updated_at"`
}

func FormatUpdatePhoto(photo Photo) UpdateFormatter {
	formatter := UpdateFormatter{
		Id:       photo.ID,
		Title:    photo.Title,
		Caption:  photo.Caption,
		PhotoUrl: photo.PhotoURL,
		UserId:   photo.UserID,
		UpdateAt: photo.UpdatedAt,
	}

	return formatter
}

type DeleteFormatter struct {
	Message string `json:"message"`
}

func FormatDeletePhoto() DeleteFormatter {
	formatter := DeleteFormatter{
		Message: "Your photo has been successfully deleted",
	}

	return formatter
}
