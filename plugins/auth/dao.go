package auth

import (
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/h2o/plugins/nut"
)

// GetUserByEmail get user by email
func GetUserByEmail(email string) (*User, error) {
	var it User
	err := orm.NewOrm().QueryTable(&it).Filter("email", email).One(&it)
	return &it, err
}

// GetUserByUID get user by uid
func GetUserByUID(uid string) (*User, error) {
	var it User
	err := orm.NewOrm().QueryTable(&it).Filter("uid", uid).One(&it)
	return &it, err
}

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

// ConfirmUser confirm user
func ConfirmUser(id uint, ip, lang string) error {
	o := orm.NewOrm()
	if _, err := o.QueryTable(new(User)).Filter("id", id).Update(orm.Params{
		"confirmed_at": time.Now(),
	}); err != nil {
		return err
	}

	return AddLog(id, ip, lang, "auth.logs.confirm")
}

// AddEmailUser add user by email
func AddEmailUser(email, password, ip, lang string) (*User, error) {
	o := orm.NewOrm()

	cnd := orm.NewCondition()

	count, err := o.QueryTable(new(User)).
		SetCond(cnd.And("email", email).OrCond(cnd.And("provider_type", UserTypeEmail).And("provider_id", email))).
		Count()
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, nut.E(lang, "auth.errors.email-already-exist")
	}

	user := User{
		Email:        email,
		Password:     string(nut.Sum([]byte(password))),
		ProviderType: UserTypeEmail,
		ProviderID:   email,
	}
	user.SetUID()
	user.SetGravatarLogo()

	if _, err = o.Insert(&user); err != nil {
		return nil, err
	}
	err = AddLog(user.ID, ip, lang, "auth.logs.sign-up")
	return &user, err
}

// AddLog add log
func AddLog(user uint, ip, lang, code string, args ...interface{}) error {
	_, err := orm.NewOrm().Insert(&Log{
		UserID:  user,
		IP:      &ip,
		Message: nut.T(lang, code, args...),
	})
	return err
}
