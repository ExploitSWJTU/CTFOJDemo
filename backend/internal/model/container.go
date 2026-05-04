package model

import (
	"fmt"
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

// ContainerStatus represents the status of a container
type ContainerStatus string

const (
	ContainerStatusRunning  ContainerStatus = "running"
	ContainerStatusStopped  ContainerStatus = "stopped"
	ContainerStatusExpired  ContainerStatus = "expired"
	ContainerStatusCreating ContainerStatus = "creating"
	ContainerStatusFailed   ContainerStatus = "failed"
)

// Container represents a running Docker container instance for a challenge
type Container struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	UserID      uint           `gorm:"index" json:"user_id"`
	ChallengeID uint           `gorm:"index" json:"challenge_id"`
	ContainerID string         `gorm:"size:100" json:"container_id"`
	PortMapping datatypes.JSON `gorm:"type:jsonb" json:"port_mapping"`
	HostPort    int            `gorm:"index" json:"host_port"`
	AccessHost  string         `gorm:"size:255" json:"access_host"`
	Flag        string         `gorm:"size:255" json:"-"`
	Status      string         `gorm:"size:20;default:'running'" json:"status"`
	LastError   string         `gorm:"type:text" json:"last_error,omitempty"`
	StartedAt   *time.Time     `json:"started_at,omitempty"`
	ExpiresAt   time.Time      `json:"expires_at"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`

	// Associations
	User      User      `gorm:"foreignKey:UserID" json:"-"`
	Challenge Challenge `gorm:"foreignKey:ChallengeID" json:"-"`
}

// PortMapping represents the port mapping of a container
type PortMapping struct {
	ContainerPort int `json:"container_port"` // Port inside container
	HostPort      int `json:"host_port"`      // Port on host machine
}

// IsExpired checks if the container has expired
func (c *Container) IsExpired() bool {
	return time.Now().After(c.ExpiresAt)
}

// TimeLeft returns the remaining time as a formatted string
func (c *Container) TimeLeft() string {
	if c.IsExpired() {
		return "00:00:00"
	}

	duration := time.Until(c.ExpiresAt)
	hours := int(duration.Hours())
	minutes := int(duration.Minutes()) % 60
	seconds := int(duration.Seconds()) % 60

	return fmt.Sprintf("%02d:%02d:%02d", hours, minutes, seconds)
}
