package comment

import (
	"gorm.io/gorm"
)

type Repository interface {
	FindByID(id int) (Comment, error)
	FindAll(comments *[]Comment) error
	Store(comment Comment) (Comment, error)
	Update(comment Comment) (Comment, error)
	Delete(comment Comment) error
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindByID(id int) (Comment, error) {
	var comment Comment
	err := r.db.Preload("User").Preload("Photo").Where("id = ?", id).Find(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) FindAll(comments *[]Comment) error {
	err := r.db.Preload("User").Preload("Photo").Find(comments).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Store(comment Comment) (Comment, error) {
	err := r.db.Preload("User").Preload("Photo").Create(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) Update(comment Comment) (Comment, error) {
	err := r.db.Preload("User").Preload("Photo").Save(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, nil
}

func (r *repository) Delete(comment Comment) error {
	err := r.db.Preload("User").Preload("Photo").Delete(&comment).Error
	if err != nil {
		return err
	}
	return nil
}
