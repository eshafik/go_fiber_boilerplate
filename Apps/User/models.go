package User

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	UserName  string    `gorm:"unique" json:"username"`
	FirstName string    `json:"first_name"`
	LastName  *string   `json:"last_name"`
	Email     *string   `json:"email"`
	Password  *string   `json:"password"`
	Phone     *string   `json:"phone"`
	IsActive  bool      `gorm:"default:true" json:"is_active"`
	JoinedAt  time.Time `json:"joined_at"`
}
