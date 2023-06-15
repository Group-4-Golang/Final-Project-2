package user

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	Update(ID int, input LoginUserInput) (User, error)
	Delete(ID int) (User, error)
	FindByID(userID int) (User, error)
	FindByUsername(username string) (User, error)
	FindByEmail(email string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Register(input RegisterUserInput) (User, error) {
	user := User{}
	user.Username = input.Username
	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Age = input.Age

	newUser, err := s.repository.Store(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}

func (s *service) Login(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, errors.New("wrong password")
	}

	return user, nil
}

func (s *service) Update(ID int, input LoginUserInput) (User, error) {
	user, err := s.repository.FindByID(ID)
	if err != nil {
		return user, err
	}

	if user.Email != input.Email {
		return user, errors.New("unauthorized")
	}

	user.Email = input.Email

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)

	updatedUser, err := s.repository.Update(user)
	if err != nil {
		return updatedUser, err
	}

	return updatedUser, nil
}

func (s *service) Delete(ID int) (User, error) {
	user, err := s.repository.FindByID(ID)

	if user.ID != ID {
		return user, errors.New("unauthorized")
	}

	if err != nil {
		return user, err
	}
	err = s.repository.Delete(user)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) FindByID(userID int) (User, error) {
	user, err := s.repository.FindByID(userID)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) FindByUsername(username string) (User, error) {
	user, err := s.repository.FindByUsername(username)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *service) FindByEmail(email string) (User, error) {
	user, err := s.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}
	return user, nil
}
