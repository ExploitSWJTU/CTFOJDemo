package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ChallengeType represents the type of challenge
type ChallengeType string

const (
	ChallengeTypeAttachment ChallengeType = "attachment" // Static files, no container
	ChallengeTypeStatic     ChallengeType = "static"     // Pre-built container, fixed flag
	ChallengeTypeDynamic    ChallengeType = "dynamic"    // Per-user container, unique flag
)

// ChallengeCategory represents the category of challenge
type ChallengeCategory string

const (
	CategoryWeb        ChallengeCategory = "Web"
	CategoryPwn        ChallengeCategory = "Pwn"
	CategoryCrypto     ChallengeCategory = "Crypto"
	CategoryMisc       ChallengeCategory = "Misc"
	CategoryReverse    ChallengeCategory = "Reverse"
	CategoryMobile     ChallengeCategory = "Mobile"
	CategoryBlockchain ChallengeCategory = "Blockchain"
	CategoryAI         ChallengeCategory = "AI"
)

// ChallengeDifficulty represents the difficulty level
type ChallengeDifficulty string

const (
	DifficultyEasy   ChallengeDifficulty = "Easy"
	DifficultyMedium ChallengeDifficulty = "Medium"
	DifficultyHard   ChallengeDifficulty = "Hard"
)

// Challenge represents a CTF challenge
type Challenge struct {
	ID                       uint           `gorm:"primaryKey" json:"id"`
	Title                    string         `gorm:"size:100" json:"title"`
	Description              string         `gorm:"type:text" json:"description"`
	Category                 string         `gorm:"size:20" json:"category"`
	Points                   int            `gorm:"default:100" json:"points"`
	Difficulty               string         `gorm:"size:10" json:"difficulty"`
	Type                     string         `gorm:"size:30;default:'web_static_flag'" json:"type"`
	Image                    string         `gorm:"size:255" json:"image"`
	InternalPort             int            `gorm:"default:80" json:"internal_port"`
	ExpectedFlag             string         `gorm:"size:255" json:"-"`
	ContainerDurationSeconds int            `gorm:"default:3600" json:"container_duration_seconds"`
	IsVisible                bool           `gorm:"default:true" json:"is_visible"`
	ContainerConfig          datatypes.JSON `gorm:"type:jsonb" json:"container_config"`
	Flag                     string         `gorm:"size:255" json:"-"`
	Attachments              datatypes.JSON `gorm:"type:jsonb" json:"attachments"`
	Metadata                 datatypes.JSON `gorm:"type:jsonb" json:"metadata"`
	SolvedCount              int            `gorm:"default:0" json:"solved_count"`
	CreatedAt                time.Time      `json:"created_at"`
	UpdatedAt                time.Time      `json:"updated_at"`
	DeletedAt                gorm.DeletedAt `gorm:"index" json:"-"`
}

// ContainerConfig represents the Docker container configuration
type ContainerConfig struct {
	Image       string   `json:"image"`             
	ExposedPort int      `json:"exposed_port"`      
	Command     string   `json:"command,omitempty"` 
	Env         []string `json:"env,omitempty"`     
}

// Attachment represents a file attachment
type Attachment struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
