package lib

import (
	api "url_shortener/backend/shortener/api"

	"github.com/julienschmidt/httprouter"
)

func GetDefaultRouter() *httprouter.Router {
	r := httprouter.New()
	r.POST("/create", api.RequestCreateShortUrl)
	return r
}
