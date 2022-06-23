package models

import (
	"net/http"
	"time"
	"url_shortener/backend/lib/errs"

	"gorm.io/gorm"
)

type Template struct {
	ID          uint64
	Pan         uint64
	AutoPay     bool       `gorm:"column:autopay"`
	RepeatType  string     `gorm:"column:repeat_type"`
	PayDay      *uint64    `gorm:"column:payday"`
	CountPay    *uint64    `gorm:"column:countpay"`
	EndOfPay    *time.Time `gorm:"column:endofpay"`
	NextPayDate *time.Time `gorm:"column:nextpaydate"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const TemplateTable = "template"

func (obj *Template) Create(db *gorm.DB) *errs.Error {
	err := db.Table(TemplateTable).Create(&obj).Error
	if err != nil {
		return errs.New().SetCode(http.StatusInternalServerError).SetMsg("unable to create template")
	}
	return nil
}

func (obj *Template) Read(db *gorm.DB) *errs.Error {
	err := db.Table(TemplateTable).First(&obj).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return errs.New().SetCode(http.StatusNotFound).SetMsg("template not found")
		}
		return errs.New().SetCode(http.StatusInternalServerError).SetMsg("unable to create template")
	}
	return nil
}

func (obj *Template) Update(db *gorm.DB) *errs.Error {
	err := db.Table(TemplateTable).Save(&obj).Error
	if err != nil {
		return errs.New().SetCode(http.StatusInternalServerError).SetMsg("unable to create template")
	}
	return nil
}

func LoadTemplateById(db *gorm.DB, id uint64) (*Template, *errs.Error) {
	tmpl := &Template{}
	err := db.Table(TemplateTable).Where("id=?", id).Scan(&tmpl).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.New().SetCode(http.StatusNotFound).SetMsg("template not found")
		}
		return nil, errs.New().SetCode(http.StatusInternalServerError).SetMsg("unable to load template")
	}
	return tmpl, nil
}

func DeleteTemplateById(db *gorm.DB, id uint64) *errs.Error {
	tx := db.Begin()
	err := tx.Table(TemplateTable).Delete(&Template{}, id).Error
	if err != nil {
		tx.Rollback()
		if err == gorm.ErrRecordNotFound {
			return errs.New().SetCode(http.StatusNotFound).SetMsg("template not found")
		}
		return errs.New().SetCode(http.StatusInternalServerError).SetMsg("unable to load template")
	}
	tx.Commit()
	return nil
}
