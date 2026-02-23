package model

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// User represents a user in the system
type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Username  string         `gorm:"uniqueIndex;size:50" json:"username"`
	Password  string         `gorm:"size:255" json:"-"` // bcrypt hash, never expose
	Email     string         `gorm:"size:100" json:"email"`
	Role      string         `gorm:"default:'user'" json:"role"` // 'admin' | 'user'
	Score     int            `gorm:"default:0" json:"score"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

// HashPassword hashes the password using bcrypt
func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

// CheckPassword compares the provided password with the hashed password
func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

// BeforeCreate hook to hash password if changed
func (u *User) BeforeCreate(tx *gorm.DB) error {
	if u.Password != "" {
		if err := u.HashPassword(u.Password); err != nil {
			return err
		}
	}
	return nil
}
