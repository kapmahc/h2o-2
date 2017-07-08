package nut

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Timestamp timestamp
type Timestamp struct {
	ID        uint      `orm:"column(id)"`
	CreatedAt time.Time `orm:"auto_now_add"`
}

// Model model
type Model struct {
	Timestamp
	UpdatedAt time.Time `orm:"auto_now"`
}

// Media media
type Media struct {
	Model
	Body string
	Type string
}

// Setting setting
type Setting struct {
	Model

	Key    string
	Val    string
	Encode bool
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
