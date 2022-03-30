package api

import (
	"net/http"
	"url_shortener/backend/lib"

	"gorm.io/gorm"
)

type ReqCreateShortUrl struct {
	UrlLong string `json:"url_long"`
}

type RplCreateShortUrl struct {
	UrlShort string `json:"url_short"`
}

func (obj *ReqCreateShortUrl) Authorize() *lib.Error {
	return nil
}

func (obj *ReqCreateShortUrl) Validate() *lib.Error {
	if obj.UrlLong == "" {
		return lib.NewError().SetCode(http.StatusBadRequest).SetMsg("Validate ERROR: UrlLong must be not empty")
	}
	return nil
}

func (obj *ReqCreateShortUrl) Execute(db *gorm.DB, log lib.Logger) (*RplCreateShortUrl, *lib.Error) {
	log.SetID("23").SetMethod("ReqCreateShortUrl")
	log.Error("fsdfsdfsf")

	return nil, lib.NewError().SetCode(http.StatusInternalServerError).SetMsg("Execute ERROR: some text about error")
}
