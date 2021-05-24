package models

import (
	"fmt"
	"net/http"

	"github.com/dimitrijed93/glossary/api/util"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Type struct {
	Id    uint   `gorm:"primaryKey" json:"id"`
	Value string `json:"value" validate:"required"`
}

func NewType() *Type {
	return &Type{}
}

func (t *Type) Url() string {
	return "/types"
}

func (t *Type) DetailUrl() string {
	return "/types/:id"
}

func (t *Type) GetAll(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		var types []Type
		db.Find(&types)
		return c.JSON(http.StatusOK, types)
	}

}

func (t *Type) Create(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		typeObj := new(Type)

		if err := c.Bind(typeObj); err != nil {
			return fmt.Errorf("id is mandatory for update")
		}

		if err := v.Validate(typeObj); err != nil {
			return err
		}

		db.Save(typeObj)

		return c.JSON(http.StatusCreated, typeObj)
	}

}

func (t *Type) Update(db *gorm.DB, v *util.Validator) func(c echo.Context) error {

	return func(c echo.Context) error {
		typeObj := new(Type)
		if err := c.Bind(typeObj); err != nil {
			return err
		}

		typeId := c.Param("id")

		if typeId == "0" {
			return v.InvalidId()
		}

		if err := v.Validate(typeObj); err != nil {
			return err
		}

		db.Model(&typeObj).Where("id=?", typeId).Updates(&typeObj)

		return c.JSON(http.StatusOK, typeObj)
	}

}

func (t *Type) Delete(db *gorm.DB, v *util.Validator) func(c echo.Context) error {
	return func(c echo.Context) error {
		typeId := c.Param("id")
		typeObj := &Type{}
		db.First(&typeObj, typeId)

		if typeId == "0" {
			return v.InvalidId()
		}

		if typeObj != nil {
			db.Delete(&typeObj)
		}

		return c.JSON(http.StatusNoContent, typeObj)
	}
}
