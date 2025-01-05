package entities

import "time"

type Student struct {
	ID        int       `gorm:"primaryKey" json:"id"`
	FirstName string    `gorm:"column:first_name" json:"first_name"`
	LastName  string    `gorm:"column:last_name" json:"last_name"`
	GroupID   int       `gorm:"column:group_id" json:"group_id"`
	Birthday  time.Time `gorm:"column:birthday" json:"birthday"`
	Comments  string    `gorm:"column:comments" json:"comments"`
}
