package api

import (
	"net/http"
	"regexp"
	"url_shortener/backend/lib"
	"url_shortener/backend/lib/errs"
	"url_shortener/backend/lib/hashid"
	"url_shortener/backend/shortener/internal/models"

	"gorm.io/gorm"
)

type ReqCreateShortUrl struct {
	Meta    string
	UrlLong string `json:"url_long"`
}

type RplCreateShortUrl struct {
	UrlShort string `json:"url_short"`
}

func (obj *ReqCreateShortUrl) Authorize() *errs.Error {
	return nil
}

func (obj *ReqCreateShortUrl) Validate() *errs.Error {

	if obj.UrlLong == "" {
		return errs.NewError().SetCode(http.StatusBadRequest).SetMsg("Validate ERROR: UrlLong must be not empty")
	}

	urlPattern := `^((https?|ftp|file):\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$`
	re := regexp.MustCompile(urlPattern)
	if re.MatchString(obj.UrlLong) {
		return errs.NewError().SetCode(http.StatusBadRequest).SetMsg("Validate ERROR: UrlLong must be match url pattern")
	}

	return nil
}

func (obj *ReqCreateShortUrl) Execute(db *gorm.DB, log lib.Logger) (*RplCreateShortUrl, *errs.Error) {
	rpl := &RplCreateShortUrl{}
	log.SetID("23").SetMethod("ReqCreateShortUrl")
	log.Infof("%+v", obj)

	hashid := hashid.GetHashId(obj.UrlLong)

	url := &models.UrlRepresentation{
		UrlLong:  obj.UrlLong,
		UrlShort: hashid,
	}

	if err := url.Create(db); err != nil {
		return nil, err
	}

	return rpl, nil
}
