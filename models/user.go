package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email    string `gorm:"not null;unique_index" json:"email"`
	Password string `gorm:"type:varchar(300);" json:"password"`
	State    string `gorm:"type:varchar(300);not null;default:'active'" json:"state"`

	Person   Person `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"person"`
	PersonID uint   `json:"personId"`
}
