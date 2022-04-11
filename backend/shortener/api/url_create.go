package api

import (
	"net/http"
	"regexp"
	"url_shortener/backend/lib/errs"
	"url_shortener/backend/lib/hashid"
	r "url_shortener/backend/lib/request"
	"url_shortener/backend/shortener/internal/models"
)

type ReqCreateShortUrl struct {
	r.Header
	UrlLong string `json:"url_long"`
}

type RplCreateShortUrl struct {
	r.Header
	UrlShort string `json:"url_short"`
}

func (obj *ReqCreateShortUrl) Authorize() *errs.Error {
	return nil
}

func (obj *ReqCreateShortUrl) Validate() *errs.Error {

	if obj.UrlLong == "" {
		return errs.New().SetCode(http.StatusBadRequest).SetMsg("Validate ERROR: UrlLong must be not empty")
	}

	urlPattern := `^((https?|ftp|file):\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	re := regexp.MustCompile(urlPattern)
	if !re.MatchString(obj.UrlLong) {
		return errs.New().SetCode(http.StatusBadRequest).SetMsg("Validate ERROR: UrlLong must be match url pattern")
	}

	return nil
}

func (obj *ReqCreateShortUrl) Execute() (r.Reply, *errs.Error) {
	rpl := &RplCreateShortUrl{}
	db, log := obj.DB(), obj.Log()

	log.SetID(hashid.NewUUID()).SetMethod("url/short/create")
	log.Infof("%+v", obj)
	defer log.Infof("%+v", rpl)

	hashid := hashid.GetHashId()

	url := &models.UrlRepresentation{
		UrlLong:  obj.UrlLong,
		UrlShort: hashid,
	}

	if err := url.Create(db); err != nil {
		return nil, err
	}

	rpl.UrlShort = url.UrlShort

	return rpl, nil
}
