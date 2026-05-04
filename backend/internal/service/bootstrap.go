package service

import (
	"errors"

	"gorm.io/gorm"

	"swjtu-ctf-oj/internal/model"
	"swjtu-ctf-oj/internal/repository"
)

func EnsureDefaultAdmin(db *repository.Database) error {
	username := getEnvOrDefault("SEED_ADMIN_USERNAME", "admin")
	password := getEnvOrDefault("SEED_ADMIN_PASSWORD", "admin123456")
	email := getEnvOrDefault("SEED_ADMIN_EMAIL", "admin@example.com")

	var existing model.User
	err := db.DB.Where("username = ?", username).First(&existing).Error
	if err == nil {
		return nil
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}

	admin := model.User{
		Username: username,
		Password: password,
		Email:    email,
		Role:     "admin",
	}
	return db.DB.Create(&admin).Error
}

func SeedChallengeVisibility(db *repository.Database) error {
	return db.DB.Model(&model.Challenge{}).
		Where("is_visible IS NULL").
		Update("is_visible", true).Error
}
