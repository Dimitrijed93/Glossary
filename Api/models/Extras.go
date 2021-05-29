package models

import (
	"fmt"
	"net/http"

	"github.com/dimitrijed93/glossary/api/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type ExtraItem struct {
	Id          uint `json:"id" gorm:"primaryKey"`
	Type        Type `json:"type" gorm:"association_foreignkey:Id" validate:"required"`
	TypeID      int
	EntryItem   EntryItem `json:"entryItem" gorm:"association_foreignkey:Id"`
	EntryItemID int
	Name        string `json:"name" validate:"required"`
	Value       string `json:"value" validate:"required"`
}

func NewExtraItem() *ExtraItem {
	return &ExtraItem{}
}

func (ei *ExtraItem) Url() string {
	return "/extras"
}

func (ei *ExtraItem) DetailUrl() string {
	return "/extras/:id"
}

func (ei *ExtraItem) GetAll(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		var items []ExtraItem
		db.Find(&items)
		return c.JSON(http.StatusOK, items)
	}

}

func (ei *ExtraItem) GetById(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		var item Entry
		itemId := c.Param("id")
		db.Model(&item).Where("id=?", itemId)

		return c.JSON(http.StatusOK, item)
	}

}

func (ei *ExtraItem) Create(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		item := new(ExtraItem)

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

func (ei *ExtraItem) Update(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		item := new(Type)
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

func (ei *ExtraItem) Delete(db *gorm.DB, v *util.Validator) func(c echo.Context) error {
	return func(c echo.Context) error {
		itemId := c.Param("id")
		item := &ExtraItem{}
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
