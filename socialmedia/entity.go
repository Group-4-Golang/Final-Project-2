package socialmedia

import (
	"hacktiv-final2/user"
	"time"
)

type SocialMedia struct {
	ID             int
	Name           string
	SocialMediaURL string
	UserID         int
	CreatedAt      time.Time
	UpdatedAt      time.Time
	User           user.User `gorm:"Constraint:OnDelete:CASCADE;"`
	
}
