package models

import "time"

type UrlRepresentation struct {
	ID        string `gorm:"column:id"`
	UrlLong   string `gorm:"column:url_long"`
	UrlShort  string `gorm:"column:url_short"`
	CreatedAt time.Time
}
