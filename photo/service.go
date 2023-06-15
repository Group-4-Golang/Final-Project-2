package photo

import "errors"

type Service interface {
	FindByID(ID int, UserID int) (Photo, error)
	FindAll(UserID int) ([]Photo, error)
	Store(input Input) (Photo, error)
	Update(ID int, input Input) (Photo, error)
	Delete(ID int, UserID int) (Photo, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByID(ID int, UserID int) (Photo, error) {
	photo, err := s.repository.FindByID(ID)
	if err != nil {
		return photo, err
	}

	if photo.UserID != UserID {
		return photo, errors.New("unauthorized")
	}

	return photo, nil
}

func (s *service) FindAll(UserID int) ([]Photo, error) {
	var photos []Photo
	err := s.repository.FindAll(&photos)
	if err != nil {
		return photos, err
	}

	if len(photos) == 0 {
		return photos, errors.New("no photo found")
	}

	var filteredPhotos []Photo
	for _, photo := range photos {
		if photo.UserID == UserID {
			filteredPhotos = append(filteredPhotos, photo)
		}
	}

	return filteredPhotos, nil
}

func (s *service) Store(input Input) (Photo, error) {
	photo := Photo{}
	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoURL = input.PhotoURL
	photo.UserID = input.User.ID

	newPhoto, err := s.repository.Store(photo)
	if err != nil {
		return newPhoto, err
	}

	return newPhoto, nil
}

func (s *service) Update(ID int, input Input) (Photo, error) {
	photo, err := s.repository.FindByID(ID)
	if err != nil {
		return photo, err
	}

	if photo.UserID != input.User.ID {
		return photo, errors.New("unauthorized")
	}

	photo.Title = input.Title
	photo.Caption = input.Caption
	photo.PhotoURL = input.PhotoURL

	updatedPhoto, err := s.repository.Update(photo)
	if err != nil {
		return updatedPhoto, err
	}

	return updatedPhoto, nil
}

func (s *service) Delete(ID int, UserID int) (Photo, error) {
	photo, err := s.repository.FindByID(ID)
	if err != nil {
		return photo, err
	}

	if photo.UserID != UserID {
		return photo, errors.New("unauthorized")
	}

	err = s.repository.Delete(photo)
	if err != nil {
		return photo, err
	}

	return photo, nil
}
