package comment

import (
	"errors"
	"hacktiv-final2/photo"
)

type Service interface {
	FindByID(ID int, UserID int) (Comment, error)
	FindAll(UserID int) ([]Comment, error)
	Store(input Input) (Comment, error)
	Update(ID int, input UpdateInput) (Comment, error)
	Delete(ID int, UserID int) (Comment, error)
}

type service struct {
	repository      Repository
	photoRepository photo.Repository
}

func NewService(repository Repository, photoRepository photo.Repository) *service {
	return &service{repository, photoRepository}
}

func (s *service) FindByID(ID int, UserID int) (Comment, error) {
	comment, err := s.repository.FindByID(ID)
	if err != nil {
		return comment, err
	}

	if comment.UserID != UserID {
		return comment, errors.New("unauthorized")
	}

	return comment, nil
}

func (s *service) FindAll(UserID int) ([]Comment, error) {
	var comments []Comment
	err := s.repository.FindAll(&comments)
	if err != nil {
		return comments, err
	}

	if len(comments) == 0 {
		return comments, errors.New("no comment found")
	}

	var filteredComments []Comment
	for _, comment := range comments {
		if comment.UserID == UserID {
			filteredComments = append(filteredComments, comment)
		}
	}

	return filteredComments, nil
}

func (s *service) Store(input Input) (Comment, error) {
	comment := Comment{}
	comment.Message = input.Message
	comment.PhotoID = input.PhotoID
	comment.UserID = input.User.ID

	if input.User.ID == 0 {
		return comment, errors.New("unauthorized")
	}

	photos, err := s.photoRepository.FindByID(input.PhotoID)
	if err != nil {
		return comment, err
	}

	if photos.ID == 0 {
		return comment, errors.New("photo not found")
	}

	newComment, err := s.repository.Store(comment)
	if err != nil {
		return newComment, err
	}

	return newComment, nil
}

func (s *service) Update(ID int, input UpdateInput) (Comment, error) {
	comment, err := s.repository.FindByID(ID)
	if err != nil {
		return comment, err
	}

	if comment.UserID != input.User.ID {
		return comment, errors.New("unauthorized")
	}

	comment.Message = input.Message

	updatedComment, err := s.repository.Update(comment)
	if err != nil {
		return updatedComment, err
	}

	return updatedComment, nil
}

func (s *service) Delete(ID int, UserID int) (Comment, error) {
	comment, err := s.repository.FindByID(ID)
	if err != nil {
		return comment, err
	}

	if comment.UserID != UserID {
		return comment, errors.New("unauthorized")
	}

	err = s.repository.Delete(comment)
	if err != nil {
		return comment, err
	}

	return comment, nil
}
