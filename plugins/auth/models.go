package auth

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/google/uuid"
	"github.com/kapmahc/h2o/plugins/nut"
)

const (
	// RoleAdmin admin role
	RoleAdmin = "admin"
	// RoleRoot root role
	RoleRoot = "root"
	// UserTypeEmail email user
	UserTypeEmail = "email"

	// DefaultResourceType default resource type
	DefaultResourceType = "-"
	// DefaultResourceID default resourc id
	DefaultResourceID = 0
)

// User user
type User struct {
	nut.Model

	Name            string
	Email           string
	UID             string `orm:"column(uid)"`
	Password        string
	ProviderID      string `orm:"column(provider_id)"`
	ProviderType    string
	Home            string
	Logo            string
	SignInCount     uint
	LastSignInAt    *time.Time
	LastSignInIP    *string `orm:"column(last_sign_in_ip)"`
	CurrentSignInAt *time.Time
	CurrentSignInIP *string `orm:"column(current_sign_in_ip)"`
	ConfirmedAt     *time.Time
	LockedAt        *time.Time
}

// TableName table name
func (*User) TableName() string {
	return "users"
}

// IsConfirm is confirm?
func (p *User) IsConfirm() bool {
	return p.ConfirmedAt != nil
}

// IsLock is lock?
func (p *User) IsLock() bool {
	return p.LockedAt != nil
}

//SetGravatarLogo set logo by gravatar
func (p *User) SetGravatarLogo() {
	buf := md5.Sum([]byte(strings.ToLower(p.Email)))
	p.Logo = fmt.Sprintf("https://gravatar.com/avatar/%s.png", hex.EncodeToString(buf[:]))
}

//SetUID generate uid
func (p *User) SetUID() {
	p.UID = uuid.New().String()
}

func (p User) String() string {
	return fmt.Sprintf("%s<%s>", p.Name, p.Email)
}

// Attachment attachment
type Attachment struct {
	nut.Model

	Title        string
	URL          string `orm:"column(url)"`
	Length       int64
	MediaType    string
	ResourceID   uint `orm:"column(resource_id)"`
	ResourceType string
	UserID       uint `orm:"column(user_id)"`
}

// TableName table name
func (*Attachment) TableName() string {
	return "attachments"
}

// IsPicture is picture?
func (p *Attachment) IsPicture() bool {
	return strings.HasPrefix(p.MediaType, "image/")
}

// Log log
type Log struct {
	nut.Timestamp

	Message string
	Type    string
	IP      *string `orm:"column(ip)"`
	UserID  uint    `orm:"column(user_id)"`
}

// TableName table name
func (*Log) TableName() string {
	return "logs"
}

func (p Log) String() string {
	return fmt.Sprintf("%s: [%s]\t %s", p.CreatedAt.Format(time.ANSIC), *p.IP, p.Message)
}

// Policy policy
type Policy struct {
	nut.Model

	StartUp  time.Time
	ShutDown time.Time
	UserID   uint `orm:"column(user_id)"`
	RoleID   uint `orm:"column(role_id)"`
}

//Enable is enable?
func (p *Policy) Enable() bool {
	now := time.Now()
	return now.After(p.StartUp) && now.Before(p.ShutDown)
}

// TableName table name
func (*Policy) TableName() string {
	return "policies"
}

// Role role
type Role struct {
	nut.Model

	Name         string
	ResourceID   uint `orm:"column(resource_id)"`
	ResourceType string
}

// TableName table name
func (*Role) TableName() string {
	return "roles"
}

func (p Role) String() string {
	return fmt.Sprintf("%s@%s://%d", p.Name, p.ResourceType, p.ResourceID)
}

func init() {
	orm.RegisterModel(new(User), new(Role), new(Policy), new(Log), new(Attachment))
}
