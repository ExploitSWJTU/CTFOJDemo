package model

import (
	"time"

	"gorm.io/gorm"
)

// Submission records a user's flag submission for a challenge.
type Submission struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	UserID        uint           `gorm:"index" json:"user_id"`
	ChallengeID   uint           `gorm:"index" json:"challenge_id"`
	SubmittedFlag string         `gorm:"size:255" json:"submitted_flag"`
	IsCorrect     bool           `gorm:"default:false" json:"is_correct"`
	IsSolve       bool           `gorm:"default:false" json:"is_solve"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
}
