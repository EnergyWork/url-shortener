package api

import (
	"net/http"
	"url_shortener/backend/lib"
	"url_shortener/backend/lib/errs"
	"url_shortener/backend/lib/hashid"
	"url_shortener/backend/shortener/internal/models"

	"gorm.io/gorm"
)

type ReqGetShortUrl struct {
	Meta     string
	UrlShort string `json:"url_short"`
}

type RplGetShortUrl struct {
	UrlLong string `json:"url_long"`
}

func (obj *ReqGetShortUrl) Authorize() *errs.Error {
	return nil
}

func (obj *ReqGetShortUrl) Validate() *errs.Error {
	if obj.UrlShort == "" {
		return errs.New().SetCode(http.StatusBadRequest).SetMsg("Validate ERROR: UrlLong must be not empty")
	}
	return nil
}

func (obj *ReqGetShortUrl) Execute(db *gorm.DB, log lib.Logger) (*RplGetShortUrl, *errs.Error) {
	rpl := &RplGetShortUrl{}
	log.SetID(hashid.NewUUID()).SetMethod("url/long/get")

	log.Infof("%+v", obj)

	url, err := models.LoadUrlReprById(db, obj.UrlShort)
	if err != nil {
		return nil, err
	}

	rpl.UrlLong = url.UrlLong

	return rpl, nil
}
