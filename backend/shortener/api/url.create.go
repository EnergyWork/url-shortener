package api

import (
	"encoding/json"
	"net/http"
	"url_shortener/backend/lib"

	"github.com/julienschmidt/httprouter"
)

func RequestCreateShortUrl(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	req := &requestCreateShortUrl{}
	lib.InitHeaders(w)
	l := lib.NewLogger().SetMethod("RequestCreateShortUrl")
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		l.Error(err)
		lib.Error(w, r, http.StatusBadRequest, err)
	}
	l.Infof("%+v", req)
	lib.Respond(w, r, http.StatusOK, nil)
}

type requestCreateShortUrl struct {
	UrlLong string `json:"url_long"`
}
