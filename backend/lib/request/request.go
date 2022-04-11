package request

import (
	"url_shortener/backend/lib"
	"url_shortener/backend/lib/errs"

	"gorm.io/gorm"
)

const (
	RequestHeader   = "X-Request-Name"
	SignatureHeader = "X-Signature"
	DBMeta          = "db"
	LogMeta         = "log"
)

type Request interface {
	SetHeader(string, string)
	SetMeta(*gorm.DB, lib.Logger)
	Authorize() *errs.Error
	Validate() *errs.Error
	Execute() (Reply, *errs.Error)
}

type Reply interface {
}

type Header struct {
	Request   string      `json:",omitempty"`
	Signature string      `json:",omitempty"`
	Digest    string      `json:",omitempty"`
	db        *gorm.DB    //`json:",omitempty"`
	log       *lib.Logger //`json:",omitempty"`
	// and PrivateKey
	Error *errs.Error `json:"error,omitempty"`
}

func (obj *Header) SetMeta(db *gorm.DB, log lib.Logger) {
	obj.db = db
	obj.log = &log
}

func (obj *Header) SetHeader(key, value string) {
	switch key {
	case RequestHeader:
		obj.Request = value
	case SignatureHeader:
		obj.Signature = value
	}
}

func (obj *Header) DB() *gorm.DB {
	return obj.db
}

func (obj *Header) Log() *lib.Logger {
	return obj.log
}
