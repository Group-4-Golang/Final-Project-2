package photo

import (
	"hacktiv-final2/user"
	"time"
)

type Photo struct {
	ID        int
	Title     string
	Caption   string
	PhotoURL  string
	UserID    int
	CreatedAt time.Time
	UpdatedAt time.Time
	User      user.User `gorm:"Constraint:OnDelete:CASCADE;"`
}
