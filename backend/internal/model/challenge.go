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
	ID              uint           `gorm:"primaryKey" json:"id"`
	Title           string         `gorm:"size:100" json:"title"`
	Description     string         `gorm:"type:text" json:"description"`
	Category        string         `gorm:"size:20" json:"category"` // Web, Pwn, Crypto...
	Points          int            `gorm:"default:100" json:"points"`
	Difficulty      string         `gorm:"size:10" json:"difficulty"`          // Easy, Medium, Hard
	Type            string         `gorm:"size:20" json:"type"`                // attachment | static | dynamic
	ContainerConfig datatypes.JSON `gorm:"type:jsonb" json:"container_config"` // {"image": "nginx:alpine", "exposed_port": 80}
	Flag            string         `gorm:"size:255" json:"-"`                  // Static flag (for static challenges)
	Attachments     datatypes.JSON `gorm:"type:jsonb" json:"attachments"`      // File attachments (for attachment challenges)
	Metadata        datatypes.JSON `gorm:"type:jsonb" json:"metadata"`         // {"author": "admin", "tags": ["easy"]}
	SolvedCount     int            `gorm:"default:0" json:"solved_count"`
	CreatedAt       time.Time      `json:"created_at"`
	UpdatedAt       time.Time      `json:"updated_at"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}

// ContainerConfig represents the Docker container configuration
type ContainerConfig struct {
	Image       string   `json:"image"`             // Docker image name
	ExposedPort int      `json:"exposed_port"`      // Port exposed by the container
	Command     string   `json:"command,omitempty"` // Optional command to run
	Env         []string `json:"env,omitempty"`     // Environment variables
}

// Attachment represents a file attachment
type Attachment struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}
