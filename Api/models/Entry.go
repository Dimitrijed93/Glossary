package models

import "gorm.io/gorm"

type Entry struct {
	gorm.Model

	Id            uint `gorm:"primaryKey"`
	Definition    string
	Folder        Folder `gorm:"association_foreignkey:Id"`
	FolderID      int
	SourceID      int
	DestinationID int
	Source        EntryItem `gorm:"association_foreignkey:Id"`
	Destination   EntryItem `gorm:"association_foreignkey:Id"`
}
