package models

type UsersEvents struct {
	UsersID  int `json:"usersid" gorm:"primaryKey"`
	EventsID int `json:"eventsid" gorm:"primaryKey"`
}
