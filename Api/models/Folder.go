package models

import "gorm.io/gorm"

type Folder struct {
	gorm.Model

	Id   uint `gorm:"primaryKey"`
	Name string
}
