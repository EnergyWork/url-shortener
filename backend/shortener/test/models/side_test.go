package models

import (
	"log"
	"testing"
	"time"
	"url_shortener/backend/shortener/internal/models"
)

func TestCreate(t *testing.T) {
	defer PanicHandler(t)
	db := GetConnection(t).Debug()

	var payday uint64 = 15
	var countpay uint64 = 5
	var endofpay = time.Now().AddDate(0, 4, -3)
	var nextpaydate = time.Now().AddDate(0, 1, 1)

	tmpl := models.Template{
		Pan:         556622117733,
		RepeatType:  "once_a_month",
		AutoPay:     true,
		PayDay:      &payday,
		CountPay:    nil,          //sql.NullInt64{Int64: 5, Valid: false},
		EndOfPay:    nil,          //time.Now().AddDate(0, 4, -3), //sql.NullTime{Time: time.Now().AddDate(0, 5, -3), Valid: true},
		NextPayDate: &nextpaydate, //time.Now().AddDate(0, 0, 1),  //sql.NullTime{Time: time.Now().AddDate(0, 0, 1), Valid: true},
	}

	if err := tmpl.Create(db); err != nil {
		log.Fatal(err)
	}
	log.Printf("TMPL: %+v", tmpl)
	////
	t.Log(countpay, endofpay, nextpaydate, payday)
}

func TestRead(t *testing.T) {
	defer PanicHandler(t)
	db := GetConnection(t).Debug()

	tmpl, err := models.LoadTemplateById(db, 4)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("TMPL: %+v", tmpl)
}

func TestUpdate(t *testing.T) {
	defer PanicHandler(t)
	db := GetConnection(t).Debug()

	tmpl, err := models.LoadTemplateById(db, 12)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("BEFORE UPDATE: %+v", tmpl)

	var cp uint64 = 5
	var t1 = time.Now().AddDate(0, 1, 5)

	tmpl.CountPay = &cp //nil //sql.NullInt64{Int64: 10, Valid: true}
	tmpl.EndOfPay = nil //&t1 //sql.NullTime{Time: time.Now().AddDate(0, 1, 0), Valid: false}

	err = tmpl.Update(db)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("AFTER UPDATE: %+v", tmpl)
	t.Log(cp, t1)
}
