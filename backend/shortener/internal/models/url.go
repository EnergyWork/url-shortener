package models

import (
	"fmt"
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
		return errs.New().SetCode(http.StatusInternalServerError).SetMsg(res.Error.Error())
	}
	return nil
}

func LoadUrlReprById(db *gorm.DB, urlShort string) (*UrlRepresentation, *errs.Error) {
	tmp := &UrlRepresentation{}
	res := db.Table(UrlRepresentationTable).Where("url_short = ?", urlShort).Scan(tmp)
	switch res.Error {
	case gorm.ErrRecordNotFound:
		return nil, errs.New().SetCode(http.StatusInternalServerError).SetMsg(fmt.Sprintf("record nor found id: %s", urlShort))
	case nil:
	default:
		return nil, errs.New().SetCode(http.StatusInternalServerError).SetMsg(fmt.Sprintf("unable to load url repr by id: %s", urlShort))
	}
	return tmp, nil
}
