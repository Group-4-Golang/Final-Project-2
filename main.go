package main

import (
	"hacktiv-final2/auth"
	"hacktiv-final2/comment"
	"hacktiv-final2/handler"
	"hacktiv-final2/helper"
	"hacktiv-final2/photo"
	"hacktiv-final2/socialmedia"
	"hacktiv-final2/user"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := helper.LoadConfig()
	db := helper.InitDB()

	authService := auth.NewService()

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService, authService)

	photoRepository := photo.NewRepository(db)
	photoService := photo.NewService(photoRepository)
	photoHandler := handler.NewPhotoHandler(photoService, authService)

	commentRepository := comment.NewRepository(db)
	commentService := comment.NewService(commentRepository, photoRepository)
	commentHandler := handler.NewCommentHandler(commentService, authService)

	socialMediaRepository := socialmedia.NewRepository(db)
	socialMediaService := socialmedia.NewService(socialMediaRepository)
	socialMediaHandler := handler.NewSocialMediaHandler(socialMediaService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	// User
	api.POST("/users/register", userHandler.RegisterUser)
	api.POST("/users/login", userHandler.LoginUser)
	api.PUT("/users", authService.AuthMiddleware(authService, userService), userHandler.UpdateUser)
	api.DELETE("/users", authService.AuthMiddleware(authService, userService), userHandler.DeleteUser)

	// Photo
	api.POST("/photos", authService.AuthMiddleware(authService, userService), photoHandler.CreatePhoto)
	api.GET("/photos", authService.AuthMiddleware(authService, userService), photoHandler.GetPhoto)
	api.PUT("/photos/:id", authService.AuthMiddleware(authService, userService), photoHandler.UpdatePhoto)
	api.DELETE("/photos/:id", authService.AuthMiddleware(authService, userService), photoHandler.DeletePhoto)

	// Comment
	api.POST("/comments", authService.AuthMiddleware(authService, userService), commentHandler.CreateComment)
	api.GET("/comments", authService.AuthMiddleware(authService, userService), commentHandler.GetComment)
	api.PUT("/comments/:id", authService.AuthMiddleware(authService, userService), commentHandler.UpdateComment)
	api.DELETE("/comments/:id", authService.AuthMiddleware(authService, userService), commentHandler.DeleteComment)

	// Social Media
	api.POST("/socialmedias", authService.AuthMiddleware(authService, userService), socialMediaHandler.CreateSocialMedia)
	api.GET("/socialmedias", authService.AuthMiddleware(authService, userService), socialMediaHandler.GetSocialMedia)
	api.PUT("/socialmedias/:id", authService.AuthMiddleware(authService, userService), socialMediaHandler.UpdateSocialMedia)
	api.DELETE("/socialmedias/:id", authService.AuthMiddleware(authService, userService), socialMediaHandler.DeleteSocialMedia)

	err := router.Run(":" + cfg.ServerPort)
	if err != nil {
		return
	}
}
