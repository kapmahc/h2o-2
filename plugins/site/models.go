package site

import "time"

// LeaveWord leave-word
type LeaveWord struct {
	ID        uint      `gorm:"primary_key" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Body      string    `json:"body"`
	Type      string    `json:"type"`
}

// TableName table name
func (LeaveWord) TableName() string {
	return "leave_words"
}
