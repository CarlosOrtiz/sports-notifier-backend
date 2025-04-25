package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model

	Name     string `gorm:"type:varchar(300);not null" json:"name"`
	Lastname string `gorm:"type:varchar(300)" json:"lastname"`

	UserID uint `gorm:"uniqueIndex" json:"userId"`
}
