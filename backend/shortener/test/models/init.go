package models

import (
	"testing"
	u "url_shortener/backend/lib"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetConnection(t *testing.T) *gorm.DB {
	config := u.NewConfig("../../cmd/config.yml")
	db, err := gorm.Open(postgres.Open(config.GetDBConnection()), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	return db
}

func PanicHandler(t *testing.T) {
	if err := recover(); err != nil {
		t.Fatal(err)
	}
}
