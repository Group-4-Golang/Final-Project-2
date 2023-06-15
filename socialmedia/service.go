package socialmedia

import "errors"

type Service interface {
	FindByID(ID int, UserID int) (SocialMedia, error)
	FindAll(UserID int) ([]SocialMedia, error)
	Store(input Input) (SocialMedia, error)
	Update(ID int, input Input) (SocialMedia, error)
	Delete(ID int, UserID int) (SocialMedia, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindByID(ID int, UserID int) (SocialMedia, error) {
	socialMedia, err := s.repository.FindByID(ID)
	if err != nil {
		return socialMedia, err
	}

	if socialMedia.UserID != UserID {
		return socialMedia, errors.New("unauthorized")
	}

	return socialMedia, nil
}

func (s *service) FindAll(UserID int) ([]SocialMedia, error) {
	var socialMedias []SocialMedia
	err := s.repository.FindAll(&socialMedias)
	if err != nil {
		return socialMedias, err
	}

	if len(socialMedias) == 0 {
		return socialMedias, errors.New("no social media found")
	}

	var filteredSocialMedias []SocialMedia
	for _, socialMedia := range socialMedias {
		if socialMedia.UserID == UserID {
			filteredSocialMedias = append(filteredSocialMedias, socialMedia)
		}
	}

	return filteredSocialMedias, nil
}

func (s *service) Store(input Input) (SocialMedia, error) {
	socialMedia := SocialMedia{}
	socialMedia.Name = input.Name
	socialMedia.SocialMediaURL = input.SocialMediaURL
	socialMedia.UserID = input.User.ID

	newSocialMedia, err := s.repository.Store(socialMedia)
	if err != nil {
		return newSocialMedia, err
	}

	return newSocialMedia, nil
}

func (s *service) Update(ID int, input Input) (SocialMedia, error) {
	socialMedia, err := s.repository.FindByID(ID)
	if err != nil {
		return socialMedia, err
	}

	if socialMedia.UserID != input.User.ID {
		return socialMedia, errors.New("unauthorized")
	}

	socialMedia.Name = input.Name
	socialMedia.SocialMediaURL = input.SocialMediaURL
	updatedSocialMedia, err := s.repository.Update(socialMedia)
	if err != nil {
		return updatedSocialMedia, err
	}
	return updatedSocialMedia, nil
}

func (s *service) Delete(ID int, UserID int) (SocialMedia, error) {
	socialMedia, err := s.repository.FindByID(ID)
	if err != nil {
		return socialMedia, err
	}

	if socialMedia.UserID != UserID {
		return socialMedia, errors.New("unauthorized")
	}

	err = s.repository.Delete(socialMedia)
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}
