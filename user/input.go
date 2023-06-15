package user

type RegisterUserInput struct {
	Username string `json:"username" binding:"required" gorm:"unique"`
	Email    string `json:"email" binding:"email,required" gorm:"unique"`
	Password string `json:"password" binding:"required,min=6"`
	Age      int    `json:"age" binding:"required,min=8"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"email,required"`
	Password string `json:"password" binding:"required,min=6"`
}
