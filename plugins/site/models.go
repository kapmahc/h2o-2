package site

import (
	"github.com/astaxie/beego/orm"
	"github.com/kapmahc/h2o/plugins/nut"
)

// Vote vote
type Vote struct {
	nut.Model

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
	nut.Timestamp
	Body string
	Type string
}

// TableName table name
func (*LeaveWord) TableName() string {
	return "leave_words"
}

// Link link
type Link struct {
	nut.Model

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
	nut.Model

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
	nut.Model

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
		new(Vote),
		new(LeaveWord), new(FriendLink),
		new(Link), new(Card),
	)

}
