package models

import (
	"testing"
	m "url_shortener/backend/shortener/internal/models"
)

func TestUrlCreate(t *testing.T) {
	defer PanicHandler(t)

	db := GetConnection(t)
	url_repr := &m.UrlRepresentation{}
	url_repr.UrlLong = "Test"
	err := url_repr.Create(db)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("UrlRepr created: %+v", url_repr)
}
