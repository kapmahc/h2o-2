package nut

import (
	"encoding/json"
	"time"

	"github.com/astaxie/beego/orm"
)

func getRole(role, rty string, rid uint) (*Role, error) {
	var it Role
	o := orm.NewOrm()
	err := o.QueryTable(&it).Filter("name", role).Filter("resource_type", rty).Filter("resource_id", rid).One(&it, "id")
	return &it, err
}

func getPolicy(user, role uint) (*Policy, error) {
	var it Policy
	o := orm.NewOrm()
	err := o.QueryTable(&it).Filter("user_id", user).Filter("role_id", role).One(&it, "id")
	return &it, err
}

// Allow allow role to user
func Allow(user uint, role, rty string, rid uint, years, months, days int) error {
	o := orm.NewOrm()
	ro, err := getRole(role, rty, rid)
	switch err {
	case orm.ErrNoRows:
		ro.Name = role
		ro.ResourceID = rid
		ro.ResourceType = rty
		if _, err = o.Insert(ro); err != nil {
			return err
		}
	case nil:
		break
	default:
		return err
	}

	begin := time.Now()
	end := begin.AddDate(years, months, days)
	pl, err := getPolicy(user, ro.ID)
	switch err {
	case nil:
		_, err = o.QueryTable(pl).Filter("id", ro.ID).Update(orm.Params{
			"start_up":  begin,
			"shut_down": end,
		})
	case orm.ErrNoRows:
		_, err = o.Insert(&Policy{
			UserID:   user,
			RoleID:   ro.ID,
			StartUp:  begin,
			ShutDown: end,
		})
	}

	return err
}

// Deny deny role from user@resource
func Deny(user uint, role, rty string, rid uint, days uint) error {
	ro, err := getRole(role, rty, rid)
	if err != nil {
		return err
	}
	pl, err := getPolicy(user, ro.ID)
	if err != nil {
		return err
	}
	o := orm.NewOrm()
	_, err = o.Delete(pl)
	return err
}

// Can check policy
func Can(user uint, role, rty string, rid uint) bool {
	ro, err := getRole(role, rty, rid)
	if err != nil {
		return false
	}
	pl, err := getPolicy(user, ro.ID)
	if err != nil {
		return false
	}
	return pl.Enable()
}

// AddEmailUser add user by email
func AddEmailUser(email, password string) (*User, error) {

	user := User{
		Email:        email,
		Password:     string(Sum([]byte(password))),
		ProviderType: UserTypeEmail,
		ProviderID:   email,
	}
	user.SetUID()
	user.SetGravatarLogo()

	if _, err := orm.NewOrm().Insert(&user); err != nil {
		return nil, err
	}
	return &user, nil
}

// AddLog add log
func AddLog(user uint, ip, lang, code string, args ...interface{}) error {
	_, err := orm.NewOrm().Insert(&Log{
		UserID:  user,
		IP:      ip,
		Message: T(lang, code, args...),
	})
	return err
}

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
func Get(key string, val interface{}, enc bool) error {
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
