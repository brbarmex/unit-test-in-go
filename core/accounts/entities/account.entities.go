package entities

import "time"

type Account struct {
	ID        uint   `gorm:"primaryKey" json:"-"`
	HashCode  string `gorm:"index:idx_hascode,unique,not null,size:37" json:"id"`
	UserName  string `gorm:"not null" json:"name"`
	Email     string `gorm:"index:idx_email,unique,not null" json:"email"`
	Password  string `gorm:"not null" json:"-"`
	Actived   bool   `gorm:"not null, default:false"`
	CreatedAt time.Time
	UpdatedAt time.Time `json:"-"`
}
