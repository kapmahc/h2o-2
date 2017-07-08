package nut

import "github.com/astaxie/beego/orm"

// GetMessage get locale
func GetMessage(lang, code string) (string, error) {
	o := orm.NewOrm()
	var it Locale
	err := o.QueryTable(&it).Filter("code", code).Filter("lang", lang).One(&it, "message")
	return it.Message, err
}

// SetMessage Set locale
func SetMessage(lang, code, message string) error {
	o := orm.NewOrm()
	var it Locale

	err := o.QueryTable(&it).Filter("code", code).Filter("lang", lang).One(&it, "id", "message")
	switch err {
	case nil:
		_, err = o.QueryTable(&it).Filter("id", it.ID).Update(orm.Params{"message": message})
	case orm.ErrNoRows:
		_, err = o.Insert(&Locale{Lang: lang, Code: code, Message: message})
	}
	return err
}
