package database

import (
	"time"

	"gorm.io/gorm"
)

type Sprint struct {
	gorm.Model
	Title       string `gorm:"unique"`
	StartDate   time.Time
	EndDate     time.Time
	Velocity    *float32
	ProjectID   uint        // 1:n (Project:Sprint)
	UserStories []UserStory // 1:n (Sprint:UserStory)
}
