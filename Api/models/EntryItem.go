package models

import "gorm.io/gorm"

type EntryItem struct {
	gorm.Model

	Id     uint `gorm:"primaryKey"`
	Value  string
	Extras []ExtraItem `gorm:"association_foreignkey:Id"`
}
