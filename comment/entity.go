package comment

import (
	"hacktiv-final2/photo"
	"hacktiv-final2/user"
	"time"
)

type Comment struct {
	ID        int
	UserID    int
	PhotoID   int
	Message   string
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User   `gorm:"Constraint:OnDelete:CASCADE;"`
	Photo     photo.Photo `gorm:"Constraint:OnDelete:CASCADE;"`
}
