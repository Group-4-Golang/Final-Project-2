package user

import "time"

type Formatter struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Age      int    `json:"age"`
}

func FormatUser(user User) Formatter {
	formatter := Formatter{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}

	return formatter
}

type LoginFormatter struct {
	Token string `json:"token"`
}

func FormatLogin(user User, token string) LoginFormatter {
	formatter := LoginFormatter{
		Token: token,
	}

	return formatter
}

type UpdateFormatter struct {
	ID       int       `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Age      int       `json:"age"`
	UpdateAt time.Time `json:"update_at"`
}

func FormatUpdate(user User) UpdateFormatter {
	formatter := UpdateFormatter{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
		UpdateAt: time.Now(),
	}

	return formatter
}

type DeleteFormatter struct {
	Message string `json:"message"`
}

func FormatDelete() DeleteFormatter {
	formatter := DeleteFormatter{
		Message: "Your account has been successfully deleted",
	}

	return formatter
}
