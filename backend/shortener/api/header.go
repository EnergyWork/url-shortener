package api

import (
	"url_shortener/backend/lib"
	"url_shortener/backend/lib/errs"

	"gorm.io/gorm"
)

type Header struct {
	Meta
	Error *errs.Error
}

type Meta struct {
	Request   string
	Signature string
	Digest    string
	DB        *gorm.DB
	Log       *lib.Logger
	// PrivateKey
}
