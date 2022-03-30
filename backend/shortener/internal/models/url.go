package models

import (
	"time"

	"gorm.io/gorm"
)

type UrlRepresentation struct {
	gorm.Model
	ID        string `gorm:"column:id"`
	UrlLong   string `gorm:"column:url_long"`
	UrlShort  string `gorm:"column:url_short"`
	CreatedAt time.Time
}
