package entities

import "time"

type Account struct {
	ID              uint64    `gorm:"primaryKey" json:"-"`
	HashCode        string    `gorm:"index:idx_act_hascode,unique,not null,size:37" json:"id"`
	FullName        string    `gorm:"not null" json:"fullName" validate:"min=3,max=80,regexp=^[a-zA-Z]*$"`
	Identification  string    `gorm:"index:idx_identification,unique,not null" json:"identification" validate:"len=10,nonzero"`
	Birth           string    `gorm:"not null" json:"birth"`
	AddressPostCode string    `gorm:"not null" json:"postCode" validate:"nonzero, len=8"`
	Address         string    `gorm:"not null" json:"address" validate:"nonzero"`
	AddressNumber   string    `gorm:"not null" json:"number" validate:"nonzero"`
	AddressDetails  string    `gorm:"not null" json:"complement"`
	AddressNeighbor string    `gorm:"not null" json:"neighbor" validate:"nonzero"`
	AddressCity     string    `gorm:"not null" json:"city" validate:"nonzero"`
	AddressState    string    `gorm:"not null" json:"state" validate:"nonzero"`
	AddressCountry  string    `gorm:"not null" json:"country" validate:"nonzero"`
	PhoneNumber     string    `gorm:"not null" json:"phone" validate:"nonzero"`
	Email           string    `gorm:"index:idx_act_email,unique,not null" json:"email"`
	Password        string    `gorm:"not null" json:"-" validate:"nonzero"`
	CreatedAt       time.Time `gorm:"not null" json:"create_at"`
	UpdatedAt       time.Time `gorm:"not null" json:"update_at"`
}
