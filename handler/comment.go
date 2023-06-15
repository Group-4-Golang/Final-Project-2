package handler

import (
	"hacktiv-final2/auth"
	"hacktiv-final2/comment"
	"hacktiv-final2/helper"
	"hacktiv-final2/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type commentHandler struct {
	commentService comment.Service
	authService    auth.Service
}

func NewCommentHandler(commentService comment.Service, authService auth.Service) *commentHandler {
	return &commentHandler{commentService, authService}
}

func (h *commentHandler) CreateComment(c *gin.Context) {
	var input comment.Input
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	input.User = currentUser

	newComment, err := h.commentService.Store(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := comment.FormatComment(newComment)
	c.JSON(http.StatusOK, formatter)
}

func (h *commentHandler) GetComment(c *gin.Context) {
	currentUser := c.MustGet("currentUser").(user.User)

	comments, err := h.commentService.FindAll(currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := comment.FormatGetAllComment(comments)
	c.JSON(http.StatusOK, formatter)
}

func (h *commentHandler) UpdateComment(c *gin.Context) {
	var inputID comment.GetCommentInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"errors": "ID not found"}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	var inputData comment.UpdateInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		c.JSON(http.StatusUnprocessableEntity, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)
	inputData.User = currentUser

	comments, err := h.commentService.Update(inputID.ID, inputData)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := comment.FormatComment(comments)
	c.JSON(http.StatusOK, formatter)
}

func (h *commentHandler) DeleteComment(c *gin.Context) {
	var inputID comment.GetCommentInput
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		errorMessage := gin.H{"errors": "ID not found"}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	currentUser := c.MustGet("currentUser").(user.User)

	_, err = h.commentService.Delete(inputID.ID, currentUser.ID)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		c.JSON(http.StatusBadRequest, errorMessage)
		return
	}

	formatter := comment.FormatDeleteComment()
	c.JSON(http.StatusOK, formatter)
}
