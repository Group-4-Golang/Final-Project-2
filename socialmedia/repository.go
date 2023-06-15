package socialmedia

import "gorm.io/gorm"

type Repository interface {
	FindByID(id int) (SocialMedia, error)
	FindAll(socialMedias *[]SocialMedia) error
	Store(socialMedia SocialMedia) (SocialMedia, error)
	Update(socialMedia SocialMedia) (SocialMedia, error)
	Delete(socialMedia SocialMedia) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(id int) (SocialMedia, error) {
	var socialMedia SocialMedia
	err := r.db.Preload("User").Where("ID = ?", id).First(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (r *repository) FindAll(socialMedias *[]SocialMedia) error {
	err := r.db.Preload("User").Find(socialMedias).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Store(socialMedia SocialMedia) (SocialMedia, error) {
	err := r.db.Preload("User").Create(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (r *repository) Update(socialMedia SocialMedia) (SocialMedia, error) {
	err := r.db.Preload("User").Save(&socialMedia).Error
	if err != nil {
		return socialMedia, err
	}
	return socialMedia, nil
}

func (r *repository) Delete(socialMedia SocialMedia) error {
	err := r.db.Preload("User").Delete(&socialMedia).Error
	if err != nil {
		return err
	}
	return nil
}
