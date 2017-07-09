package nut

import (
	"encoding/json"

	"github.com/astaxie/beego/orm"
)

// Set set setting key-val
func Set(key string, val interface{}, enc bool) error {
	buf, err := json.Marshal(val)
	if err != nil {
		return err
	}
	if enc {
		if buf, err = Encrypt(buf); err != nil {
			return err
		}
	}

	o := orm.NewOrm()
	var it Setting
	err = o.QueryTable(&it).Filter("key", key).One(&it, "id")
	switch err {
	case nil:
		_, err = o.QueryTable(&it).Filter("id", it.ID).Update(orm.Params{"val": string(buf), "encode": enc})
	case orm.ErrNoRows:
		_, err = o.Insert(&Setting{Key: key, Val: string(buf), Encode: enc})
	}
	return err
}

// Get get setting key=>val
func Get(key string, val interface{}) error {
	o := orm.NewOrm()
	var it Setting
	err := o.QueryTable(&it).Filter("key", key).One(&it, "val")
	if err != nil {
		return err
	}
	buf := []byte(it.Val)
	if it.Encode {
		if buf, err = Decrypt(buf); err != nil {
			return err
		}
	}
	return json.Unmarshal(buf, val)
}

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
