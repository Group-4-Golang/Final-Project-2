package handler

import (
	"hacktiv-final2/auth"
	"hacktiv-final2/helper"
	"hacktiv-final2/photo"
	"hacktiv-final2/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type photoHandler struct {
	photoService photo.Service
	authService  auth.Service
}

func NewPhotoHandler(photoService photo.Service, authService auth.Service) *photoHandler {
	return &photoHandler{photoService, authService}
}

func (h *photoHandler) CreatePhoto(c *gin.Context) {
	var input photo.Input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newPhoto, err := h.photoService.Store(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := photo.FormatPhoto(newPhoto)
	c.JSON(http.StatusOK, formatter)
}

func (h *photoHandler) GetPhoto(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	photos, err := h.photoService.FindAll(currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := photo.FormatGetAllPhoto(photos)
	c.JSON(http.StatusOK, formatter)
}

func (h *photoHandler) UpdatePhoto(c *gin.Context) {
	var inputID photo.GetPhotoInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	var inputData photo.Input
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	updatedPhoto, err := h.photoService.Update(inputID.ID, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := photo.FormatUpdatePhoto(updatedPhoto)
	c.JSON(http.StatusOK, formatter)
}

func (h *photoHandler) DeletePhoto(c *gin.Context) {
	var inputID photo.GetPhotoInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	_, err = h.photoService.Delete(inputID.ID, currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := photo.FormatDeletePhoto()
	c.JSON(http.StatusOK, formatter)
}
