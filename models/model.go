package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID    uint   `gorm:"column:id;primaryKey"`
	Name  string `gorm:"column:name"`
	Email string `gorm:"column:email;unique"`
}
