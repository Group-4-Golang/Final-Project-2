package comment

import "time"

type Formatter struct {
	ID      int       `json:"id"`
	Message string    `json:"message"`
	PhotoID int       `json:"photo_id"`
	UserID  int       `json:"user_id"`
	Created time.Time `json:"created_at"`
}

func FormatComment(comment Comment) Formatter {
	formatter := Formatter{
		ID:      comment.ID,
		Message: comment.Message,
		PhotoID: comment.PhotoID,
		UserID:  comment.UserID,
		Created: comment.CreatedAt,
	}

	return formatter
}

type GetFormatter struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	User      struct {
		Id       int    `json:"id"`
		Email    string `json:"email"`
		Username string `json:"username"`
	} `json:"User"`
	Photo struct {
		Id       int    `json:"id"`
		Title    string `json:"title"`
		Caption  string `json:"caption"`
		PhotoUrl string `json:"photo_url"`
		UserId   int    `json:"user_id"`
	} `json:"Photo"`
}

func FormatGetComment(comment []Comment) GetFormatter {
	var formatter GetFormatter
	for _, value := range comment {
		formatter.Id = value.ID
		formatter.Message = value.Message
		formatter.PhotoId = value.PhotoID
		formatter.UserId = value.UserID
		formatter.UpdatedAt = value.UpdatedAt
		formatter.CreatedAt = value.CreatedAt
		formatter.User.Id = value.User.ID
		formatter.User.Email = value.User.Email
		formatter.User.Username = value.User.Username
		formatter.Photo.Id = value.Photo.ID
		formatter.Photo.Title = value.Photo.Title
		formatter.Photo.Caption = value.Photo.Caption
		formatter.Photo.PhotoUrl = value.Photo.PhotoURL
		formatter.Photo.UserId = value.Photo.UserID
	}
	return formatter
}

func FormatGetAllComment(comment []Comment) []GetFormatter {
	var formatter []GetFormatter
	for _, value := range comment {
		formatter = append(formatter, GetFormatter{
			Id:        value.ID,
			Message:   value.Message,
			PhotoId:   value.PhotoID,
			UserId:    value.UserID,
			UpdatedAt: value.UpdatedAt,
			CreatedAt: value.CreatedAt,
			User: struct {
				Id       int    `json:"id"`
				Email    string `json:"email"`
				Username string `json:"username"`
			}{
				Id:       value.User.ID,
				Email:    value.User.Email,
				Username: value.User.Username,
			},
			Photo: struct {
				Id       int    `json:"id"`
				Title    string `json:"title"`
				Caption  string `json:"caption"`
				PhotoUrl string `json:"photo_url"`
				UserId   int    `json:"user_id"`
			}{
				Id:       value.Photo.ID,
				Title:    value.Photo.Title,
				Caption:  value.Photo.Caption,
				PhotoUrl: value.Photo.PhotoURL,
				UserId:   value.Photo.UserID,
			},
		})
	}
	return formatter
}

type UpdateFormatter struct {
	Id        int       `json:"id"`
	Message   string    `json:"message"`
	PhotoId   int       `json:"photo_id"`
	UserId    int       `json:"user_id"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatUpdateComment(comment Comment) UpdateFormatter {
	formatter := UpdateFormatter{
		Id:        comment.ID,
		Message:   comment.Message,
		PhotoId:   comment.PhotoID,
		UserId:    comment.UserID,
		UpdatedAt: comment.UpdatedAt,
	}

	return formatter
}

type DeleteFormatter struct {
	Message string `json:"message"`
}

func FormatDeleteComment() DeleteFormatter {
	formatter := DeleteFormatter{
		Message: "Your comment has been successfully deleted",
	}

	return formatter
}
