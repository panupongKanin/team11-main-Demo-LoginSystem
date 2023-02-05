package entity

import (
	"gorm.io/gorm"
)

type Role struct {
	gorm.Model
	RoleType string

	Customers         []Customer         `gorm:"ForeignKey:ROLE_ID"`
}