package nut

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
	"github.com/google/uuid"
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

// User user
type User struct {
	Model

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
	LastSignInIP    string `orm:"column(last_sign_in_ip)"`
	CurrentSignInAt *time.Time
	CurrentSignInIP string `orm:"column(current_sign_in_ip)"`
	ConfirmedAt     *time.Time
	LockedAt        *time.Time

	Logs []*Log `orm:"reverse(many)"`
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
	Model

	Title        string
	URL          string `orm:"column(url)"`
	Length       int64
	MediaType    string
	ResourceID   uint `orm:"column(resource_id)"`
	ResourceType string

	User *User `orm:"rel(fk)"`
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
	Timestamp
	Message string
	Type    string
	IP      string

	User *User `orm:"rel(fk)"`
}

// TableName table name
func (*Log) TableName() string {
	return "logs"
}

func (p Log) String() string {
	return fmt.Sprintf("%s: [%s]\t %s", p.CreatedAt.Format(time.ANSIC), p.IP, p.Message)
}

// Policy policy
type Policy struct {
	Model

	StartUp  time.Time
	ShutDown time.Time

	User *User `orm:"rel(fk)"`
	Role *Role `orm:"rel(fk)"`
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
	Model

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

// Vote vote
type Vote struct {
	Model

	Point        int
	ResourceID   uint `orm:"column(resource_id)"`
	ResourceType string
}

// TableName table name
func (*Vote) TableName() string {
	return "votes"
}

// LeaveWord leave-word
type LeaveWord struct {
	Timestamp
	Body string
	Type string
}

// TableName table name
func (*LeaveWord) TableName() string {
	return "leave_words"
}

// Link link
type Link struct {
	Model
	Loc       string
	Href      string
	Label     string
	SortOrder int
}

// TableName table name
func (*Link) TableName() string {
	return "links"
}

// Card card
type Card struct {
	Model

	Loc       string
	Title     string
	Summary   string
	Type      string
	Href      string
	Logo      string
	SortOrder int
	Action    string
}

// TableName table name
func (*Card) TableName() string {
	return "cards"
}

// FriendLink friend_links
type FriendLink struct {
	Model

	Title     string
	Home      string
	Logo      string
	SortOrder int
}

// TableName table name
func (*FriendLink) TableName() string {
	return "friend_links"
}

func init() {
	orm.RegisterModel(
		new(Locale), new(Setting),
		new(User), new(Role), new(Policy), new(Log),
		new(Attachment), new(Vote),
		new(LeaveWord), new(FriendLink),
		new(Link), new(Card),
	)

}
