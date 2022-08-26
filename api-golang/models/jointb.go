package models

type UsersEvents struct {
	UsersID  int `gorm:"primaryKey"`
	EventsID int `gorm:"primaryKey"`
}
