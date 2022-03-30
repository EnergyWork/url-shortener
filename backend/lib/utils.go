package lib

import (
	"encoding/json"
	"net/http"
)

func InitHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func RespondError(w http.ResponseWriter, r *http.Request, code int, err error) {
	Respond(w, r, code, map[string]string{"error": err.Error()})
}

func Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	InitHeaders(w)
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}
