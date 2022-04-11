package api

import (
	"net/http"
	"url_shortener/backend/lib/errs"
	"url_shortener/backend/lib/hashid"
	r "url_shortener/backend/lib/request"
	"url_shortener/backend/shortener/internal/models"
)

type ReqGetShortUrl struct {
	r.Header
	UrlShort string `json:"url_short"`
}

type RplGetShortUrl struct {
	r.Header
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

func (obj *ReqGetShortUrl) Execute() (r.Reply, *errs.Error) {
	rpl := &RplGetShortUrl{}
	db, log := obj.DB(), obj.Log()

	log.SetID(hashid.NewUUID()).SetMethod("url/long/get")

	log.Infof("%+v", obj)

	url, err := models.LoadUrlReprById(db, obj.UrlShort)
	if err != nil {
		return nil, err
	}

	rpl.UrlLong = url.UrlLong

	return rpl, nil
}
