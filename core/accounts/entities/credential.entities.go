package entities

import "time"

type Credential struct {
	ID        uint      `gorm:"primaryKey" json:"-"`
	HashCode  string    `gorm:"index:idx_hascode,unique,not null,size:37" json:"id"`
	Email     string    `gorm:"index:idx_email,unique,not null" json:"email"`
	Password  string    `gorm:"not null" json:"-"`
	Actived   bool      `gorm:"not null, default:false"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
