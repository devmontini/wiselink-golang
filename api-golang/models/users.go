package models

type Users struct {
	Id       int      `json:"id" gorm:"primaryKey"`
	Username string   `json:"username" gorm:"unique;not null"`
	Password string   `json:"password" gorm:"not null"`
	Admin    bool     `json:"admin" gorm:"default:false"`
	Eventss  []Events `json:"events" gorm:"many2many:users_eventss;"`
}
