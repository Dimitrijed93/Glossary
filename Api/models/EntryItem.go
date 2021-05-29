package models

import (
	"fmt"
	"net/http"

	"github.com/dimitrijed93/glossary/api/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type EntryItem struct {
	gorm.Model

	Id     uint `gorm:"primaryKey"`
	Value  string
	Extras []ExtraItem `gorm:"association_foreignkey:Id"`
}

func NewEntryItem() *EntryItem {
	return &EntryItem{}
}

func (entryItem *EntryItem) Url() string {
	return "/items"
}

func (entryItem *EntryItem) DetailUrl() string {
	return "/items/:id"
}

func (entryItem *EntryItem) GetAll(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		var items []EntryItem
		db.Find(&items)
		return c.JSON(http.StatusOK, items)
	}

}

func (entryItem *EntryItem) GetById(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		var entryItem Entry
		itemId := c.Param("id")
		db.Model(&entryItem).Where("id=?", itemId)

		return c.JSON(http.StatusOK, entryItem)
	}

}

func (entryItem *EntryItem) Create(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		item := new(EntryItem)

		if err := c.Bind(item); err != nil {
			return fmt.Errorf("id is mandatory for update")
		}

		if err := v.Validate(item); err != nil {
			return err
		}

		db.Save(item)

		return c.JSON(http.StatusCreated, item)
	}

}

func (entryItem *EntryItem) Update(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		item := new(EntryItem)
		if err := c.Bind(item); err != nil {
			return err
		}

		itemId := c.Param("id")

		if itemId == "0" {
			return v.InvalidId()
		}

		if err := v.Validate(item); err != nil {
			return err
		}

		db.Model(&item).Where("id=?", itemId).Updates(&item)

		return c.JSON(http.StatusOK, item)
	}

}

func (entryItem *EntryItem) Delete(db *gorm.DB, v *util.Validator) func(c echo.Context) error {
	return func(c echo.Context) error {
		itemId := c.Param("id")
		item := &EntryItem{}
		db.First(&item, itemId)

		if itemId == "0" {
			return v.InvalidId()
		}

		if item != nil {
			db.Delete(&item)
		}

		return c.JSON(http.StatusNoContent, item)
	}
}
