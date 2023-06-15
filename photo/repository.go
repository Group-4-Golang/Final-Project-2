package photo

import "gorm.io/gorm"

type Repository interface {
	FindByID(id int) (Photo, error)
	FindAll(photos *[]Photo) error
	Store(photo Photo) (Photo, error)
	Update(photo Photo) (Photo, error)
	Delete(photo Photo) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(id int) (Photo, error) {
	var photo Photo
	err := r.db.Preload("User").Where("ID = ?", id).First(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *repository) FindAll(photos *[]Photo) error {
	err := r.db.Preload("User").Find(photos).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Store(photo Photo) (Photo, error) {
	err := r.db.Preload("User").Create(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *repository) Update(photo Photo) (Photo, error) {
	err := r.db.Preload("User").Save(&photo).Error
	if err != nil {
		return photo, err
	}
	return photo, nil
}

func (r *repository) Delete(photo Photo) error {
	err := r.db.Preload("User").Delete(&photo).Error
	if err != nil {
		return err
	}
	return nil
}
