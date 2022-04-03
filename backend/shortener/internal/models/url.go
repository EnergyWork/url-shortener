package models

import (
	"net/http"
	"time"
	"url_shortener/backend/lib/errs"

	"gorm.io/gorm"
)

type UrlRepresentation struct {
	ID        uint64
	UrlLong   string
	UrlShort  string
	CreatedAt time.Time
}

const UrlRepresentationTable = "url_representation"

func (obj *UrlRepresentation) Create(db *gorm.DB) *errs.Error {
	res := db.Table(UrlRepresentationTable).Create(obj)
	if res.Error != nil {
		return errs.NewError().SetCode(http.StatusInternalServerError).SetMsg(res.Error.Error())
	}
	return nil
}
