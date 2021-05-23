package models

import "gorm.io/gorm"

type ExtraItem struct {
	gorm.Model

	Id          uint `gorm:"primaryKey"`
	Type        Type `gorm:"association_foreignkey:Id"`
	TypeID      int
	EntryItem   EntryItem `gorm:"association_foreignkey:Id"`
	EntryItemID int
	Name        string
	Value       string
}

type Extras = []ExtraItem
