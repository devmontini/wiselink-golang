package models

import (
	"time"
)

type Events struct {
	Id        int       `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"unique;not null"`
	ShortDesc string    `json:"short_desc" gorm:"not null"`
	LongDesc  string    `json:"long_desc" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"not null"`
	Organizer string    `json:"organizer" gorm:"not null"`
	Place     string    `json:"place" gorm:"not null"`
	Status    bool      `json:"status" gorm:"default:false"`
}
