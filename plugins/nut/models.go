package nut

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Model model
type Model struct {
	ID        uint      `orm:"column(id)"`
	CreatedAt time.Time `orm:"auto_now_add"`
	UpdatedAt time.Time `orm:"auto_now"`
}

// Setting setting
type Setting struct {
	Model

	Key    string
	Val    string
	Encode string
}

// TableName table name
func (u *Setting) TableName() string {
	return "settings"
}

// Locale locale
type Locale struct {
	Model

	Lang    string
	Code    string
	Message string
}

// TableName table name
func (u *Locale) TableName() string {
	return "locales"
}
func init() {
	orm.RegisterModel(new(Locale), new(Setting))
}
