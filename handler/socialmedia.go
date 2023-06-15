package handler

import (
	"hacktiv-final2/auth"
	"hacktiv-final2/helper"
	"hacktiv-final2/socialmedia"
	"hacktiv-final2/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type socialMediaHandler struct {
	socialMediaService socialmedia.Service
	authService        auth.Service
}

func NewSocialMediaHandler(socialMediaService socialmedia.Service, authService auth.Service) *socialMediaHandler {
	return &socialMediaHandler{socialMediaService, authService}
}

func (h *socialMediaHandler) CreateSocialMedia(c *gin.Context) {
	var input socialmedia.Input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newSocialMedia, err := h.socialMediaService.Store(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := socialmedia.FormatSocialMedia(newSocialMedia)
	c.JSON(http.StatusOK, formatter)
}

func (h *socialMediaHandler) GetSocialMedia(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	socialMedia, err := h.socialMediaService.FindAll(currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := socialmedia.FormatSocialMedias(socialMedia)
	c.JSON(http.StatusOK, formatter)
}

func (h *socialMediaHandler) UpdateSocialMedia(c *gin.Context) {
	var inputID socialmedia.GetSocialMediaInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	var inputData socialmedia.Input
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedSocialMedia, err := h.socialMediaService.Update(inputID.ID, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := socialmedia.FormatSocialMedia(updatedSocialMedia)
	c.JSON(http.StatusOK, formatter)
}

func (h *socialMediaHandler) DeleteSocialMedia(c *gin.Context) {
	var inputID socialmedia.GetSocialMediaInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	UserID := currentUser.ID

	_, err = h.socialMediaService.Delete(inputID.ID, UserID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := socialmedia.FormatDeleteSocialMedia()
	c.JSON(http.StatusOK, formatter)
}
